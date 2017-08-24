package burzedzisnet

import (
	"os"
)

func apikeyFromEnv() string {
	return os.Getenv("BURZEDZISNET_APIKEY")
}

// client represents a
type client struct {
	apiKey string
}

// NewClient returns
func NewClient(apiKey string) *client {
	return &client{apiKey: apiKey}
}

// APIKey sends a client key and returns true, when the key is valid.
//
// Otherwise it returns false
func (c *client) APIKey() (bool, error) {
	req, err := renderAPIKeyRequest(apiKeyReqParams{apikey: c.apiKey}.transf())
	if err != nil {
		return false, err
	}
	return parseAuth(doSoapRequest(req)), nil
}

// MyComplexTypeMiejscowosc sends
func (c *client) MyComplexTypeMiejscowosc(name string) (myComplexTypeMiejscowosc, error) {
	req, err := renderMyComplexTypeMiejscowoscRequest(myComplexTypeMiejscowoscReqParams{apikey: c.apiKey, name: name}.transf())
	if err != nil {
		return myComplexTypeMiejscowosc{X: 0, Y: 0}, err
	}
	return parseMyComplexTypeMiejscowosc(doSoapRequest(req)), nil
}

// MyComplexTypeBurza sends
func (c *client) MyComplexTypeBurza(x, y float64, r int) (myComplexTypeBurza, error) {
	req, err := renderMyComplexTypeBurzaRequest(myComplexTypeBurzaReqParams{apikey: c.apiKey, x: x, y: y, radius: r}.transf())
	if err != nil {
		return myComplexTypeBurza{}, err
	}
	return parseMyComplexTypeBurza(doSoapRequest(req)), nil
}

// MyComplexTypeOstrzezenia sends
func (c *client) MyComplexTypeOstrzezenia(x, y float64) (myComplexTypeOstrzezenia, error) {
	req, err := renderMyComplexTypeOstrzezeniaRequest(myComplexTypeOstrzezeniaReqParams{apikey: c.apiKey, x: x, y: y}.transf())
	if err != nil {
		return myComplexTypeOstrzezenia{}, err
	}
	return parseMyComplexTypeOstrzezenia(doSoapRequest(req)), nil
}
