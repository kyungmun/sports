package team

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TeamRecord struct {
	ID      string `gorm:"column:m_id;" json:"m_id"`           // validate:"required"`
	Check   string `gorm:"column:m_check;" json:"m_check"`     // validate:"required"`
	DestDay string `gorm:"column:m_destday;" json:"m_destday"` // validate:"required"`
	Name    string `gorm:"column:m_name;" json:"m_name"`       // validate:"required"`
	Date    string `gorm:"column:m_date;" json:"m_date"`       // validate:"required"`
	Result  string `gorm:"column:m_result;" json:"m_result"`   // validate:"required"`
	Mvp     int    `gorm:"column:m_mvp;" json:"m_mvp"`         // validate:"required"`
	Point   int    `gorm:"column:m_point;" json:"m_point"`     // validate:"required"`
	Gole    int    `gorm:"column:m_gole;" json:"m_gole"`       // validate:"required"`
	Assist  int    `gorm:"column:m_assist;" json:"m_assist"`   // validate:"required"`
	Mobile  int    `gorm:"column:m_mobile;" json:"m_mobile"`   // validate:"required"`
}

type Repository struct {
	DB *gorm.DB
}

func Hello(context *fiber.Ctx) error {
	return context.SendString("storm personal record")
}

func (r *Repository) GetRecordByID(context *fiber.Ctx) error {
	teamRecord := &[]TeamRecord{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Table("xe_member_attend_check").Where("m_id = ?", id).Order("m_destday desc").Find(teamRecord).Error
	if (err != nil) || (len(*teamRecord) == 0) {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get record"})
		return err
	}

	//log.Println(*personalRecords)
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "record id gotten successfully",
		"data":    teamRecord,
		"count":   len(*teamRecord),
	})
	return nil
}

func (r *Repository) CreateRecord(context *fiber.Ctx) error {
	teamRecord := TeamRecord{}

	err := context.BodyParser(&teamRecord)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	validator := validator.New()
	err = validator.Struct(TeamRecord{})

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": err},
		)
		return err
	}

	err = r.DB.Table("xe_member_attend_check").Create(&teamRecord).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create personalRecord"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "personalRecord has been successfully added",
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api/team")
	api.Post("/record", r.CreateRecord)
	api.Get("/records/:id", r.GetRecordByID)
}
