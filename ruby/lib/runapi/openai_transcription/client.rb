# frozen_string_literal: true

module RunApi
  module OpenaiTranscription
    class Client < RunApi::Core::Client
      attr_reader :speech_to_text

      def initialize(api_key: nil, **options)
        super
        @speech_to_text = Resources::SpeechToText.new(http)
      end
    end
  end
end
