-- +goose Up
-- +goose StatementBegin
CREATE TABLE materials (
    material_id UUID,

    type VARCHAR(32) NOT NULL,
    title TEXT NOT NULL UNIQUE ,
    description TEXT,

    url TEXT,

    file_name VARCHAR(128),
    mime_type VARCHAR(128),
    file_size_in_bytes BIGINT,

    telegram_file_id TEXT,

    created_at TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (material_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE materials;
-- +goose StatementEnd
