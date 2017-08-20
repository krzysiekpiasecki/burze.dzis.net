package burzedzisnet

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
)

func makePanic(xmlDecoder *xml.Decoder) {
	var faultCode, faultString, last string
	for {
		tok, err := xmlDecoder.Token()
		if tok == nil {
			break
		}
		if err != nil {
			fmt.Printf("Parsing SOAP response failure %s", err)
			os.Exit(1)
		}
		switch se := tok.(type) {
		case xml.StartElement:
			{
				last = se.Name.Local
			}
		case xml.CharData:
			{
				if last == "faultcode" {
					last = ""
					faultCode = string(se)
				}
				if last == "faultstring" {
					last = ""
					faultString = string(se)
					break
				}
			}
		}
	}

	panic(fmt.Sprintf("SOAP ERROR with code %s and message %s", faultCode, faultString))
}

func parseAuth(response []byte) bool {

	xmlDecoder := xml.NewDecoder(bytes.NewReader(response))

	last := ""

	for {
		tok, err := xmlDecoder.Token()
		if tok == nil {
			break
		}
		if err != nil {
			fmt.Printf("Parsing SOAP response failure %s", err)
			os.Exit(1)
		}

		switch se := tok.(type) {
		case xml.StartElement:
			{
				last = se.Name.Local
			}
		case xml.CharData:
			{
				if last == "return" {
					last = ""
					r, err := strconv.ParseBool(string(se))
					if err != nil {
						fmt.Println(err, se)
						os.Exit(1)
					}
					return r
				}
			}
		}
	}

	panic("Parsing SOAP response failure")
}

func parseLocation(response []byte) myComplexTypeMiejscowosc {

	xmlDecoder := xml.NewDecoder(bytes.NewReader(response))

	last := ""
	loc := myComplexTypeMiejscowosc{}

	for {
		tok, err := xmlDecoder.Token()
		if tok == nil {
			break
		}
		if err != nil {
			fmt.Printf("Parsing SOAP response failure %s", err)
			os.Exit(1)
		}

		switch se := tok.(type) {
		case xml.StartElement:
			{
				if se.Name.Local == "Fault" {
					makePanic(xmlDecoder)
				}
				last = se.Name.Local
				if se.Name.Local == "x" {
					//			fmt.Print(se)
				} else if se.Name.Local == "y" {
					//			fmt.Print(se)
				}
			}
		case xml.CharData:
			{
				if last == "x" {
					last = ""
					loc.X, err = strconv.ParseFloat(string(se), 64)
					if err != nil {
						panic("Invalid externa data when parsing x myComplexTypeMiejscowosc")
					}
				}
				if last == "y" {
					last = ""
					loc.Y, err = strconv.ParseFloat(string(se), 64)
					if err != nil {
						panic("Invalid externa data when parsing y myComplexTypeMiejscowosc")
					}
				}

			}
		}
	}

	return loc

}

func parseMyComplexTypeBurza(response []byte) myComplexTypeBurza {

	xmlDecoder := xml.NewDecoder(bytes.NewReader(response))

	last := ""
	MyComplexTypeBurza := myComplexTypeBurza{}

	for {
		tok, err := xmlDecoder.Token()
		if tok == nil {
			break
		}
		if err != nil {
			fmt.Printf("Parsing SOAP response failure %s", err)
			os.Exit(1)
		}

		switch se := tok.(type) {
		case xml.StartElement:
			{
				if se.Name.Local == "Fault" {
					makePanic(xmlDecoder)
				}
				last = se.Name.Local
				if se.Name.Local == "x" {
					//			fmt.Print(se)
				} else if se.Name.Local == "y" {
					//			fmt.Print(se)
				}
			}
		case xml.CharData:
			{
				switch last {
				case "odleglosc":
					{
						MyComplexTypeBurza.Odleglosc, err = strconv.ParseFloat(string(se), 0)
					}
				case "liczba":
					{
						MyComplexTypeBurza.Liczba, err = strconv.Atoi(string(se))
					}
				case "okres":
					{
						MyComplexTypeBurza.Okres, err = strconv.Atoi(string(se))
					}
				case "kierunek":
					{
						MyComplexTypeBurza.Kierunek = string(se)
					}
				}

				last = ""
				if err != nil {
					panic("Invalid external data when parsing x myComplexTypeMiejscowosc")
				}
			}
		}
	}

	return MyComplexTypeBurza

}

