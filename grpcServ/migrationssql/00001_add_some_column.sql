-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS askbid
(
    id     BIGSERIAL PRIMARY KEY,
    price  VARCHAR(100),
    volume VARCHAR(100),
    amount VARCHAR(100),
    factor VARCHAR(100),
    type   VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS currencyinfo
(
    id     BIGSERIAL PRIMARY KEY,
    timestamp  INTEGER,
    ask INTEGER REFERENCES askbid(id),
    bid INTEGER REFERENCES askbid(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS currencyinfo;

DROP TABLE IF EXISTS askbid;
-- +goose StatementEnd
