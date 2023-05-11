package productsdto

import "time"

type ProductResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" gorm:"type: varchar(255)"  validate:"required"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