func parseMyComplexTypeOstrzezenia(response []byte) myComplexTypeOstrzezenia {

	xmlDecoder := xml.NewDecoder(bytes.NewReader(response))

	last := ""
	MyComplexTypeOstrzezenia := myComplexTypeOstrzezenia{}

	for {
		tok, err := xmlDecoder.Token()
		if tok == nil {
			break
		}
		if err != nil {
			fmt.Printf("Parsing SOAP response failure %s", err)
			os.Exit(1)
		}

		switch se := tok.(type) {
		case xml.StartElement:
			{
				if se.Name.Local == "Fault" {
					makePanic(xmlDecoder)
				}
				last = se.Name.Local
			}
		case xml.CharData:
			{
				var err error

				switch last {
				case "odDnia":
					{
						MyComplexTypeOstrzezenia.OdDnia = string(se)
					}
				case "doDnia":
					{
						MyComplexTypeOstrzezenia.DoDnia = string(se)
					}
				case "mroz":
					{
						MyComplexTypeOstrzezenia.Mroz, err = strconv.Atoi(string(se))
					}
				case "mrozOdDnia":
					{
						MyComplexTypeOstrzezenia.MrozOdDnia = string(se)
					}
				case "mrozDoDnia":
					{
						MyComplexTypeOstrzezenia.MrozDoDnia = string(se)
					}
				case "upal":
					{
						MyComplexTypeOstrzezenia.Upal, err = strconv.Atoi(string(se))
					}
				case "upalOdDnia":
					{
						MyComplexTypeOstrzezenia.UpalOdDnia = string(se)
					}
				case "upalDoDnia":
					{
						MyComplexTypeOstrzezenia.UpalDoDnia = string(se)
					}
				case "wiatr":
					{
						MyComplexTypeOstrzezenia.Wiatr, err = strconv.Atoi(string(se))
					}
				case "wiatrOdDnia":
					{
						MyComplexTypeOstrzezenia.WiatrOdDnia = string(se)
					}
				case "wiatrDoDnia":
					{
						MyComplexTypeOstrzezenia.WiatrDoDnia = string(se)
					}
				case "opad":
					{
						MyComplexTypeOstrzezenia.Opad, err = strconv.Atoi(string(se))
					}
				case "opadOdDnia":
					{
						MyComplexTypeOstrzezenia.OpadOdDnia = string(se)
					}
				case "opadDoDnia":
					{
						MyComplexTypeOstrzezenia.OpadDoDnia = string(se)
					}
				case "burza":
					{
						MyComplexTypeOstrzezenia.Burza, err = strconv.Atoi(string(se))
					}
				case "burzaOdDnia":
					{
						MyComplexTypeOstrzezenia.BurzaOdDnia = string(se)
					}
				case "burzaDoDnia":
					{
						MyComplexTypeOstrzezenia.BurzaDoDnia = string(se)
					}
				case "traba":
					{
						MyComplexTypeOstrzezenia.Traba, err = strconv.Atoi(string(se))
					}
				case "trabaOdDnia":
					{
						MyComplexTypeOstrzezenia.TrabaOdDnia = string(se)
					}
				case "trabaDoDnia":
					{
						MyComplexTypeOstrzezenia.TrabaDoDnia = string(se)
					}
				}

				last = ""
				if err != nil {
					panic("Invalid external data when parsing x myComplexTypeMiejscowosc")
				}
			}
		}

	}

	return MyComplexTypeOstrzezenia
}
