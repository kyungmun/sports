package repository

import (
	"fmt"
	"log"
	"storm/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewPersonalRecordRepository(db *gorm.DB) (*PersonalRecordRepository, error) {
	return &PersonalRecordRepository{db: db}, nil
}

type PersonalRecordRepository struct {
	db *gorm.DB
}

func (r *PersonalRecordRepository) GetIndex() (*[]models.PersonalRecord, error) {
	personalRecords := &[]models.PersonalRecord{}

	err := r.db.Find(personalRecords).Error
	if err != nil {
		return nil, err
	}

	return personalRecords, nil
}

func (r *PersonalRecordRepository) UpdateRecord(personalRecord *models.PersonalRecord) (*models.PersonalRecord, error) {

	err := r.db.Model(&personalRecord).Where("m_id = ?", personalRecord.UserID).Updates(personalRecord).Error
	if err != nil {
		return nil, err
	}

	return personalRecord, nil
}

func (r *PersonalRecordRepository) DeleteByID(id string) error {

	personalRecord := &models.PersonalRecord{}

	result := r.db.Where("m_id = ?", id).First(personalRecord)
	if result.RowsAffected == 0 { // returns count of records found
		fmt.Printf("count : 0")
		return fiber.NewError(fiber.StatusNotFound, "No Record Found")
	}

	err := r.db.Where("m_id = ?", id).Delete(personalRecord).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *PersonalRecordRepository) GetByID(id string) (*models.PersonalRecord, error) {
	//fmt.Println(">> ", m.DBEngine)

	personalRecord := &models.PersonalRecord{}
	//id := context.Params("id")

	err := r.db.Where("m_id = ?", id).Order("m_destday desc").First(personalRecord).Error
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return personalRecord, nil
}

func (r *PersonalRecordRepository) Create(personalRecord *models.PersonalRecord) (*models.PersonalRecord, error) {

	//personalRecords := &models.PersonalRecords{}

	//personalRecords.Assist = personalRecord.Assist

	err := r.db.Create(&personalRecord).Error

	if err != nil {
		return nil, err
	}

	return personalRecord, nil
}

/*
GetByID(id int64) (*PersonalRecord, error)
Fetch(offset, limit int) ([]*PersonalRecord, error)
Create(personalRecord *PersonalRecord) (*PersonalRecord, error)
Update(id int64, personalRecord *PersonalRecord) (*PersonalRecord, error)
Delete(id int64) error
*/
