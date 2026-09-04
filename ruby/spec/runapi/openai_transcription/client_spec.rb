# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::OpenaiTranscription::Resources::SpeechToText do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:resource) { described_class.new(http) }

  it "posts repeated multipart fields and preserves raw text" do
    allow(http).to receive(:request).and_return("hello from audio\n")
    result = resource.run(file: "/tmp/interview.mp3", model: "gpt-transcribe", languages: ["en", "zh"], keywords: ["RunAPI", "release notes"], response_format: "text")

    expect(http).to have_received(:request) do |method, path, body:, options:|
      expect([method, path]).to eq([:post, "/v1/audio/transcriptions"])
      expect(body.fields["languages[]"]).to eq(["en", "zh"])
      expect(body.fields["keywords[]"]).to eq(["RunAPI", "release notes"])
      expect(body.files["file"].filename).to eq("interview.mp3")
    end
    expect(result).to eq("hello from audio\n")
  end

  it "defaults model and preserves JSON" do
    allow(http).to receive(:request).and_return(RunApi::Core::Response.new(body: {"text" => "hello"}, headers: {}))
    result = resource.run(file: "/tmp/sample.wav")

    expect(http).to have_received(:request) do |_, _, body:, options:|
      expect(body.fields["model"]).to eq("whisper-1")
    end
    expect(result).to eq("text" => "hello")
  end

  it "follows an accepted task result and preserves VTT text" do
    location = "https://api.runapi.ai/api/v1/tasks/task-1/result"
    allow(http).to receive(:request).and_return(
      RunApi::Core::Response.new(
        body: {"id" => "task-1", "status" => "processing"},
        headers: {"Location" => location, "Content-Type" => "application/json"},
        status: 202
      ),
      RunApi::Core::Response.new(
        body: {"id" => "task-1", "status" => "processing"},
        headers: {"Content-Type" => "application/json", "Retry-After" => "0"},
        status: 200
      ),
      RunApi::Core::Response.new(
        body: {
          "id" => "task-1",
          "status" => "completed",
          "response" => {
            "status" => 200,
            "content_type" => "text/vtt",
            "headers" => {},
            "body" => "WEBVTT\n\n00:00.000 --> 00:01.000\nHello\n"
          }
        },
        headers: {"Content-Type" => "application/json"},
        status: 200
      )
    )

    result = resource.run(file: "/tmp/interview.mp3", response_format: "vtt")

    expect(result).to eq("WEBVTT\n\n00:00.000 --> 00:01.000\nHello\n")
  end
end
