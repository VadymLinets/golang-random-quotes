-- +goose Up
CREATE TABLE IF NOT EXISTS "quotes" (
  "id" text NOT NULL PRIMARY KEY,
  "quote" text NOT NULL,
  "author" text,
  "likes" int
);
-- +goose Down
DROP TABLE "quotes";
