package burzedzisnet

import (
	"bytes"
	"text/template"
)

func renderAuthRequest(params map[string]string) ([]byte, error) {
	return renderSoapRequest(`<?xml version="1.0" encoding="UTF-8" standalone="no"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="https://burze.dzis.net/soap.php" xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/" xmlns:soap-enc="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" ><SOAP-ENV:Body><mns:KeyAPI xmlns:mns="https://burze.dzis.net/soap.php" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><klucz xsi:type="xsd:string">{{.apikey}}</klucz></mns:KeyAPI></SOAP-ENV:Body></SOAP-ENV:Envelope>`, params)
}

func renderLocationRequest(params map[string]string) ([]byte, error) {
	return renderSoapRequest(`<?xml version="1.0" encoding="UTF-8" standalone="no"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="https://burze.dzis.net/soap.php" xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/" xmlns:soap-enc="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" ><SOAP-ENV:Body><mns:miejscowosc xmlns:mns="https://burze.dzis.net/soap.php" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><nazwa xsi:type="xsd:string">{{.name}}</nazwa><klucz xsi:type="xsd:string">{{.apikey}}</klucz></mns:miejscowosc></SOAP-ENV:Body></SOAP-ENV:Envelope>`, params)
}

func renderMyComplexTypeOstrzezeniaRequest(params map[string]string) ([]byte, error) {
	return renderSoapRequest(`<?xml version="1.0" encoding="UTF-8" standalone="no"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="https://burze.dzis.net/soap.php" xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/" xmlns:soap-enc="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" ><SOAP-ENV:Body><mns:ostrzezenia_pogodowe xmlns:mns="https://burze.dzis.net/soap.php" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><y xsi:type="xsd:float">{{.y}}</y><x xsi:type="xsd:float">{{.x}}</x><klucz xsi:type="xsd:string">{{.apikey}}</klucz></mns:ostrzezenia_pogodowe></SOAP-ENV:Body></SOAP-ENV:Envelope>`, params)
}

func renderMyComplexTypeBurzaRequest(params map[string]string) ([]byte, error) {
	return renderSoapRequest(`<?xml version="1.0" encoding="UTF-8" standalone="no"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="https://burze.dzis.net/soap.php" xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/" xmlns:soap-enc="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" ><SOAP-ENV:Body><mns:szukaj_burzy xmlns:mns="https://burze.dzis.net/soap.php" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><y xsi:type="xsd:string">{{.y}}</y><x xsi:type="xsd:string">{{.x}}</x><promien xsi:type="xsd:int">{{.radius}}</promien><klucz xsi:type="xsd:string">{{.apikey}}</klucz></mns:szukaj_burzy></SOAP-ENV:Body></SOAP-ENV:Envelope>`, params)
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
