package postgresql

import (
	"context"
	"database/sql"
	"dudanseai_bot/internal/domain"
	"errors"
	"fmt"
)

type materialRow struct {
	ID   domain.MaterialID
	Type domain.MaterialType

	Title       sql.NullString
	Description sql.NullString
	URL         sql.NullString

	FileName        sql.NullString
	MimeType        sql.NullString
	FileSizeInBytes sql.NullInt64

	TelegramFileID sql.NullString
	CreatedAt      sql.NullTime
}

func (row *materialRow) mapToMaterial() (m domain.Material, err error) {
	m.ID = row.ID
	m.Type = row.Type

	if !row.CreatedAt.Valid {
		return domain.Material{}, errors.New("material creation time is not set")
	}
	m.CreatedAt = row.CreatedAt.Time

	if row.URL.Valid {
		if err := m.ParseAndSetURL(row.URL.String); err != nil {
			return domain.Material{}, err
		}
	}

	if row.Title.Valid {
		m.Title = row.Title.String
	}

	if row.Description.Valid {
		m.Description = row.Description.String
	}

	if row.FileName.Valid {
		m.FileName = row.FileName.String
	}

	if row.MimeType.Valid {
		m.MimeType = row.MimeType.String
	}

	if row.FileSizeInBytes.Valid {
		m.FileSizeInBytes = row.FileSizeInBytes.Int64
	}

	if row.TelegramFileID.Valid {
		m.TelegramFileID = row.TelegramFileID.String
	}
	return m, nil
}

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
		m.URL.String(), m.FileName, m.MimeType, m.FileSizeInBytes, m.TelegramFileID)
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

func (r MaterialsRepository) FindOneMaterial(ctx context.Context, id domain.MaterialID) (m domain.Material, err error) {
	q := `
	SELECT material_id, type, title, description, url, file_name,
			mime_type, file_size_in_bytes, telegram_file_id, created_at
	FROM materials
	WHERE material_id = $1
	`

	var row materialRow
	err = r.client.QueryRow(ctx, q, id).Scan(&row.ID, &row.Type, &row.Title, &row.Description,
		&row.URL, &row.FileName, &row.MimeType, &row.FileSizeInBytes, &row.TelegramFileID, &row.CreatedAt)
	if err != nil {
		return domain.Material{}, err
	}
	if m, err = row.mapToMaterial(); err != nil {
		return domain.Material{}, err
	}
	return m, nil
}

func (r MaterialsRepository) FindMaterialByTitle(ctx context.Context, title string) (m domain.Material, err error) {
	q := `
	SELECT material_id, type, title, description, url, file_name,
			mime_type, file_size_in_bytes, telegram_file_id, created_at
	FROM materials
	WHERE title = $1
	`

	var row materialRow
	err = r.client.QueryRow(ctx, q, title).Scan(&row.ID, &row.Type, &row.Title, &row.Description,
		&row.URL, &row.FileName, &row.MimeType, &row.FileSizeInBytes, &row.TelegramFileID, &row.CreatedAt)
	if err != nil {
		return domain.Material{}, err
	}
	if m, err = row.mapToMaterial(); err != nil {
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
	commandTag, err := r.client.Exec(ctx, q, m.ID, m.Type, m.Title, m.Description, m.URL.String(), m.MimeType, m.FileName, m.FileSizeInBytes, m.TelegramFileID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("material not found id: %s", m.ID.String())
	}
	return nil
}

func (r MaterialsRepository) CreateUserMaterial(ctx context.Context, u domain.UserMaterial) error {
	q := `
	INSERT INTO user_materials
	(user_id, material_id, sent_at)
	VALUES ($1, $2, $3)
`
	_, err := r.client.Exec(ctx, q, u.UserID, u.MaterialID, u.SentAt)
	if err != nil {
		return err
	}
	return nil
}

func (r MaterialsRepository) FindUserMaterialByMaterialID(ctx context.Context, userID domain.UserID, materialID domain.MaterialID) (um domain.UserMaterial, err error) {
	q := `
	SELECT user_material_id, user_id, material_id, sent_at
	FROM user_materials
	WHERE user_id = $1 AND material_id = $2
`
	err = r.client.QueryRow(ctx, q, userID, materialID).Scan(um.ID, um.UserID, um.MaterialID, um.SentAt)
	if err != nil {
		return domain.UserMaterial{}, err
	}
	return um, nil
}

func NewMaterialsRepository(client *Client) (MaterialsRepository, error) {
	if err := client.Ping(context.Background()); err != nil {
		return MaterialsRepository{}, err
	}
	return MaterialsRepository{client: client}, nil
}
