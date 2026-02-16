package domain

import "net/url"

type ButtonCallback func(username string, data []byte)
type Button struct {
	Text   string
	URL    *url.URL
	Action ActionType
}
type Keyboard struct {
	markup [][]Button
}
type Message struct {
	ChatID ChatID
	Text   string
	Keys   Keyboard
}

func NewMessage(chatID ChatID, text string, keys Keyboard) *Message {
	return &Message{ChatID: chatID, Text: text, Keys: keys}
}

func NewButton(text string, action ActionType) Button {
	return Button{Text: text, Action: action}
}

func NewButtonURL(text string, u *url.URL) Button {
	return Button{Text: text, URL: u}
}

func NewKeyboard() Keyboard {
	return Keyboard{markup: [][]Button{{}}}
}

func (k *Keyboard) AddButton(name string, action ActionType) {
	k.markup[len(k.markup)-1] = append(k.markup[len(k.markup)-1], NewButton(name, action))
}

func (k *Keyboard) AddButtonURL(name string, u *url.URL) {
	k.markup[len(k.markup)-1] = append(k.markup[len(k.markup)-1], NewButtonURL(name, u))
}

func (k *Keyboard) Row() *Keyboard {
	if len(k.markup[len(k.markup)-1]) > 0 {
		k.markup = append(k.markup, []Button{})
	}
	return k
}

func (k *Keyboard) Markup() [][]Button {
	return k.markup
}
