-- +goose Up
CREATE TABLE IF NOT EXISTS "views" (
  "user_id" text NOT NULL,
  "quote_id" text NOT NULL REFERENCES quotes (id) ON DELETE CASCADE,
  "liked" boolean,
  CONSTRAINT views_pkey PRIMARY KEY ("user_id","quote_id")
);
-- +goose Down
DROP TABLE "views";
