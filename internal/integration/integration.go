package integration

import (
	"chat/internal/config"
	"fmt"
)

type Integrations struct {
	CentrifugoServer CentrifugoServer
}

func NewIntegrations(cfg *config.Config) (*Integrations, error) {
	centrifugoServer, err := NewCentrifugoServerClient(cfg.Integrations.CentrifugoServerUrl, cfg.Integrations.CentrifugoServerXApiKey)
	if err != nil {
		return nil, fmt.Errorf("error on init NewCentrifugoServerClient")
	}
	return &Integrations{
		CentrifugoServer: centrifugoServer,
	}, nil
}
