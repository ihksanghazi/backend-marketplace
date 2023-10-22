package web

type CheckoutRequest struct {
	OriginCity       string `json:"origin_city" binding:"required"`
	DestionationCity string `json:"destination_city" binding:"required"`
	Courier          string `json:"courier" binding:"required"`
	WeightOnGram     string `json:"weight_on_gram" binding:"required,number"`
	Service          string `json:"service" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Price            int    `json:"price" binding:"required,number"`
}
