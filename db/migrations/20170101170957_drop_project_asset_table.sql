-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE project_assets;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
CREATE TABLE project_assets (
  	id serial PRIMARY KEY,
    project_id serial REFERENCES projects (id),
    url varchar,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);