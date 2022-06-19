package api

import (
	"context"
	"fmt"

	"github.com/ghostec/hall/server/app/api/errors"
)

func HomeHandler(ctx context.Context, in HomeInput) (HomeOutput, error) {
	return HomeOutput{Greeting: fmt.Sprintf("Hello, %s!", in.Name)}, errors.NewBadRequestError("failed to do something")
}

type HomeInput struct {
	Name     string `query:"name"`
	Location string `json:"location"`
}

type HomeOutput struct {
	Greeting string `json:"greeting"`
}
