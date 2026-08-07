# frozen_string_literal: true

require "runapi/core"
require_relative "openai_transcription/types"
require_relative "openai_transcription/contract_gen"
require_relative "openai_transcription/resources/speech_to_text"
require_relative "openai_transcription/client"

module RunApi
  module OpenaiTranscription
    AuthenticationError = RunApi::Core::AuthenticationError
    RateLimitError = RunApi::Core::RateLimitError
    InsufficientCreditsError = RunApi::Core::InsufficientCreditsError
    NotFoundError = RunApi::Core::NotFoundError
    ValidationError = RunApi::Core::ValidationError
  end
end
