package burzedzisnet

// MyComplexTypeMiejscowosc represents a pair of coordinates of a location (DMS).
type MyComplexTypeMiejscowosc struct {
	X float64
	Y float64
}

// IsSpec ensures that the coordinates of a location are specified.
func (c MyComplexTypeMiejscowosc) IsSpec() bool {
	return c.X != 0 && c.Y != 0
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
	Od_dnia       string
	Do_dnia       string
	Mroz          int
	Mroz_od_dnia  string
	Mroz_do_dnia  string
	Upal          int
	Upal_od_dnia  string
	Upal_do_dnia  string
	Wiatr         int
	Wiatr_od_dnia string
	Wiatr_do_dnia string
	Opad          int
	Opad_od_dnia  string
	Opad_do_dnia  string
	Burza         int
	Burza_od_dnia string
	Burza_do_dnia string
	Traba         int
	Traba_od_dnia string
	Traba_do_dnia string
}

// IsSafe returns true when no MyComplexTypeOstrzezenias were registered otherwise it resturns false
func (w MyComplexTypeOstrzezenia) IsSafe() bool {
	return 0 == w.Mroz+w.Upal+w.Opad+w.Burza+w.Traba
}
