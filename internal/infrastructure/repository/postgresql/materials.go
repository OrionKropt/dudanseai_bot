package postgresql

import (
	"context"
	"dudanseai_bot/internal/domain"
	"fmt"
)

type MaterialsRepository struct {
	client *Client
}

func (r MaterialsRepository) CreateMaterial(ctx context.Context, m domain.Material) error {
	q := `
	INSERT INTO materials 
		(material_id, type, title, description, url,
		 file_name, mime_type, file_size_in_bytes, telegram_file_id)
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9)
`
	_, err := r.client.Exec(ctx, q, m.ID, m.Type, m.Title, m.Description,
		m.URL, m.FileName, m.MimeType, m.FileSizeInBytes, m.TelegramFileID)
	if err != nil {
		return err
	}
	return nil
}

func (r MaterialsRepository) DeleteMaterial(ctx context.Context, id domain.MaterialID) error {
	q := `
	DELETE FROM materials
	WHERE material_id = $1
	`
	commandTag, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("user not found id: %s", id.String())
	}
	return nil

}

func (r MaterialsRepository) FindOneMaterial(ctx context.Context, id domain.MaterialID) (domain.Material, error) {
	q := `
	SELECT material_id, type, title, description, url,
		 file_name, mime_type, file_size_in_bytes, telegram_file_id
	FROM materials
	WHERE material_id = $1
	`
	var m domain.Material
	err := r.client.QueryRow(ctx, q, id).Scan(&m.ID, &m.Type, &m.Title, &m.Description,
		&m.URL, &m.FileName, &m.MimeType, &m.FileSizeInBytes, &m.TelegramFileID)
	if err != nil {
		return domain.Material{}, err
	}
	return m, nil
}

func (r MaterialsRepository) FindMaterialByTitle(ctx context.Context, title string) (domain.Material, error) {
	q := `
	SELECT material_id, type, title, description, url,
		 file_name, mime_type, file_size_in_bytes, telegram_file_id
	FROM materials
	WHERE title = $1
	`
	var m domain.Material
	err := r.client.QueryRow(ctx, q, title).Scan(&m.ID, &m.Type, &m.Title, &m.Description,
		&m.URL, &m.FileName, &m.MimeType, &m.FileSizeInBytes, &m.TelegramFileID)
	if err != nil {
		return domain.Material{}, err
	}
	return m, nil
}

func (r MaterialsRepository) UpdateMaterial(ctx context.Context, m domain.Material) error {
	q := `
		UPDATE materials
		SET type = $2, title = $3, description = $4, url = $5,
		    file_name = $6, mime_type = $7, file_size_in_bytes = $8, telegram_file_id = $9
		WHERE material_id = $1
`
	commandTag, err := r.client.Exec(ctx, q, m.ID, m.Type, m.Title, m.Description, m.URL, m.MimeType, m.FileName, m.FileSizeInBytes, m.TelegramFileID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("material not found id: %s", m.ID.String())
	}
	return nil
}

func NewMaterialsRepository(client *Client) (MaterialsRepository, error) {
	if err := client.Ping(context.Background()); err != nil {
		return MaterialsRepository{}, err
	}
	return MaterialsRepository{client: client}, nil
}
