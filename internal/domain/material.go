package domain

import (
	"dudanseai_bot/internal/infrastructure/uuid"
	"net/url"
	"time"
)

type MaterialID uuid.UUID

func (m MaterialID) String() string {
	return uuid.UUID(m).String()
}

type MaterialType string

const (
	MaterialPDF  MaterialType = "pdf"
	MaterialLink MaterialType = "link"
)

type MaterialFile struct {
	MaterialID MaterialID
	Data       []byte
}

type Material struct {
	ID   MaterialID
	Type MaterialType

	Title       string
	Description string
	URL         url.URL

	FileName        string
	MimeType        string
	FileSizeInBytes int64

	TelegramFileID string
	CreatedAt      time.Time
}

func NewMaterial(t MaterialType, title, description string) Material {
	return Material{ID: MaterialID(uuid.Generate([]byte(title))), Type: t, Title: title, Description: description}
}

func (m *Material) ParseAndSetURL(urlString string) error {
	u, err := url.Parse(urlString)
	if err != nil {
		return err
	}
	m.URL = *u
	return nil
}
