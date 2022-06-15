package memory

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
}

func (u UserRepository) GetUserBalance(ctx context.Context, userID int) (int, error) {
	if userID < 5 {
		return 0, errors.New("user id wrong")
	}

	if userID == 10 {
		return 0, ErrUserNotFound
	}

	return 99, nil
}
