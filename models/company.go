package models

type Company struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	CreateAt     int64    `json:"create_at"`
	BuildingId   string   `json:"building_id"`
	Address      string   `json:"address"`
	PhoneNumbers []string `json:"phone_numbers"`
	Categories   []string `json:"categories"`
}

func (comp *Company) Reset() {
	comp.Id = ""
	comp.Name = ""
	comp.CreateAt = 0
	comp.BuildingId = ""
	comp.Address = ""
	if comp.PhoneNumbers != nil {
		comp.PhoneNumbers = comp.PhoneNumbers[:0]
	}
	if comp.Categories != nil {
		comp.Categories = comp.Categories[:0]
	}
}
