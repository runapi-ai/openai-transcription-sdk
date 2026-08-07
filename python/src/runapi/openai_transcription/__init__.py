"""OpenAI-compatible audio transcription client for RunAPI."""

from runapi.core import AuthenticationError, InsufficientCreditsError, NotFoundError, RateLimitError, ValidationError

from .client import OpenaiTranscriptionClient

__all__ = [
    "OpenaiTranscriptionClient",
    "AuthenticationError",
    "RateLimitError",
    "InsufficientCreditsError",
    "NotFoundError",
    "ValidationError",
]
