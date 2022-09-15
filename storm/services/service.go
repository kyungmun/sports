package services

import (
	"gorm.io/gorm"
)

type ServiceRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{DB: db}
}

/*
func (s *ServiceRepository) SetupRoutes(app *fiber.App) {

	personalRepo, err := repository.NewPersonalRecordRepository(s.DB)
	if err != nil {
		log.Fatal("could not personal repository create")
	}
	personalService, err := personal.NewPersonalRecordServices(personalRepo)
	if err != nil {
		log.Fatal("could not personal services create")
	}
	personalService.SetupRoutes(app)

	// team := &team.Repository{
	// 	DB: r.DB,
	// }

	// team.SetupRoutes(app)
}
*/
