package repository

import (
	"version1-medium-gotth/backend/model"
)

type DataManager struct {
}

func NewDataManager() *DataManager {
	return &DataManager{}
}

func (dm *DataManager) GetSampleData() model.PageData {
	return model.PageData{
		Title:       "Ready to transform your challenges into changes?",
		Subtitle:    "Get in touch with us today and start your transformation journey.",
		Description: "Our global team is ready to support your business transformation needs with innovative solutions and expert guidance.",
		Locations: []model.Location{
			{
				Name:         "Ireland Head Office",
				Image:        "https://www.version1.com/wp-content/uploads/2024/05/Ireland-1-scaled-640x427.jpg",
				Address:      "Millennium House,\nMillennium Walkway,\nDublin 1, D01 F5P8, Ireland",
				Country:      "Ireland",
				Region:       "Europe",
				IsHeadOffice: true,
			},
			{
				Name:         "UK Head Office",
				Image:        "https://www.version1.com/wp-content/uploads/2024/05/UK-1-scaled-640x427.jpg",
				Address:      "Waterloo Court, Suites 6 & 7,\nFirst Floor, 10 Theed St,\nLondon SE1 8ST, United Kingdom",
				Country:      "United Kingdom",
				Region:       "Europe",
				IsHeadOffice: true,
			},
			{
				Name:    "Malaga, Spain",
				Image:   "https://www.version1.com/wp-content/uploads/2024/05/Spain-1-640x427.jpg",
				Address: "2nd Floor Avda Ortega y Gasset 283,\nCP29006, Malaga, Spain",
				Country: "Spain",
				Region:  "Europe",
			},
			{
				Name:    "Ljubljana, Slovenia",
				Image:   "https://www.version1.com/wp-content/uploads/2024/06/Slovenia-640x340.jpg",
				Address: "Motnica 7,\n1236 Trzin, Slovenia",
				Country: "Slovenia",
				Region:  "Europe",
			},
			{
				Name:    "Belfast, Northern Ireland",
				Image:   "https://www.version1.com/wp-content/uploads/2024/05/Belfast-1-1-640x436.jpg",
				Address: "11th Floor Lanyon Plaza, West Tower,\n8 Lanyon Place, Belfast, BT1 3LP,\nNorthern Ireland",
				Country: "Northern Ireland",
				Region:  "Europe",
			},
			{
				Name:    "Birmingham, UK",
				Image:   "https://www.version1.com/wp-content/uploads/2024/05/Birmingham-640x360.png",
				Address: "Suite 3 D&E, Third Floor,\n31 Temple Street, Birmingham,\nB2 5DB, United Kingdom",
				Country: "United Kingdom",
				Region:  "Europe",
			},
			{
				Name:    "Cork, Ireland",
				Image:   "https://www.version1.com/wp-content/uploads/2024/05/Cork-Ireland-640x384.jpg",
				Address: "5 Lapp's Quay, Centre, Cork,\nT12 RW7D, Ireland",
				Country: "Ireland",
				Region:  "Europe",
			},
			{
				Name:    "Edinburgh, UK",
				Image:   "https://www.version1.com/wp-content/uploads/2024/05/Edinburgh-scaled-640x427.jpg",
				Address: "6 Waterloo Place, Edinburgh,\nEH1 3EG, United Kingdom",
				Country: "United Kingdom",
				Region:  "Europe",
			},
			{
				Name:    "Newcastle, UK",
				Image:   "https://www.version1.com/wp-content/uploads/2024/05/Newcastle-UK-640x366.jpg",
				Address: "Bank House, Pilgrim St,\nNewcastle-upon-Tyne NE1 6QE",
				Country: "United Kingdom",
				Region:  "Europe",
			},
			{
				Name:         "United States Head Office",
				Image:        "https://www.version1.com/wp-content/uploads/2024/05/New-York-640x427.png",
				Address:      "9th Floor, Office 9019, WeWork,\n1460 Broadway, New York 10036, USA",
				Country:      "United States",
				Region:       "Americas",
				IsHeadOffice: true,
			},
			{
				Name:    "Illinois, United States",
				Image:   "https://www.version1.com/wp-content/uploads/2024/05/Illinois-USA-1-640x427.jpg",
				Address: "15255 S. 94th Avenue, Suite 500,\nOrland Park, IL 60462, USA",
				Country: "United States",
				Region:  "Americas",
			},
			{
				Name:    "Sydney, Australia",
				Image:   "https://www.version1.com/wp-content/uploads/2024/05/Australia-scaled-640x427.jpg",
				Address: "Suite 8.02, Level 8, 9 Help Street,\nChatswood, New South Wales, 2067, Australia",
				Country: "Australia",
				Region:  "Asia Pacific",
			},
			{
				Name:    "Bangalore, India",
				Image:   "https://www.version1.com/wp-content/uploads/2024/05/India-Bangalore-scaled-640x321.jpeg",
				Address: "Madhuvan North Avenue – M2 Block, A Wing, 8th floor,\nManyata Embassy Business Park\nBengaluru, Karnataka 560045",
				Country: "India",
				Region:  "Asia Pacific",
			},
			{
				Name:    "Pune, India",
				Image:   "https://www.version1.com/wp-content/uploads/2025/04/pune-1-640x360.jpg",
				Address: "Baner One, Floor 701 & 702,\nPancard Club Road, Baner,\nPune, Maharashtra, 411045",
				Country: "India",
				Region:  "Asia Pacific",
			},
		},
	}
}
