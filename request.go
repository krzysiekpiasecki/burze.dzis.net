package burzedzisnet

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
)

type transf interface {
	transf(m *map[string]string)
}

func ptransf(t transf) map[string]string {
	m := make(map[string]string)
	t.transf(&m)
	return m
}

// authReqParams represents
type authReqParams struct {
	apikey string
}

func (p authReqParams) transf(m *map[string]string) {
	(*m)["apiKey"] = p.apikey
}

// myComplexTypeMiejscowoscReqParams represents
type myComplexTypeMiejscowoscReqParams struct {
	apikey string
	name   string
}

func (p myComplexTypeMiejscowoscReqParams) transf(m *map[string]string) {
	(*m)["apiKey"] = p.apikey
	(*m)["name"] = p.name
}

// myComplexTypeBurzaReqParams represents
type myComplexTypeBurzaReqParams struct {
	x      float64
	y      float64
	radius int
	apikey string
}

func (p myComplexTypeBurzaReqParams) transf(m *map[string]string) {
	(*m)["apiKey"] = p.apikey
	(*m)["x"] = strconv.FormatFloat(p.x, 'f', 2, 64)
	(*m)["y"] = strconv.FormatFloat(p.y, 'f', 2, 64)
	(*m)["radius"] = strconv.FormatInt(int64(p.radius), 10)
}

// myComplexTypeOstrzezeniaReqParams represents
type myComplexTypeOstrzezeniaReqParams struct {
	x      float64
	y      float64
	apikey string
}

func (p myComplexTypeOstrzezeniaReqParams) transf(m *map[string]string) {
	(*m)["apiKey"] = p.apikey
	(*m)["x"] = strconv.FormatFloat(p.x, 'f', 2, 64)
	(*m)["y"] = strconv.FormatFloat(p.y, 'f', 2, 64)
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

