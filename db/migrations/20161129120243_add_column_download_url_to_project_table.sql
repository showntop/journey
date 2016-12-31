
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE projects ADD COLUMN download_url varchar;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE projects DROP COLUMN download_url RESTRICT;
