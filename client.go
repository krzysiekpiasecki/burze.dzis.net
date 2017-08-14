package burzedzisnet

import (
	"os"
)

func apikeyFromEnv() string {
	return os.Getenv("BURZEDZISNET_apikey")
}

// Client represents a
type Client struct {
	apikey string
}

// NewClient returns
func NewClient(apikey string) *Client {
	return &Client{apikey: apikey}
}

// APIKey returns used API key by the client.
func (c *Client) APIKey() string {
	k := (*c).apikey
	if k == "" {
		return apikeyFromEnv()
	}
	return k
}

// Auth sends a client key and returns true, when the key is valid.
//
// Otherwise it returns false
func (c *Client) Auth() (bool, error) {
	p := AuthReqParams{apikey: c.APIKey()}
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
func (c *Client) AuthKey(apikey string) (bool, error) {
	p := AuthReqParams{apikey: apikey}
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
func (c *Client) Locate(name string) (MyComplexTypeMiejscowosc, error) {
	return c.AuthLocate(name, c.APIKey())
}

// AuthLocate sends
// TODO: Handle variadic args to support apikey arg
func (c *Client) AuthLocate(name string, apikey string) (MyComplexTypeMiejscowosc, error) {
	p := LocationReqParams{apikey: apikey, name: name}
	pm := make(map[string]string)
	p.print(&pm)
	req, err := renderLocationRequest(pm)
	if err != nil {
		return MyComplexTypeMiejscowosc{X: 0, Y: 0}, err
	}
	resp := soapRequest(req)
	loc := parseLocation(resp)
	return loc, nil
}

// MyComplexTypeBurza sends
func (c *Client) MyComplexTypeBurza(x, y float64, r int) (MyComplexTypeBurza, error) {
	return c.AuthMyComplexTypeBurza(x, y, r, c.APIKey())
}

// MyComplexTypeBurza
func (c *Client) AuthMyComplexTypeBurza(x, y float64, r int, apikey string) (MyComplexTypeBurza, error) {
	p := MyComplexTypeBurzaReqParams{apikey: apikey, x: x, y: y, radius: r}
	pm := make(map[string]string)
	p.print(&pm)
	req, err := renderMyComplexTypeBurzaRequest(pm)
	if err != nil {
		return MyComplexTypeBurza{}, err
	}
	resp := soapRequest(req)
	MyComplexTypeBurza := parseMyComplexTypeBurza(resp)
	return MyComplexTypeBurza, nil
}

// MyComplexTypeBurza sends
func (c *Client) MyComplexTypeOstrzezenia(x, y float64) (MyComplexTypeOstrzezenia, error) {
	return c.AuthMyComplexTypeOstrzezenia(x, y, c.APIKey())
}

// MyComplexTypeBurza
func (c *Client) AuthMyComplexTypeOstrzezenia(x, y float64, apikey string) (MyComplexTypeOstrzezenia, error) {
	p := MyComplexTypeOstrzezeniaReqParams{apikey: apikey, x: x, y: y}
	pm := make(map[string]string)
	p.print(&pm)
	req, err := renderMyComplexTypeOstrzezeniaRequest(pm)
	if err != nil {
		return MyComplexTypeOstrzezenia{}, err
	}
	resp := soapRequest(req)
	MyComplexTypeOstrzezenia := parseMyComplexTypeOstrzezenia(resp)
	return MyComplexTypeOstrzezenia, nil
}
