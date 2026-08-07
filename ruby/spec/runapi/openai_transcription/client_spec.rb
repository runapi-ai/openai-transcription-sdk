# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::OpenaiTranscription::Resources::SpeechToText do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:resource) { described_class.new(http) }

  it "posts repeated multipart fields and preserves raw text" do
    allow(http).to receive(:request).and_return("hello from audio\n")
    result = resource.run(file: "/tmp/interview.mp3", model: "gpt-transcribe", languages: ["en", "zh"], keywords: ["RunAPI", "release notes"], response_format: "text")

    expect(http).to have_received(:request) do |method, path, body:|
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

    expect(http).to have_received(:request) do |_, _, body:|
      expect(body.fields["model"]).to eq("whisper-1")
    end
    expect(result).to eq("text" => "hello")
  end
end
