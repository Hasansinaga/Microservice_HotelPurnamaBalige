package request

type RequestRentFacilityCreate struct {
	CheckIn      string `json:"checkin" gorm:"column:check_in" db:"check_in"`
	CheckInTime  string `json:"checkin_time" gorm:"column:checkin_time" db:"checkin_time" validate:"required"`
	CheckOutTime string `json:"checkout_time" gorm:"column:checkout_time" db:"checkout_time" validate:"required"`
}

type RequestRentFacilityUpdate struct {
	CheckIn      string `json:"checkin" gorm:"column:check_in" db:"check_in"`
	CheckInTime  string `json:"checkin_time" gorm:"column:checkin_time" db:"checkin_time" validate:"required"`
	CheckOutTime string `json:"checkout_time" gorm:"column:checkout_time" db:"checkout_time" validate:"required"`
}
