package burzedzisnet

// myComplexTypeMiejscowosc represents
type myComplexTypeMiejscowosc struct {
	X float64
	Y float64
}

// IsSpec ensures that
func (c myComplexTypeMiejscowosc) IsSpec() bool {
	return c.X != 0 || c.Y != 0
}

// myComplexTypeBurza represents
type myComplexTypeBurza struct {
	Liczba    int
	Odleglosc float64
	Kierunek  string
	Okres     int
}

// myComplexTypeOstrzezenia represents
type myComplexTypeOstrzezenia struct {
	OdDnia      string
	DoDnia      string
	Mroz        int
	MrozOdDnia  string
	MrozDoDnia  string
	Upal        int
	UpalOdDnia  string
	UpalDoDnia  string
	Wiatr       int
	WiatrOdDnia string
	WiatrDoDnia string
	Opad        int
	OpadOdDnia  string
	OpadDoDnia  string
	Burza       int
	BurzaOdDnia string
	BurzaDoDnia string
	Traba       int
	TrabaOdDnia string
	TrabaDoDnia string
}

// IsSafe returns
func (w myComplexTypeOstrzezenia) IsSafe() bool {
	return w.Mroz+w.Upal+w.Wiatr+w.Opad+w.Burza+w.Traba == 0
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

// MyComplexTypeOstrzezenia sends
func (c *client) LookupNames(name, country string) ([]string, error) {
	req, err := renderListaMiejscowosciRequest(myComplexTypeMiejscowoscListReqParams{apikey: c.apiKey, name: name, country: country}.transf())
	if err != nil {
		return make([]string, 0), err
	}

	return parseLookupNameResponse(doSoapRequest(req)), nil
}
