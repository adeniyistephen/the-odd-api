package business

type SportKey string

type SportTitle string

type Sport struct {
	Key           SportKey   `json:"key"`
	Active        bool       `json:"active"`
	Group         string     `json:"group"`
	Description   string     `json:"description"`
	Title         SportTitle `json:"title"`
	Has_Outrights bool       `json:"has_outrights"`
}