package postgresql

import (
	"context"
	"dudanseai_bot/internal/domain"
	"fmt"
)

type UserRepository struct {
	client *Client
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	q := `
		INSERT INTO users
    		(user_id, chat_id, name, surname, username, role, state)
		VALUES
	    	($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.client.Exec(ctx, q, u.ID, u.ChatID, u.Name, u.Surname, u.Username, u.Role, u.State)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) Delete(ctx context.Context, id domain.UserID) error {
	q := `
	DELETE FROM users 
	WHERE user_id = $1
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

func (r *UserRepository) FindOne(ctx context.Context, id domain.UserID) (domain.User, error) {
	q := `
	SELECT 
		user_id, chat_id, name, surname, username, role, state, last_interaction, created_at
	FROM users
	WHERE user_id = $1
	`
	var u domain.User
	err := r.client.QueryRow(ctx, q, id).Scan(&u.ID, &u.ChatID, &u.Name, &u.Surname, &u.Username, &u.Role, &u.State, &u.LastInteraction, &u.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (r *UserRepository) Update(ctx context.Context, u domain.User) error {
	q := `
		UPDATE users
		SET name = $1, surname = $2, username = $3, role = $4, state = $5, last_interaction = $6
		WHERE user_id = $7
	`
	commandTag, err := r.client.Exec(ctx, q, u.Name, u.Surname, u.Username, u.Role, u.State, u.LastInteraction, u.ID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("user not found id: %s", u.ID.String())
	}
	return nil
}

func NewUserRepository(client *Client) (UserRepository, error) {
	if err := client.Ping(context.Background()); err != nil {
		return UserRepository{}, err
	}
	return UserRepository{client: client}, nil
}
