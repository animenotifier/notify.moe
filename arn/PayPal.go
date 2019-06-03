package arn

import (
	"os"

	paypalsdk "github.com/logpacker/PayPal-Go-SDK"
)

var payPal *paypalsdk.Client

// PayPal returns the new PayPal SDK client.
func PayPal() (*paypalsdk.Client, error) {
	if payPal != nil {
		return payPal, nil
	}

	apiBase := paypalsdk.APIBaseSandBox

	if IsProduction() {
		apiBase = paypalsdk.APIBaseLive
	}

	// Create a client instance
	c, err := paypalsdk.NewClient(APIKeys.PayPal.ID, APIKeys.PayPal.Secret, apiBase)
	c.SetLog(os.Stdout)

	if err != nil {
		return nil, err
	}

	return c, nil
}
