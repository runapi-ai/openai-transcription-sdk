# Changelog

## [js/v0.2.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/js%2Fv0.2.0), [ruby/v0.3.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/ruby%2Fv0.3.0), [go/v0.3.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/go%2Fv0.3.0), [python/v0.3.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/python%2Fv0.3.0), [java/v0.2.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/java%2Fv0.2.0) - 2026-10-08

### Changed
- Send request parameters to the service without local validation. Model ids and parameter values the service supports work without an SDK upgrade; static types and enum constants remain for completion.
  Migration: Invalid parameters now fail with the validation error built from the service's 400 response, including its status and message, instead of a validation error raised locally before the request. The error type is unchanged: `ValidationError` in JavaScript, Python, and Ruby, `ValidationException` in Java and PHP, and `ErrValidation` in Go.


## [python/v0.2.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/python%2Fv0.2.0) - 2026-09-04

### Added
- Automatically follow accepted transcription Tasks while preserving JSON, text, SRT, and VTT terminal results.

## [go/v0.2.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/go%2Fv0.2.0) - 2026-09-04

### Added
- Add Create, Subscribe, and automatic Run support for accepted transcription Tasks while preserving JSON, text, SRT, and VTT response bytes.

## [ruby/v0.2.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/ruby%2Fv0.2.0) - 2026-09-04

### Added
- Add run and subscribe support for hybrid Task responses.


## [ruby/v0.1.1](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/ruby%2Fv0.1.1) - 2026-08-18

### Changed
- Allow Ruby clients to install the core SDK release that adds persistent Files and multipart Uploads alongside this model SDK.


## [js/v0.1.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/js%2Fv0.1.0), [ruby/v0.1.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/ruby%2Fv0.1.0), [go/v0.1.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/go%2Fv0.1.0), [python/v0.1.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/python%2Fv0.1.0), [java/v0.1.0](https://github.com/runapi-ai/openai-transcription-sdk/releases/tag/java%2Fv0.1.0) - 2026-08-07

### Added
- Add synchronous multipart audio transcription with JSON, text, and subtitle response preservation.
