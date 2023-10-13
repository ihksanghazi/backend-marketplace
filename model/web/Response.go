package web

type BasicResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Pagination struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	CurrentPage string      `json:"current_page"`
	TotalPage   int         `json:"total_page"`
	Data        interface{} `json:"data"`
}