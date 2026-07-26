package util

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/medama-io/medama/db"
	"github.com/medama-io/medama/model"
)

type ApiKeysService struct {
	client db.AppClient
}

func NewApiKeysService(
	client db.AppClient,
) (*ApiKeysService, error) {
	return &ApiKeysService{
		client: client,
	}, nil
}

func (s *ApiKeysService) ExchangeApiKey(ctx context.Context, apiKey string) (*model.User, error) {
	user, err := s.client.GetUserByApiKey(ctx, apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "exchange api key")
	}

	return user, nil
}
