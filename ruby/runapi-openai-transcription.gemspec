# frozen_string_literal: true

Dir.chdir(__dir__) do

  Gem::Specification.new do |spec|
    spec.name = "runapi-openai-transcription"
    spec.version = "0.1.0"
    spec.metadata["runapi_slug"] = "openai-transcription"
    spec.authors = ["RunAPI"]
    spec.email = ["contact@runapi.ai"]
    spec.summary = "RunAPI OpenAI Transcription SDK for Ruby"
    spec.description = "OpenAI-compatible audio transcription through RunAPI."
    spec.homepage = "https://runapi.ai/models/openai-transcription"
    spec.license = "Apache-2.0"
    spec.required_ruby_version = ">= 3.1.0"
    spec.metadata["homepage_uri"] = "https://runapi.ai/models/openai-transcription"
    spec.metadata["documentation_uri"] = "https://github.com/runapi-ai/openai-transcription-sdk/blob/main/ruby/README.md"
    spec.metadata["source_code_uri"] = "https://github.com/runapi-ai/openai-transcription-sdk"
    spec.metadata["bug_tracker_uri"] = "https://github.com/runapi-ai/openai-transcription-sdk/issues"
    spec.metadata["changelog_uri"] = "https://github.com/runapi-ai/openai-transcription-sdk/blob/main/CHANGELOG.md"

    spec.files = Dir.glob("lib/**/*") + %w[LICENSE README.md]
    spec.extra_rdoc_files = ["README.md"]
        spec.require_paths = ["lib"]
    spec.add_dependency "runapi-core", "~> 0.3.3"
  end
end
