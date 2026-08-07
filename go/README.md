# RunAPI OpenAI Transcription SDK for Go

OpenAI-compatible audio transcription through RunAPI.

```go
result, err := client.SpeechToText.Run(ctx, openaitranscription.SpeechToTextParams{File: "audio.mp3"})
```

The returned bytes preserve JSON, text, and subtitle response formats. See [OpenAI Transcription](https://runapi.ai/models/openai-transcription) for supported models, formats, and pricing.
