CREATE TABLE IF NOT EXISTS refresh_tokens (
  id BIGSERIAL PRIMARY KEY,
  token TEXT NOT NULL,
  user_id BIGINT NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  revoked BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  CONSTRAINT uni_refresh_tokens_token UNIQUE (token)
);
