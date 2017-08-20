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
