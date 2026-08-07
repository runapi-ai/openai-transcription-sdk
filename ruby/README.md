# RunAPI OpenAI Transcription SDK for Ruby

OpenAI-compatible audio transcription through RunAPI.

```ruby
client = RunApi::OpenaiTranscription::Client.new
result = client.speech_to_text.run(file: "audio.mp3")
```

The result is a Hash for JSON response formats and a String for text or subtitle formats. See [OpenAI Transcription](https://runapi.ai/models/openai-transcription) for supported models, formats, and pricing.
