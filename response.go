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

func parseLocation(response []byte) MyComplexTypeMiejscowosc {

	xmlDecoder := xml.NewDecoder(bytes.NewReader(response))

	last := ""
	loc := MyComplexTypeMiejscowosc{}

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
						panic("Invalid externa data when parsing x MyComplexTypeMiejscowosc")
					}
				}
				if last == "y" {
					last = ""
					loc.Y, err = strconv.ParseFloat(string(se), 64)
					if err != nil {
						panic("Invalid externa data when parsing y MyComplexTypeMiejscowosc")
					}
				}

			}
		}
	}

	return loc

}

func parseMyComplexTypeBurza(response []byte) MyComplexTypeBurza {

	xmlDecoder := xml.NewDecoder(bytes.NewReader(response))

	last := ""
	MyComplexTypeBurza := MyComplexTypeBurza{}

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
					panic("Invalid external data when parsing x MyComplexTypeMiejscowosc")
				}
			}
		}
	}

	return MyComplexTypeBurza

}

func parseMyComplexTypeOstrzezenia(response []byte) MyComplexTypeOstrzezenia {

	xmlDecoder := xml.NewDecoder(bytes.NewReader(response))

	last := ""
	MyComplexTypeOstrzezenia := MyComplexTypeOstrzezenia{}

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
				case "od_dnia":
					{
						MyComplexTypeOstrzezenia.Od_dnia = string(se)
					}
				case "do_dnia":
					{
						MyComplexTypeOstrzezenia.Do_dnia = string(se)
					}
				case "mroz":
					{
						MyComplexTypeOstrzezenia.Mroz, err = strconv.Atoi(string(se))
					}
				case "mroz_od_dnia":
					{
						MyComplexTypeOstrzezenia.Mroz_od_dnia = string(se)
					}
				case "mroz_do_dnia":
					{
						MyComplexTypeOstrzezenia.Mroz_do_dnia = string(se)
					}
				case "upal":
					{
						MyComplexTypeOstrzezenia.Upal, err = strconv.Atoi(string(se))
					}
				case "upal_od_dnia":
					{
						MyComplexTypeOstrzezenia.Upal_od_dnia = string(se)
					}
				case "upal_do_dnia":
					{
						MyComplexTypeOstrzezenia.Upal_do_dnia = string(se)
					}
				case "wiatr":
					{
						MyComplexTypeOstrzezenia.Wiatr, err = strconv.Atoi(string(se))
					}
				case "wiatr_od_dnia":
					{
						MyComplexTypeOstrzezenia.Wiatr_od_dnia = string(se)
					}
				case "wiatr_do_dnia":
					{
						MyComplexTypeOstrzezenia.Wiatr_do_dnia = string(se)
					}
				case "opad":
					{
						MyComplexTypeOstrzezenia.Opad, err = strconv.Atoi(string(se))
					}
				case "opad_od_dnia":
					{
						MyComplexTypeOstrzezenia.Opad_od_dnia = string(se)
					}
				case "opad_do_dnia":
					{
						MyComplexTypeOstrzezenia.Opad_do_dnia = string(se)
					}
				case "burza":
					{
						MyComplexTypeOstrzezenia.Burza, err = strconv.Atoi(string(se))
					}
				case "burza_od_dnia":
					{
						MyComplexTypeOstrzezenia.Burza_od_dnia = string(se)
					}
				case "burza_do_dnia":
					{
						MyComplexTypeOstrzezenia.Burza_do_dnia = string(se)
					}
				case "traba":
					{
						MyComplexTypeOstrzezenia.Traba, err = strconv.Atoi(string(se))
					}
				case "traba_od_dnia":
					{
						MyComplexTypeOstrzezenia.Traba_od_dnia = string(se)
					}
				case "traba_do_dnia":
					{
						MyComplexTypeOstrzezenia.Traba_do_dnia = string(se)
					}
				}

				last = ""
				if err != nil {
					panic("Invalid external data when parsing x MyComplexTypeMiejscowosc")
				}
			}
		}

	}

	return MyComplexTypeOstrzezenia
}
