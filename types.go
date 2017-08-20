package burzedzisnet

// MyComplexTypeMiejscowosc represents a pair of coordinates of a location (DMS).
type MyComplexTypeMiejscowosc struct {
	X float64
	Y float64
}

// IsSpec ensures that the coordinates of a location are specified.
func (c MyComplexTypeMiejscowosc) IsSpec() bool {
	return c.X != 0 || c.Y != 0
}

// MyComplexTypeBurza represents information about a number of lightnings
type MyComplexTypeBurza struct {
	Liczba    int
	Odleglosc float64
	Kierunek  string
	Okres     int
}

// MyComplexTypeOstrzezenia represents information about wheater MyComplexTypeOstrzezenia for the given location
type MyComplexTypeOstrzezenia struct {
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

// IsSafe returns true when no MyComplexTypeOstrzezenias were registered otherwise it resturns false
func (w MyComplexTypeOstrzezenia) IsSafe() bool {
	return 0 == w.Mroz+w.Upal+w.Wiatr+w.Opad+w.Burza+w.Traba
}
