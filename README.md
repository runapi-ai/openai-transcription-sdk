<p align="center">
  <a href="https://runapi.ai"><img src="https://runapi.ai/icon.svg" height="56" alt="RunAPI"></a>
</p>

<h3 align="center"><a href="https://github.com/runapi-ai/openai-transcription-sdk">OpenAI Transcription API SDK for RunAPI</a></h3>

<p align="center">OpenAI-compatible audio transcription SDKs for JavaScript, Python, Ruby, Go, Java, and PHP on RunAPI.</p>

<div align="center">

[![npm](https://img.shields.io/npm/v/@runapi.ai/openai-transcription)](https://www.npmjs.com/package/@runapi.ai/openai-transcription)
[![PyPI](https://img.shields.io/pypi/v/runapi-openai-transcription)](https://pypi.org/project/runapi-openai-transcription/)
[![RubyGems](https://img.shields.io/gem/v/runapi-openai-transcription)](https://rubygems.org/gems/runapi-openai-transcription)
[![Go Reference](https://pkg.go.dev/badge/github.com/runapi-ai/openai-transcription-sdk/go.svg)](https://pkg.go.dev/github.com/runapi-ai/openai-transcription-sdk/go)
[![Maven Central](https://img.shields.io/maven-central/v/ai.runapi/runapi-openai-transcription)](https://central.sonatype.com/artifact/ai.runapi/runapi-openai-transcription)
[![License](https://img.shields.io/github/license/runapi-ai/openai-transcription-sdk)](https://github.com/runapi-ai/openai-transcription-sdk/blob/main/LICENSE)

</div>

Transcribe audio with `whisper-1` or `gpt-transcribe`. JSON formats return
structured data; text and subtitle formats preserve their exact response body.

## Install

```bash
npm install @runapi.ai/openai-transcription
pip install runapi-openai-transcription
gem install runapi-openai-transcription
go get github.com/runapi-ai/openai-transcription-sdk/go@latest
```

For Java, install `ai.runapi:runapi-openai-transcription:0.2.0`.
The PHP package is released from the split repository at
https://github.com/runapi-ai/openai-transcription-php.

## Quick start

```typescript
import { OpenaiTranscriptionClient } from '@runapi.ai/openai-transcription';

const client = new OpenaiTranscriptionClient();
const result = await client.speechToText.run({
  file: new Blob([audioBytes], { type: 'audio/mpeg' }),
  filename: 'audio.mp3',
});
console.log(result);
```

## Variants

- [whisper-1](https://runapi.ai/models/openai-transcription/whisper-1)
- [gpt-transcribe](https://runapi.ai/models/openai-transcription/gpt-transcribe)

See the [OpenAI Transcription model page](https://runapi.ai/models/openai-transcription)
for current pricing, formats, limits, and usage details.

## License

Licensed under the Apache License, Version 2.0.
