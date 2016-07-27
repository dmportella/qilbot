package edsm

// System Simple representation of a system object as detailed by edsm
type System struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Coords *Coordinates `json:"coords"`
}

// Coordinates Simple representation of a coord object as detailed by edsm
type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
