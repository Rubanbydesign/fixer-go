package fixer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
)

func TestLatest(t *testing.T) {
	rates, err := Latest(ctx)
	if !assert.NoError(t, err) {
		return
	}
	assert.Len(t, rates, 31)
}

func TestGet(t *testing.T) {
	rates, err := Get(ctx, USD)
	if !assert.NoError(t, err) {
		return
	}
	assert.Len(t, rates, 31)
}

func TestConvert(t *testing.T) {
	amt, err := Convert(ctx, USD, EUR, 1)
	assert.NoError(t, err)
	assert.True(t, amt < 1)
}
