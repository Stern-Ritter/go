package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Random struct{}

func NewRandom() *Random {
	return &Random{}
}

func (r *Random) Int(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("min %d should be less than max %d", min, max)
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, err
	}

	return int(n.Int64()) + min, nil
}
