package example1

import (
	"context"
	"fmt"
	"log"
)

type Service struct {
	logger log.Logger
}

var ErrEmptyResult = errors.New("the result is zero")

func (s Service) Operation(ctx context.Context, param1, param2 int) error {
	op := anotherCheck(ctx, param1)

	if op > threshold {
		return nil
	}

	result := Operation(param1, param2)

	if result == 0 {
		return ErrEmptyResult
	}

	// do some other operations

	if err := s.db.Persist(ctx, Calculations); err != nil {
		return fmt.Errorf("cannot persist X: %w", err)
	}

	return nil
}
