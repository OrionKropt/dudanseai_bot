package domain

type VideoNote struct {
	ChatID ChatID
	File   File
}

func NewVideoNote(chatID ChatID, file File) *VideoNote {
	return &VideoNote{ChatID: chatID, File: file}
}
