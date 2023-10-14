package web

type CreateProductRequest struct {
	ProductName string `json:"product_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Stock       string `json:"stock" binding:"required"`
	Price       string `json:"price" binding:"required"`
	ImageUrl    string `json:"image_url"`
}