package common

// conference & person bio data
type Info interface {
	bioData() map[string]string
}

// attendant
type Attendant struct {
	name       string
	conference string
}

// implement info for attendant
func (a Attendant) bioData() map[string]string {
	attendant := map[string]string{"name": a.name, "conference": a.conference}
	return attendant
}

// visitors
type Visitor struct {
	name       string
	conference string
	isPaid     string
}

// implement visitor bio data
func (v Visitor) bioData() map[string]string {
	visitor := map[string]string{"name": v.name, "conference": v.conference, "isPaid": v.isPaid}
	return visitor

}

// conference
type Conference struct {
	name     string
	location string
	price    string
	capacity string
}

// implement conference bio data

func (c Conference) bioData() map[string]string {
	conference := map[string]string{"name": c.name, "location": c.location, "price": c.price, "capacity": c.capacity}
	return conference

}
