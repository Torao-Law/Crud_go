package productsdto

type CreateProductRequest struct {
	Name string `json:"name" form:"name" gorm:"type: varchar(255)"`
}

type UpdateProductRequest struct {
	Name string `json:"name" form:"name"`
}
