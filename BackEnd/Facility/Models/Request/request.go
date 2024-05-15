package request

type RequestFacilityCreate struct {
	Name         string  `json:"name" gorm:"type:varchar(50);index:idx_nm,unique" validate:"required"`
	Description  string  `json:"description" gorm:"type:text" validate:"required"`
	Rental_Price float64 `json:"rental_price" gorm:"type:int" validate:"required"`
}

type RequestFacilityUpdate struct {
	Name         string  `json:"name" gorm:"type:varchar(50);index:idx_nm,unique" validate:"required"`
	Description  string  `json:"description" gorm:"type:text" validate:"required"`
	Rental_Price float64 `json:"rental_price" gorm:"type:int" validate:"required"`
}
