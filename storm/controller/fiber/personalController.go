package controller

//개인기록 요청받아 서비스에서 결과를 받아와 응답주는 처리

import (
	"fmt"
	"net/http"
	"storm/models"
	"storm/services/personal"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PersonalRecordController struct {
	svc *personal.PersonalRecordServices
}

func NewPersonalRecordController(s *personal.PersonalRecordServices) *PersonalRecordController {
	return &PersonalRecordController{svc: s}
}

func (c *PersonalRecordController) GetRecordIndex(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 0
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil {
		pageSize = 0
	}
	personalRecords, err := c.svc.GetRecordIndex(page, pageSize)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get data"})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "fiber engine, records all get successfully",
		"count":   len(*personalRecords),
		"data":    personalRecords,
	})
	return nil
}

func (c *PersonalRecordController) GetRecordByID(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	//log.Println("param id :" + id)
	if userId == "" {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	personalRecord, err := c.svc.GetRecordByID(userId)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get record",
		})
		return err
	}

	//log.Println(*personalRecords)
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "record id gotten successfully",
		"data":    personalRecord,
	})
	return nil
}

func (c *PersonalRecordController) UpdateRecord(ctx *fiber.Ctx) error {
	personalRecord := &models.PersonalRecord{}

	err := ctx.BodyParser(&personalRecord)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	fmt.Println(personalRecord)

	personalRecordNew, err := c.svc.UpdateRecord(personalRecord)

	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update personalRecord"})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "personalRecord update has been successfully",
		"data":    personalRecordNew,
	})

	return nil
}

func (c *PersonalRecordController) PatchRecord(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	if userId == "" {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	var jsonMap map[string]interface{}
	err := ctx.BodyParser(&jsonMap)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	fmt.Println(jsonMap)

	personalRecordNew, err := c.svc.PatchRecord(userId, jsonMap)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update personalRecord"})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "personalRecord update has been successfully",
		"data":    personalRecordNew,
	})

	return nil
}

/*
func (s *PersonalRecordController) DeleteRecord(ctx *fiber.Ctx) error {
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



func (s *PersonalRecordController) CreateRecord(ctx *fiber.Ctx) error {
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
*/

func (c *PersonalRecordController) SetupRoutes(app *fiber.App) {
	api := app.Group("/api/personal")
	api.Get("/records", c.GetRecordIndex)
	api.Get("/record/:id", c.GetRecordByID)
	api.Put("/record/:id", c.UpdateRecord)
	api.Patch("/record/:id", c.PatchRecord)
}
