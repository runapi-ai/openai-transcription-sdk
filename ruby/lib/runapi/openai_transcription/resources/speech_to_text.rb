# frozen_string_literal: true

module RunApi
  module OpenaiTranscription
    module Resources
      class SpeechToText
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/v1/audio/transcriptions"

        def initialize(http)
          @http = http
        end

        def run(file:, options: nil, **params)
          params = compact_params({file:, model: params.delete(:model) || "whisper-1", **params})
          validate_contract!(CONTRACT["speech-to-text"], params)
          fields = params.except(:file, :languages, :keywords, :timestamp_granularities)
          [:languages, :keywords, :timestamp_granularities].each do |key|
            fields["#{key}[]"] = params[key] if params.key?(key)
          end
          body = RunApi::Core::MultipartBody.new(
            fields:,
            files: {file: RunApi::Core::MultipartFile.new(path: file, filename: File.basename(file), content_type: content_type(file))}
          )
          response = @http.request(:post, ENDPOINT, body:, **(options ? {options:} : {}))
          response.is_a?(RunApi::Core::Response) ? response.body : response
        end

        private

        def content_type(path)
          {".mp3" => "audio/mpeg", ".m4a" => "audio/mp4", ".wav" => "audio/wav", ".webm" => "audio/webm", ".ogg" => "audio/ogg"}.fetch(File.extname(path).downcase, "application/octet-stream")
        end
      end
    end
  end
end
