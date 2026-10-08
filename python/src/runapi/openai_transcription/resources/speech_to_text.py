"""Audio transcription resource."""

from __future__ import annotations

import mimetypes
from pathlib import Path
from typing import Any, Optional

from runapi.core import MultipartBody, MultipartFile, RequestOptions, Resource


class SpeechToText(Resource):
    """Transcribe an uploaded audio file."""

    ENDPOINT = "/v1/audio/transcriptions"
    RESPONSE_CLASS = None

    def run(self, file: str, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        compacted = self._compact_params({"file": file, "model": params.pop("model", None) or "whisper-1", **params})

        path = Path(file)
        fields = {
            key: value
            for key, value in compacted.items()
            if key not in {"file", "languages", "keywords", "timestamp_granularities"}
        }
        for key in ("languages", "keywords", "timestamp_granularities"):
            if key in compacted:
                fields[f"{key}[]"] = compacted[key]
        body = MultipartBody(
            fields=fields,
            files={
                "file": MultipartFile(
                    path=str(path),
                    filename=path.name,
                    content_type=mimetypes.guess_type(path.name)[0] or "application/octet-stream",
                )
            },
        )
        return self._run_hybrid("post", self.ENDPOINT, body=body, options=options)
