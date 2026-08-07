---
name: openai-transcription
description: Transcribe uploaded audio through RunAPI with an OpenAI-compatible API. Use for one-off transcription, subtitle output, multilingual hints, or application integration. Prefer the RunAPI CLI for manual requests and the target-language SDK for production integration.
documentation: https://runapi.ai/models/openai-transcription.md
provider_page: https://runapi.ai/providers/openai.md
catalog: https://runapi.ai/models.md
metadata:
  openclaw:
    homepage: https://runapi.ai/models/openai-transcription
    requires:
      bins:
      - runapi
    install:
    - kind: brew
      formula: runapi-ai/tap/runapi
      bins:
      - runapi
    envVars:
    - name: RUNAPI_API_KEY
      required: false
      description: Optional RunAPI API key; prefer environment auth or saved CLI config.
---

# OpenAI Transcription on RunAPI

Transcribe a local audio file synchronously. Integration code uses the target-language SDK; manual verification uses the `runapi` CLI.

## Critical: Integration Runtime

- Integration work (app, backend, worker, library, Rails service, Node service, Go service, webhook pipeline, or production codebase) uses the **SDK integration path** for the target language.
- One-off transcription, subtitle export, manual smoke tests, and debugging use the **CLI path** with the `runapi` binary. For full CLI-specific agent guidance, see https://github.com/runapi-ai/cli-skill.
- Never shell out to the `runapi` CLI as the production runtime integration layer.

## SDK integration path

For application integration, confirm the target SDK's install command, multipart file input, synchronous `run` method, response formats, and error classes before using CLI help or raw HTTP examples. Choose the package for the application's runtime:

- JavaScript / TypeScript: `@runapi.ai/openai-transcription`
- Python: `runapi-openai-transcription`
- Ruby: `runapi-openai-transcription`
- Go: `github.com/runapi-ai/openai-transcription-sdk/go`
- Java: `ai.runapi:runapi-openai-transcription`
- PHP: `runapi-ai/openai-transcription`

Call `speech_to_text` with a local file. An omitted model selects `whisper-1`. Append `languages[]`, `keywords[]`, and `timestamp_granularities[]` once per value when using raw multipart HTTP. Completion means the response remains a JSON object for JSON formats or the exact string for text, SRT, and VTT formats.

## CLI path

The `runapi` binary is the one-off and manual testing runtime dependency. Run
`runapi auth status` first, then inspect the current request fields:

```shell
runapi auth status
runapi openai-transcription speech-to-text --help
runapi openai-transcription speech-to-text --file audio.mp3 --response-format text
```

Completion means the command exits successfully and prints the transcription in the selected response format.

## Variants

- `whisper-1` supports JSON, text, verbose JSON, SRT, VTT, and word or segment timestamps.
- `gpt-transcribe` supports JSON or text plus repeated language and keyword hints.
- `language` and `languages` are alternative hint forms; send only one.

## References

- Model overview: https://runapi.ai/models/openai-transcription.md
- whisper-1: https://runapi.ai/models/openai-transcription/whisper-1.md
- gpt-transcribe: https://runapi.ai/models/openai-transcription/gpt-transcribe.md
