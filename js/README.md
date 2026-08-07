# `@runapi.ai/openai-transcription`

JavaScript and TypeScript client for OpenAI-compatible audio transcription through RunAPI.

```ts
const result = await client.speechToText.run({
  file: new Blob([audioBytes], { type: 'audio/mpeg' }),
  filename: 'audio.mp3',
});
```

JSON formats return an object; text and subtitle formats return their exact string. See [OpenAI Transcription](https://runapi.ai/models/openai-transcription) for supported models, formats, and pricing.
