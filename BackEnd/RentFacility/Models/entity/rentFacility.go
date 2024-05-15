package entity

import "time"

type RentFacility struct {
	Id           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	FacilityID   uint      `json:"facility_id"`
	Status       string    `json:"status" gorm:"type:varchar(255)"`
	CheckIn      string    `json:"checkin" gorm:"column:check_in" db:"check_in"`
	CheckInTime  string    `json:"checkin_time"`
	CheckOutTime string    `json:"checkout_time"`
	Total        float64   `json:"total"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}

type Facility struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name" gorm:"type:varchar(50);index:idx_nm,unique" validate:"required"`
	Description  string    `json:"description" gorm:"type:text" validate:"required"`
	Rental_Price float64   `json:"rental_price" gorm:"type:int" validate:"required"`
	Image        string    `json:"image" gorm:"type:varchar(50)" validate:"required"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name" gorm:"type:varchar(255)" validate:"required"`
	Phone     string    `json:"phone" gorm:"type:varchar(255)" validate:"required"`
	Email     string    `json:"email" gorm:"type:varchar(255)" validate:"required"`
	Username  string    `json:"username" gorm:"type:varchar(255)" validate:"required,min=5"`
	Password  string    `json:"password" gorm:"type:varchar(255)" validate:"required,min=5"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}
