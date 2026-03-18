package domain

import (
	"os"
	"strings"
)

type FileSource int

const (
	SourceFileID FileSource = iota
	SourcePath
)

type File struct {
	Source         FileSource
	TelegramFileID string
	Path           string
	Bytes          []byte
	FileName       string
}

func NewFileByID(id string) File {
	return File{Source: SourceFileID, TelegramFileID: id}
}

func NewFileFromPath(path string) File {
	parts := strings.Split(path, "/")
	name := parts[len(parts)-1]
	return File{Source: SourcePath, Path: path, FileName: name}
}

func (f *File) Load() (data []byte, err error) {
	if len(f.Bytes) > 0 {
		return f.Bytes, nil
	}
	data, err = os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	f.Bytes = data
	return data, nil
}
