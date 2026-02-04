package controller

import (
	"dudanseai_bot/internal/domain"
	"errors"
)

type Action func(user domain.User)

type ActionHandler struct {
	Actions map[domain.ActionType]Action
}

func NewActionHandler() *ActionHandler {
	return &ActionHandler{Actions: make(map[domain.ActionType]Action)}
}

func (ah *ActionHandler) RegisterHandler(actionType domain.ActionType, action Action) {
	ah.Actions[actionType] = action
}

func (ah *ActionHandler) Action(user domain.User, actionType domain.ActionType) error {
	action, ok := ah.Actions[actionType]
	if !ok {
		return errors.New("action not found")
	}
	action(user)
	return nil
}
