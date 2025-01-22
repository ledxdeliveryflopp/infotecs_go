-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE wallet (id serial PRIMARY KEY, number varchar unique, balance float);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE wallet;