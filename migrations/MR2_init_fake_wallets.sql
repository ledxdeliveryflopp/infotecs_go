-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO wallet (number, balance)

SELECT
    'wlnum' || generate_series(1, 10) AS number,
     100.0 AS balance;


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM wallet WHERE id = 1 and id = 2 and id = 3 and id = 4 and id = 5
                     and id = 6 and id = 7 and id = 8 and id = 9 and id = 10;