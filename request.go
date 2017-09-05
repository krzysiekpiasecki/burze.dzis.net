package burzedzisnet

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
)

type transf interface {
	transf() map[string]string
}

// apiKeyReqParams represents
type apiKeyReqParams struct {
	apikey string
}

// myComplexTypeMiejscowoscReqParams represents
type myComplexTypeMiejscowoscReqParams struct {
	apikey string
	name   string
}

// myComplexTypeMiejscowoscListReqParams represents
type myComplexTypeMiejscowoscListReqParams struct {
	apikey  string
	name    string
	country string
}

// myComplexTypeOstrzezeniaReqParams represents
type myComplexTypeOstrzezeniaReqParams struct {
	apikey string
	x      float64
	y      float64
}

// myComplexTypeBurzaReqParams represents
type myComplexTypeBurzaReqParams struct {
	apikey string
	x      float64
	y      float64
	radius int
}

// myListaMiejscowosciParams represents
type myListaMiejscowosciParams struct {
	apikey  string
	name    string
	country string
}

func (p apiKeyReqParams) transf() (m map[string]string) {
	m = make(map[string]string)
	m["apiKey"] = p.apikey
	return
}

func (p myComplexTypeMiejscowoscReqParams) transf() (m map[string]string) {
	m = make(map[string]string)
	m["apiKey"] = p.apikey
	m["name"] = p.name
	return
}

func (p myComplexTypeBurzaReqParams) transf() (m map[string]string) {
	m = make(map[string]string)
	m["apiKey"] = p.apikey
	m["x"] = strconv.FormatFloat(p.x, 'f', 2, 64)
	m["y"] = strconv.FormatFloat(p.y, 'f', 2, 64)
	m["radius"] = strconv.FormatInt(int64(p.radius), 10)
	return
}

func (p myComplexTypeOstrzezeniaReqParams) transf() (m map[string]string) {
	m = make(map[string]string)
	m["apiKey"] = p.apikey
	m["x"] = strconv.FormatFloat(p.x, 'f', 2, 64)
	m["y"] = strconv.FormatFloat(p.y, 'f', 2, 64)
	return
}

func (p myComplexTypeMiejscowoscListReqParams) transf() (m map[string]string) {
	m = make(map[string]string)
	m["apiKey"] = p.apikey
	m["name"] = p.name
	m["country"] = p.country
	return
}

func doSoapRequest(soapReq []byte) []byte {
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
