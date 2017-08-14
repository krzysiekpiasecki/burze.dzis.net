package burzedzisnet

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
)

// AuthReqParams represents
type AuthReqParams struct {
	apikey string
}

// LocationReqParams represents
type LocationReqParams struct {
	apikey string
	name   string
}

// MyComplexTypeBurzaReqParams represents
type MyComplexTypeBurzaReqParams struct {
	x      float64
	y      float64
	radius int
	apikey string
}

// MyComplexTypeOstrzezeniaReqParams represents
type MyComplexTypeOstrzezeniaReqParams struct {
	x      float64
	y      float64
	apikey string
}

func (p AuthReqParams) print(m *map[string]string) {
	(*m)["apikey"] = p.apikey
}

func (p LocationReqParams) print(m *map[string]string) {
	(*m)["apikey"] = p.apikey
	(*m)["name"] = p.name
}

func (p MyComplexTypeBurzaReqParams) print(m *map[string]string) {
	(*m)["apikey"] = p.apikey
	(*m)["x"] = strconv.FormatFloat(p.x, 'f', 2, 64)
	(*m)["y"] = strconv.FormatFloat(p.y, 'f', 2, 64)
	(*m)["radius"] = strconv.FormatInt(int64(p.radius), 10)
}

func (p MyComplexTypeOstrzezeniaReqParams) print(m *map[string]string) {
	(*m)["apikey"] = p.apikey
	(*m)["x"] = strconv.FormatFloat(p.x, 'f', 2, 64)
	(*m)["y"] = strconv.FormatFloat(p.y, 'f', 2, 64)
}

func authRequest(p AuthReqParams) (*http.Request, error) {
	m := make(map[string]string)
	p.print(&m)
	r, _ := renderAuthRequest(m)
	return http.NewRequest("POST", "https://burze.dzis.net/soap.php", bytes.NewBuffer(r))
}

func locationRequest(p LocationReqParams) (*http.Request, error) {
	m := make(map[string]string)
	p.print(&m)
	r, _ := renderLocationRequest(m)
	return http.NewRequest("POST", "https://burze.dzis.net/soap.php", bytes.NewBuffer(r))
}

func MyComplexTypeBurzaRequest(p MyComplexTypeBurzaReqParams) (*http.Request, error) {
	m := make(map[string]string)
	p.print(&m)
	r, _ := renderMyComplexTypeBurzaRequest(m)
	return http.NewRequest("POST", "https://burze.dzis.net/soap.php", bytes.NewBuffer(r))
}

func MyComplexTypeOstrzezeniaRequest(p MyComplexTypeOstrzezeniaReqParams) (*http.Request, error) {
	m := make(map[string]string)
	p.print(&m)
	r, _ := renderMyComplexTypeOstrzezeniaRequest(m)
	return http.NewRequest("POST", "https://burze.dzis.net/soap.php", bytes.NewBuffer(r))
}

func soapRequest(soapReq []byte) []byte {
	resp, err := http.Post("https://burze.dzis.net/soap.php", "text/xml", bytes.NewBuffer(soapReq))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}
