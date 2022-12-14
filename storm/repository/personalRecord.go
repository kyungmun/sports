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

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 2
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (r *PersonalRecordRepository) GetIndex(page, pageSize int) (*[]models.PersonalRecord, error) {
	personalRecords := &[]models.PersonalRecord{}

	if (page <= 0) && (pageSize <= 0) {
		err := r.db.Find(personalRecords).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := r.db.Scopes(Paginate(page, pageSize)).Find(personalRecords).Error
		if err != nil {
			return nil, err
		}
	}

	return personalRecords, nil
}

func (r *PersonalRecordRepository) UpdateRecord(personalRecord *models.PersonalRecord) (*models.PersonalRecord, error) {

	fmt.Println(personalRecord)

	//grom 에서는 구조체로 저장시 0 값인 필드는 업데이트 하지 않는다., 맵으로 변환하여 저장해야함.
	//Put 메소드로 Update 에서는 전체 필드를 넘겨 받아서 처리한다. 단, 값이 숫자 필드가 0인 값이 있다면 맵으로 변환해서 처리 해야함.
	//err := r.db.Model(&personalRecord).Where("m_id = ?", personalRecord.UserID).Updates(personalRecord).Error

	//맵으로 변경해야할 값을 넘기면 0 값도 저장 됨.
	//err := r.db.Model(&personalRecord).Where("m_id = ?", personalRecord.UserID).Updates(jsonData).Error

	//select 해서 변경할 필드를 지정 가능. 이렇게하면 0 도 저장 됨.
	err := r.db.Model(&personalRecord).Select("*").Where("m_id = ?", personalRecord.UserID).Updates(personalRecord).Error
	if err != nil {
		return nil, err
	}

	return personalRecord, nil
}

func (r *PersonalRecordRepository) PatchRecord(userId string, jsonData map[string]interface{}) (*models.PersonalRecord, error) {

	fmt.Println(jsonData)

	personalRecord := &models.PersonalRecord{}

	//patch update 일때는 맵으로 전달된 필드만 업데이트 한다.
	err := r.db.Model(&personalRecord).Where("m_id = ?", userId).Updates(jsonData).Error
	if err != nil {
		return nil, err
	}

	err2 := r.db.Where("m_id = ?", userId).Order("m_destday desc").First(personalRecord).Error
	if err2 != nil {
		log.Printf("%v", err2)
		return nil, err2
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
