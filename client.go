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
	p := authReqParams{apikey: c.APIKey()}
	pm := make(map[string]string)
	p.print(&pm)
	req, err := renderAuthRequest(pm)
	if err != nil {
		return false, err
	}
	resp := soapRequest(req)
	return parseAuth(resp), nil
}

// AuthKey sends a parameter as an API key and returns true, when the key is valid.
//
// Otherwise it returns false
func (c *client) AuthKey(apikey string) (bool, error) {
	p := authReqParams{apikey: apikey}
	pm := make(map[string]string)
	p.print(&pm)
	req, err := renderAuthRequest(pm)
	if err != nil {
		return false, err
	}
	resp := soapRequest(req)
	return parseAuth(resp), nil
}

// Locate sends
func (c *client) Locate(name string) (myComplexTypeMiejscowosc, error) {
	return c.AuthLocate(name, c.APIKey())
}

// AuthLocate sends
func (c *client) AuthLocate(name string, apikey string) (myComplexTypeMiejscowosc, error) {
	p := myComplexTypeMiejscowoscReqParams{apikey: apikey, name: name}
	pm := make(map[string]string)
	p.print(&pm)
	req, err := renderLocationRequest(pm)
	if err != nil {
		return myComplexTypeMiejscowosc{X: 0, Y: 0}, err
	}
	resp := soapRequest(req)
	loc := parseLocation(resp)
	return loc, nil
}

// myComplexTypeBurza sends
func (c *client) MyComplexTypeBurza(x, y float64, r int) (myComplexTypeBurza, error) {
	return c.AuthMyComplexTypeBurza(x, y, r, c.APIKey())
}

// myComplexTypeBurza
func (c *client) AuthMyComplexTypeBurza(x, y float64, r int, apikey string) (myComplexTypeBurza, error) {
	p := myComplexTypeBurzaReqParams{apikey: apikey, x: x, y: y, radius: r}
	pm := make(map[string]string)
	p.print(&pm)
	req, err := renderMyComplexTypeBurzaRequest(pm)
	if err != nil {
		return myComplexTypeBurza{}, err
	}
	resp := soapRequest(req)
	MyComplexTypeBurza := parseMyComplexTypeBurza(resp)
	return MyComplexTypeBurza, nil
}

// myComplexTypeBurza sends
func (c *client) MyComplexTypeOstrzezenia(x, y float64) (myComplexTypeOstrzezenia, error) {
	return c.AuthMyComplexTypeOstrzezenia(x, y, c.APIKey())
}

// myComplexTypeBurza
func (c *client) AuthMyComplexTypeOstrzezenia(x, y float64, apikey string) (myComplexTypeOstrzezenia, error) {
	p := myComplexTypeOstrzezeniaReqParams{apikey: apikey, x: x, y: y}
	pm := make(map[string]string)
	p.print(&pm)
	req, err := renderMyComplexTypeOstrzezeniaRequest(pm)
	if err != nil {
		return myComplexTypeOstrzezenia{}, err
	}
	resp := soapRequest(req)
	MyComplexTypeOstrzezenia := parseMyComplexTypeOstrzezenia(resp)
	return MyComplexTypeOstrzezenia, nil
}
