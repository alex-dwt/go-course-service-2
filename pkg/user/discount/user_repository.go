package discount

import (
	"context"
)

type UserRepository interface {
	GetUserBalance(ctx context.Context, userID int) (int, error)
}
