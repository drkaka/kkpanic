package kkpanic

import (
	"errors"
	"testing"
)

func TestP(t *testing.T) {
	var err error
	P(err)

	err = errors.New("Error.")
	P(err)
}
