package cognitoclient

import (
	"github.com/ONSdigital/dp-identity-api/config"
)

type Cognito struct {
	Client interface{}
	AWSUserPoolID string

}

func New(cfg *config.Config) *Cognito {
	return &Cognito{}
}