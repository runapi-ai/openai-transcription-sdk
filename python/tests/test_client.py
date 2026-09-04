from pathlib import Path

import pytest

from runapi.core import ApiResponse
from runapi.openai_transcription import OpenaiTranscriptionClient


class FakeHttp:
    def __init__(self, *responses):
        self.responses = list(responses)
        self.calls = []
        self.options = []

    def request(self, method, path, body=None, options=None):
        self.calls.append((method, path, body))
        self.options.append(options)
        return self.responses.pop(0)


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


@pytest.mark.parametrize(
    ("response_format", "content_type", "terminal_body", "expected"),
    [
        ("json", "application/json", {"text": "hello"}, {"text": "hello"}),
        (
            "vtt",
            "text/vtt; charset=utf-8",
            "WEBVTT\n\n00:00.000 --> 00:01.000\nHello",
            "WEBVTT\n\n00:00.000 --> 00:01.000\nHello",
        ),
    ],
)
def test_run_follows_accepted_task_result_and_preserves_terminal_format(
    tmp_path: Path, response_format, content_type, terminal_body, expected
):
    audio = tmp_path / "sample.wav"
    audio.write_bytes(b"audio")
    location = "https://runapi.ai/api/v1/tasks/task_1"
    fake = FakeHttp(
        ApiResponse({"id": "task_1", "status": "processing"}, {"Location": location}, status_code=202),
        ApiResponse(
            {
                "id": "task_1",
                "status": "completed",
                "response": {
                    "status": 200,
                    "content_type": content_type,
                    "headers": {},
                    "body": terminal_body,
                },
            }
        ),
    )
    client = OpenaiTranscriptionClient(api_key="k", http_client=fake)

    result = client.speech_to_text.run(str(audio), response_format=response_format)

    assert result == expected
    assert [call[:2] for call in fake.calls] == [
        ("post", "/v1/audio/transcriptions"),
        ("get", location),
    ]
    assert fake.calls[0][2].files["file"].filename == "sample.wav"
    assert fake.options[0].headers["Idempotency-Key"]
    assert fake.options[1].headers == fake.options[0].headers
