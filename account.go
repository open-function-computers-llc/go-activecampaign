package activecampaign

import (
	"errors"
	"strconv"
)

func VerifyAccountConnection() error {
	r, err := getRequest("api/3/accounts")
	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		return errors.New("Didn't get a valid response code. Got: " + strconv.Itoa(r.StatusCode))
	}
	return nil
}
