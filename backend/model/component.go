package model

type Location struct {
	Name         string
	Image        string
	Address      string
	Country      string
	Region       string
	IsHeadOffice bool
}

type PageData struct {
	Title       string
	Subtitle    string
	Description string
	Locations   []Location
}
