# RunAPI OpenAI Transcription SDK for Java

OpenAI-compatible audio transcription through RunAPI.

## Install

```kotlin
implementation("ai.runapi:runapi-openai-transcription:0.1.0")
```

```java
TranscriptionResponse result = client.speechToText().run(
    SpeechToTextParams.builder().file(Path.of("audio.mp3")).build()
);
```

The response preserves its exact body and content type. See [OpenAI Transcription](https://runapi.ai/models/openai-transcription) for supported models, formats, and pricing.
