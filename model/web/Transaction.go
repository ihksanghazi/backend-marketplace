package web

type Ongkir struct {
	Rajaongkir rajaongkir `json:"rajaongkir"`
}

type rajaongkir struct {
	DestinationDetails city     `json:"destination_details"`
	OriginDetails      city     `json:"origin_details"`
	Query              query    `json:"query"`
	Results            []result `json:"results"`
	Status             status   `json:"status"`
}

type city struct {
	CityId     string `json:"city_id"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
	Province   string `json:"province"`
	ProvinceId string `json:"province_id"`
	Type       string `json:"type"`
}

type query struct {
	Courier     string `json:"courier"`
	Destination string `json:"destination"`
	Origin      string `json:"origin"`
	Weight      int    `json:"weight"`
}

type status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type result struct {
	Code  string  `json:"code"`
	Costs []costs `json:"costs"`
	Name  string  `json:"name"`
}

type costs struct {
	Cost        []cost `json:"cost"`
	Description string `json:"description"`
	Service     string `json:"service"`
}

type cost struct {
	Etd   string `json:"etd"`
	Note  string `json:"note"`
	Value int    `json:"value"`
}

type ExpeditionWebResponse struct {
	OriginDetails      city                `json:"origin_details"`
	DestinationDetails city                `json:"destination_details"`
	Services           []ServiceExpedition `json:"services"`
}

type ServiceExpedition struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Value       int    `json:"value"`
	Etd         string `json:"etd"`
	Note        string `json:"note"`
}
