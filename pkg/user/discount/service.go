package discount

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type Service struct {
	logger         *zap.Logger
	rate           float64
	userRepository UserRepository
}

func New(logger *zap.Logger, rate float64, repository UserRepository) *Service {
	return &Service{
		logger:         logger,
		rate:           rate,
		userRepository: repository,
	}
}

func (s *Service) CalculateDiscountForUser(ctx context.Context, userID int) (float64, error) {
	if userID <= 100 {
		return 0, errors.New("wrong user")
	}

	balance, err := s.userRepository.GetUserBalance(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("get user balance: %w", err)
	}

	return float64(balance) * s.rate, nil
}
