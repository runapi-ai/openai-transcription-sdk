from pathlib import Path

from runapi.openai_transcription import OpenaiTranscriptionClient


class FakeHttp:
    def __init__(self, response):
        self.response = response
        self.calls = []

    def request(self, method, path, body=None, options=None):
        self.calls.append((method, path, body))
        return self.response


def test_posts_repeated_multipart_fields_and_preserves_text(tmp_path: Path):
    audio = tmp_path / "interview.mp3"
    audio.write_bytes(b"audio")
    fake = FakeHttp("hello from audio\n")
    client = OpenaiTranscriptionClient(api_key="k", http_client=fake)

    result = client.speech_to_text.run(
        str(audio),
        model="gpt-transcribe",
        languages=["en", "zh"],
        keywords=["RunAPI", "release notes"],
        response_format="text",
    )

    method, path, body = fake.calls[0]
    assert (method, path) == ("post", "/v1/audio/transcriptions")
    assert body.fields["languages[]"] == ["en", "zh"]
    assert body.fields["keywords[]"] == ["RunAPI", "release notes"]
    assert body.files["file"].filename == "interview.mp3"
    assert result == "hello from audio\n"


def test_defaults_model_and_preserves_json(tmp_path: Path):
    audio = tmp_path / "sample.wav"
    audio.write_bytes(b"audio")
    fake = FakeHttp({"text": "hello"})
    client = OpenaiTranscriptionClient(api_key="k", http_client=fake)

    result = client.speech_to_text.run(str(audio))

    assert fake.calls[0][2].fields["model"] == "whisper-1"
    assert result == {"text": "hello"}
