package activecampaign

import "github.com/sirupsen/logrus"

var log *logrus.Logger
var apiKey string
var apiBase string

// Init Initialize the package with the required items from Infusionsoft and get it ready for use
func Init(c Config) error {
	err := c.validate()
	if err != nil {
		return err
	}

	log = c.Logger
	apiKey = c.ApiKey
	log.Info("ActiveCampaign initialized with api key " + apiKey)
	apiBase = c.ApiBase
	log.Info("ActiveCampaign initialized with api base url " + apiBase)

	log.Info("Checking to make sure API Key and URL are valid")
	err = VerifyAccountConnection()

	if err != nil {
		return err
	}
	return nil
}
