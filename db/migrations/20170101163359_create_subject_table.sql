
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE subjects (
  	id serial PRIMARY KEY,
    title varchar,
    content text,
    asset varchar,
    action varchar,
    params varchar,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE subjects;

