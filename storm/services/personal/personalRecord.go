package personal

// 개인기록 비즈니스 로직만 처리하는 서비스

import (
	"fmt"
	"log"
	"net/http"
	"storm/models"
	"storm/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewPersonalRecordServices(r *repository.PersonalRecordRepository) (*PersonalRecordServices, error) {
	return &PersonalRecordServices{repo: r}, nil
}

type PersonalRecordServices struct {
	repo *repository.PersonalRecordRepository
}

func Hello(ctx *fiber.Ctx) error {
	return ctx.SendString("storm personal record")
}

func (s *PersonalRecordServices) GetRecordIndex2() (*[]models.PersonalRecord, error) {
	personalRecords, err := s.repo.GetIndex()
	if err != nil {
		return nil, err
	}
	return personalRecords, nil
}

func (s *PersonalRecordServices) GetRecordByID(id string) (*models.PersonalRecord, error) {

	personalRecord, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return personalRecord, nil
}

func (s *PersonalRecordServices) UpdateRecord(personalRecord *models.PersonalRecord) (*models.PersonalRecord, error) {

	personalRecordNew, err := s.repo.UpdateRecord(personalRecord)

	if err != nil {
		return nil, err
	}

	return personalRecordNew, nil
}

func (s *PersonalRecordServices) DeleteRecord(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	log.Println("param id :" + id)
	if id == "" {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := s.repo.DeleteByID(id)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "cannot delete record",
		})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": fmt.Sprintf("record id %s delete successfully", id),
	})

	return nil
}

func (s *PersonalRecordServices) CreateRecord(ctx *fiber.Ctx) error {
	personalRecord := &models.PersonalRecord{}

	err := ctx.BodyParser(&personalRecord)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	log.Printf("data : %v", personalRecord)

	validator := validator.New()
	err = validator.Struct(personalRecord)

	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": err},
		)
		return err
	}

	personalRecordNew, err := s.repo.Create(personalRecord)

	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create personalRecord"})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "personalRecord Create has been successfully",
		"data":    personalRecordNew,
	})

	return nil
}

/*
func (s *PersonalRecordServices) SetupRoutes(app *fiber.App) {
	api := app.Group("/api/personal")
	api.Get("/records", s.GetRecordIndex)
	api.Get("/record/:id", s.GetRecordByID)
	api.Post("/record", s.CreateRecord)
	api.Put("/record", s.UpdateRecord)
	api.Delete("/record/:id", s.DeleteRecord)
}
*/
