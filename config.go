package activecampaign

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Logger  *logrus.Logger
	ApiKey  string
	ApiBase string
}

func (c *Config) validate() error {
	if c.Logger == nil {
		return errors.New("Must have a valid loggin instance to use this package")
	}

	if c.ApiKey == "" {
		return errors.New("Must have an API Key set to use this package")
	}

	if c.ApiBase == "" {
		return errors.New("Each ActiveCampaign customer has their own API URL. This needs to be set to use this package")
	}

	if c.ApiBase[len(c.ApiBase)-1:] != "/" {
		return errors.New("Your API Base URL should end with a trailing slash")
	}

	return nil
}
