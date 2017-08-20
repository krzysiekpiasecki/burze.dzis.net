package burzedzisnet

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
)

// authReqParams represents
type authReqParams struct {
	apikey string
}

// myComplexTypeMiejscowoscReqParams represents
type myComplexTypeMiejscowoscReqParams struct {
	apikey string
	name   string
}

// myComplexTypeBurzaReqParams represents
type myComplexTypeBurzaReqParams struct {
	x      float64
	y      float64
	radius int
	apikey string
}

// myComplexTypeOstrzezeniaReqParams represents
type myComplexTypeOstrzezeniaReqParams struct {
	x      float64
	y      float64
	apikey string
}

func (p authReqParams) print(m *map[string]string) {
	(*m)["apikey"] = p.apikey
}

func (p myComplexTypeMiejscowoscReqParams) print(m *map[string]string) {
	(*m)["apikey"] = p.apikey
	(*m)["name"] = p.name
}

func (p myComplexTypeBurzaReqParams) print(m *map[string]string) {
	(*m)["apikey"] = p.apikey
	(*m)["x"] = strconv.FormatFloat(p.x, 'f', 2, 64)
	(*m)["y"] = strconv.FormatFloat(p.y, 'f', 2, 64)
	(*m)["radius"] = strconv.FormatInt(int64(p.radius), 10)
}

func (p myComplexTypeOstrzezeniaReqParams) print(m *map[string]string) {
	(*m)["apikey"] = p.apikey
	(*m)["x"] = strconv.FormatFloat(p.x, 'f', 2, 64)
	(*m)["y"] = strconv.FormatFloat(p.y, 'f', 2, 64)
}

func authRequest(p authReqParams) (*http.Request, error) {
	m := make(map[string]string)
	p.print(&m)
	r, _ := renderAuthRequest(m)
	return http.NewRequest("POST", "https://burze.dzis.net/soap.php", bytes.NewBuffer(r))
}

func locationRequest(p myComplexTypeMiejscowoscReqParams) (*http.Request, error) {
	m := make(map[string]string)
	p.print(&m)
	r, _ := renderSoapRequest(`<?xml version="1.0" encoding="UTF-8" standalone="no"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="https://burze.dzis.net/soap.php" xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/" xmlns:soap-enc="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" ><SOAP-ENV:Body><mns:miejscowosc xmlns:mns="https://burze.dzis.net/soap.php" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><nazwa xsi:type="xsd:string">{{.name}}</nazwa><klucz xsi:type="xsd:string">{{.apikey}}</klucz></mns:miejscowosc></SOAP-ENV:Body></SOAP-ENV:Envelope>`, m)
	return http.NewRequest("POST", "https://burze.dzis.net/soap.php", bytes.NewBuffer(r))
}

func MyComplexTypeBurzaRequest(p myComplexTypeBurzaReqParams) (*http.Request, error) {
	m := make(map[string]string)
	p.print(&m)
	r, _ := renderMyComplexTypeBurzaRequest(m)
	return http.NewRequest("POST", "https://burze.dzis.net/soap.php", bytes.NewBuffer(r))
}

func MyComplexTypeOstrzezeniaRequest(p myComplexTypeOstrzezeniaReqParams) (*http.Request, error) {
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


func renderSoapRequest(req string, params map[string]string) ([]byte, error) {
	var b bytes.Buffer
	tpl := template.Must(template.New("request").Parse(req))
	err := tpl.Execute(&b, &params)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
