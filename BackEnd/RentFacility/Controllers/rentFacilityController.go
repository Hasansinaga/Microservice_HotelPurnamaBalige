package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	database "rentFacility/Database"
	"rentFacility/Models/entity"
	"rentFacility/Models/request"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func getUserID(token string, UserID int) (*entity.User, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://127.0.0.1:8084/user/profile/%d", UserID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	req.AddCookie(&http.Cookie{Name: "jwt", Value: token})

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request user failed with status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var user entity.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return &user, nil
}

func getFacilityID(facilityID int) (*entity.Facility, error) {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:8083/facility/%d", facilityID))
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request room failed with status code: %d", resp.StatusCode)
	}

	var facility entity.Facility
	if err := json.NewDecoder(resp.Body).Decode(&facility); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return &facility, nil
}

func GetAllRentFacility(c *fiber.Ctx) error {
	var rentfacility []entity.RentFacility

	result := database.DB.Find(&rentfacility)

	if len(rentfacility) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Not Found",
		})
	}

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Not Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": rentfacility,
	})
}

func GetRentFacilityByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var rentfacility entity.RentFacility

	err := database.DB.First(&rentfacility, "id = ?", id).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "Room Not Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": rentfacility,
	})
}

func CreateRentFacility(c *fiber.Ctx) error {
	input := new(request.RequestRentFacilityCreate)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "Failed",
			"message": err.Error(),
		})
	}

	FacilityID, err := strconv.Atoi(c.FormValue("facility_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "Facility ID is required and must be an integer",
		})
	}

	facility, err := getFacilityID(FacilityID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	facility.Id = uint(FacilityID)

	UserID, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "UserID is required and must be an integer",
		})
	}

	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user, err := getUserID(cookie, UserID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	user.Id = uint(UserID)

	checkIn, err := time.Parse("2006-01-02", input.CheckIn)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "CheckIn date is invalid. Format should be YYYY-MM-DD",
		})
	}

	// Membuat objek RoomBooking
	roombooking := entity.RentFacility{
		CheckIn:      checkIn.Format("2006-01-02"),
		CheckInTime:  input.CheckInTime,
		CheckOutTime: input.CheckOutTime,
		Status:       "waiting",
		UserID:       uint(UserID),
		FacilityID:   uint(FacilityID),
	}

	// Menyimpan data ke database
	if err := database.DB.Create(&roombooking).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to create room booking",
		})
	}

	// Mengembalikan respons sukses
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": roombooking,
	})
}

func UpdateRoomBookingByUser(c *fiber.Ctx) error {
	id := c.Params("id")
	input := new(request.RequestRentFacilityUpdate)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	var booking entity.RentFacility

	err := database.DB.First(&booking, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "Booking not found",
		})
	}

	if input.CheckIn != "" {
		checkIn, err := time.Parse("2006-01-02", input.CheckIn)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "failed",
				"message": "Invalid CheckIn date. Format should be YYYY-MM-DD",
			})
		}
		booking.CheckIn = checkIn.Format("2006-01-02")
	}

	facilityID, err := strconv.Atoi(c.FormValue("facility_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "facilityID is required and must be an integer",
		})
	}

	if facilityID != 0 {
		room, err := getFacilityID(facilityID)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "failed",
				"message": err.Error(),
			})
		}
		room.Id = uint(facilityID)
		booking.FacilityID = uint(facilityID)
	}

	result := database.DB.Save(&booking)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't update booking",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": booking,
	})
}

func DeleteRentFacility(c *fiber.Ctx) error {
	id := c.Params("id")

	var booking entity.RentFacility

	err := database.DB.Debug().First(&booking, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "failed",
			"message": "RoomBooking Not Found",
		})
	}

	if err := database.DB.Debug().Delete(&booking).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "Can't Delete RoomBooking",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "RoomBooking deleted successfully!",
	})
}
