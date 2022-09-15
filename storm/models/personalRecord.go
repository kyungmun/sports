package models

import "gorm.io/gorm"

type PersonalRecords struct {
	gorm.Model
	UserID  string `gorm:"column:m_id; primary key; type:varchar(20); not null" json:"m_id" validate:"required"`  // varchar(20) NOT NULL default '',
	Check   string `gorm:"column:m_check; key; type:varchar(1); not null" json:"m_check" validate:"required"`     // varchar(1) NOT NULL default 'Y',
	DestDay string `gorm:"column:m_destday; key; type:varchar(8); not null" json:"m_destday" validate:"required"` // varchar(8) NOT NULL,
	Name    string `gorm:"column:m_name; type:varchar(20); not null" json:"m_name" validate:"required"`           // varchar(20) NOT NULL default '',
	Result  string `gorm:"column:m_result; type:varchar(1); not null" json:"m_result" validate:"required"`        // varchar(1) NOT NULL default '-',
	Mvp     int    `gorm:"column:m_mvp;" json:"m_mvp" validate:"required"`                                        // int(1) NOT NULL default '0',
	Point   int    `gorm:"column:m_point;" json:"m_point" validate:"required"`                                    // int(3) NOT NULL default '0',
	Gole    int    `gorm:"column:m_gole;" json:"m_gole" validate:"required"`                                      // int(2) NOT NULL default '0',
	Assist  int    `gorm:"column:m_assist;" json:"m_assist" validate:"required"`                                  // int(2) NOT NULL default '0',
	Mobile  int    `gorm:"column:m_mobile;" json:"m_mobile" validate:"required"`                                  // int(1) NOT NULL default '0',
}

type PersonalRecord struct {
	UserID  string `gorm:"column:m_id;" json:"m_id"            validate:"required"`
	Check   string `gorm:"column:m_check;" json:"m_check"      validate:"required"`
	DestDay string `gorm:"column:m_destday;" json:"m_destday"  validate:"required"`
	Name    string `gorm:"column:m_name;" json:"m_name"        validate:"required"`
	Result  string `gorm:"column:m_result;" json:"m_result"    validate:"required"`
	Mvp     int    `gorm:"column:m_mvp;" json:"m_mvp"          validate:"required"`
	Point   int    `gorm:"column:m_point;" json:"m_point"      validate:"required"`
	Gole    int    `gorm:"column:m_gole;" json:"m_gole"        validate:"required"`
	Assist  int    `gorm:"column:m_assist;" json:"m_assist"    validate:"required"`
	Mobile  int    `gorm:"column:m_mobile;" json:"m_mobile"    validate:"required"`
}

// TableName overrides the table name used by User to `tablename`
// 자동으로 생성되는 테이블명은 구조체명을 기준으로 스네이크 스타일 및 끝에 s 문자가 자동으로 붙기에 지정한 테이블명을 사용하려면 아래 메소드 구현 필요.
//func (PersonalRecords) TableName() string {
//	return "personal_record_test"
//}

func MigratePersonalRecords(db *gorm.DB) error {
	err := db.AutoMigrate(&PersonalRecords{})
	return err
}

// 데이터 저장용 인터페이스
type PersonalRecordRepository interface {
	GetIndex() ([]*PersonalRecord, error)
	GetByID(id string) (*PersonalRecord, error)
	Fetch(offset, limit int) ([]*PersonalRecord, error)
	Create(personalRecord *PersonalRecord) (*PersonalRecord, error)
	Update(id int64, personalRecord *PersonalRecord) (*PersonalRecord, error)
	Delete(id int64) error
}

// 데이터 비즈니스로직 처리용 인터페이스
type PersonalRecordUseCase interface {
	GetByID(id string) (*PersonalRecord, error)
	Fetch(offset, limit int) ([]*PersonalRecord, error)
	Create(personalRecord *PersonalRecord) (*PersonalRecord, error)
	Update(id int64, personalRecord *PersonalRecord) (*PersonalRecord, error)
	Delete(id int64) error
}
