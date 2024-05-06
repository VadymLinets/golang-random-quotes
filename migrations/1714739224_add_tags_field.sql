-- +goose Up
ALTER TABLE "quotes"
    ADD COLUMN IF NOT EXISTS "tags" text[];
-- +goose Down
ALTER TABLE "quotes" DROP COLUMN IF EXISTS "tags";
