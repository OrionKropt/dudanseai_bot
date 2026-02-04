-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_materials (
    user_material_id SERIAL,
    user_id UUID,
    material_id UUID,

    sent_at TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (user_material_id),

    CONSTRAINT fk_user_id FOREIGN KEY (user_id)
    REFERENCES users(user_id)
    ON DELETE CASCADE,

    CONSTRAINT fk_material_id FOREIGN KEY (material_id)
    REFERENCES materials(material_id)
    ON DELETE CASCADE

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_materials;
-- +goose StatementEnd
