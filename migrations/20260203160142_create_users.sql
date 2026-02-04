-- +goose Up
-- +goose StatementBegin

CREATE TYPE user_role AS ENUM (
    'user',
    'admin',
    'engineer'
);

CREATE TABLE users (
    user_id UUID,
    chat_id BIGINT NOT NULL UNIQUE,

    name VARCHAR(128),
    surname VARCHAR(128),
    username VARCHAR(64),
    role user_role NOT NULL DEFAULT 'user',

    state VARCHAR(128) NOT NULL,
    last_interaction TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users;
DROP TYPE user_role;
-- +goose StatementEnd
