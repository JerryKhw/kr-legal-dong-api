package model

type Si struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type Gu struct {
	Code   string `json:"code"`
	SiCode string `json:"siCode"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type Dong struct {
	Code   string `json:"code"`
	GuCode string `json:"guCode"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type Detail struct {
	Code     string `json:"code"`
	DongCode string `json:"dongCode"`
	Name     string `json:"name"`
	Active   bool   `json:"active"`
}
