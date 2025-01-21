-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO wallet (number, balance)

SELECT
    'wlnum' || generate_series(1, 10) AS number,
     100.0 AS balance;


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM wallet WHERE id = generate_series(1, 10);