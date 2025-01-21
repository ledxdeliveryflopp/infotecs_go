-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE transaction (id serial PRIMARY KEY, sender varchar, recipient varchar, amount float,
                          time timestamp DEFAULT current_timestamp);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE transaction;