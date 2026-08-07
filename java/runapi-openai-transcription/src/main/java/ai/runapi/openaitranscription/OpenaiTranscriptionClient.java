package ai.runapi.openaitranscription;

import ai.runapi.core.BaseClient;
import ai.runapi.core.ClientOptions;
import ai.runapi.core.http.HttpTransport;
import ai.runapi.openaitranscription.resources.SpeechToTextResource;
import java.net.URI;

/** OpenAI-compatible audio transcription client. */
public final class OpenaiTranscriptionClient extends BaseClient {
  private final SpeechToTextResource speechToText;

  private OpenaiTranscriptionClient(ClientOptions options) {
    super(options);
    this.speechToText = new SpeechToTextResource(transport(), options());
  }

  public static Builder builder() { return new Builder(); }
  public SpeechToTextResource speechToText() { return speechToText; }

  public static final class Builder extends BaseClient.Builder<Builder> {
    private Builder() {}
    @Override public Builder apiKey(String value) { return super.apiKey(value); }
    @Override public Builder baseUrl(String value) { return super.baseUrl(value); }
    @Override public Builder baseUrl(URI value) { return super.baseUrl(value); }
    @Override public Builder transport(HttpTransport value) { return super.transport(value); }
    @Override public OpenaiTranscriptionClient build() { return new OpenaiTranscriptionClient(options.build()); }
  }
}
