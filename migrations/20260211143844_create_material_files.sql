-- +goose Up
-- +goose StatementBegin
CREATE TABLE material_files (
    material_id UUID,
    file_data BYTEA NOT NULL,

    PRIMARY KEY (material_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE material_files
-- +goose StatementEnd
