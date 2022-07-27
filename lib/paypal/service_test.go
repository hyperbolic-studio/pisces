package paypal

import (
	"context"
	"testing"

	"github.com/cryptnode-software/pisces/lib"
	"github.com/cryptnode-software/pisces/lib/utility"
)

var (
	env = utility.NewEnv(utility.NewLogger())

	service, err = NewService(env)

	ctx = context.Background()

	order = &lib.Order{
		ID:    1,
		Total: 40.00,
	}
)

func TestGenerateClientSideToken(t *testing.T) {
	if err != nil {
		t.Error(err)
		return
	}
	response, err := service.GenerateClientToken(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	if response.Token == "" {
		t.Error("no client token was returned from response")
		return
	}
}

func TestCreateOrder(t *testing.T) {
	if err != nil {
		t.Error(err)
		return
	}

	_, err := service.CreateOrder(ctx, order)
	if err != nil {
		t.Error(err)
		return
	}
}
