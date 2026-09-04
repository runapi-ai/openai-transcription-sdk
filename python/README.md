# RunAPI OpenAI Transcription SDK for Python

OpenAI-compatible audio transcription through RunAPI.

```python
from runapi.openai_transcription import OpenaiTranscriptionClient

client = OpenaiTranscriptionClient()
result = client.speech_to_text.run("audio.mp3")
```

`run()` returns the same terminal result whether the request completes directly or is accepted for background processing. The result is a dictionary for JSON response formats and a string for text or subtitle formats. Call `subscribe(location)` to resume an accepted request from its opaque Task Result URL.

See [OpenAI Transcription](https://runapi.ai/models/openai-transcription) for supported models, formats, and pricing.
