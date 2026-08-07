"""OpenAI-compatible audio transcription client."""

from __future__ import annotations

from typing import Any, Optional

from runapi.core import ProviderClient

from .resources.speech_to_text import SpeechToText


class OpenaiTranscriptionClient(ProviderClient):
    """Client for synchronous audio transcription."""

    def __init__(self, api_key: Optional[str] = None, **options: Any) -> None:
        super().__init__(api_key, **options)
        http = self._http
        self.speech_to_text = SpeechToText(http)
