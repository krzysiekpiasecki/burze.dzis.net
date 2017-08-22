package burzedzisnet

import (
	"os"
)

func apikeyFromEnv() string {
	return os.Getenv("BURZEDZISNET_APIKEY")
}

// client represents a
type client struct {
	apikey string
}

// NewClient returns
func NewClient(apiKey string) *client {
	return &client{apikey: apiKey}
}

// APIKey returns used API key by the client.
func (c *client) APIKey() string {
	k := (*c).apikey
	if k == "" {
		return apikeyFromEnv()
	}
	return k
}

// Auth sends a client key and returns true, when the key is valid.
//
// Otherwise it returns false
func (c *client) Auth() (bool, error) {
	p := ptransf(authReqParams{apikey: c.APIKey()})
	req, err := renderAuthRequest(p)
	if err != nil {
		return false, err
	}
	resp := soapRequest(req)
	return parseAuth(resp), nil
}

// MyComplexTypeMiejscowosc sends
func (c *client) MyComplexTypeMiejscowosc(name string) (myComplexTypeMiejscowosc, error) {
	p := ptransf(myComplexTypeMiejscowoscReqParams{apikey: c.APIKey(), name: name})
	req, err := renderMyComplexTypeMiejscowoscRequest(p)
	if err != nil {
		return myComplexTypeMiejscowosc{X: 0, Y: 0}, err
	}
	resp := soapRequest(req)
	loc := parseLocation(resp)
	return loc, nil
}

// MyComplexTypeBurza sends
func (c *client) MyComplexTypeBurza(x, y float64, r int) (myComplexTypeBurza, error) {
	p := ptransf(myComplexTypeBurzaReqParams{apikey: c.APIKey(), x: x, y: y, radius: r})
	req, err := renderMyComplexTypeBurzaRequest(p)
	if err != nil {
		return myComplexTypeBurza{}, err
	}
	resp := soapRequest(req)
	MyComplexTypeBurza := parseMyComplexTypeBurza(resp)
	return MyComplexTypeBurza, nil
}

// MyComplexTypeOstrzezenia sends
func (c *client) MyComplexTypeOstrzezenia(x, y float64) (myComplexTypeOstrzezenia, error) {
	p := ptransf(myComplexTypeOstrzezeniaReqParams{apikey: c.APIKey(), x: x, y: y})
	req, err := renderMyComplexTypeOstrzezeniaRequest(p)
	if err != nil {
		return myComplexTypeOstrzezenia{}, err
	}
	resp := soapRequest(req)
	MyComplexTypeOstrzezenia := parseMyComplexTypeOstrzezenia(resp)
	return MyComplexTypeOstrzezenia, nil
}