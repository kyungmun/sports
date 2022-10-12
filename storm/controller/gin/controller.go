package controller

// 각 기능 HTTP API 요청을 Gin 프레임워크를 사용해서 서비스와 연결 설정

import (
	"log"
	"storm/repository"
	"storm/services"
	"storm/services/personal"

	"github.com/gin-gonic/gin"
)

type GinHendler struct {
	App *gin.Engine
}

func NewGin() *GinHendler {
	app := gin.New()
	return &GinHendler{App: app}
}

func (f *GinHendler) Listen(port string) {
	err := f.App.Run(port)
	if err != nil {
		log.Panic("service start fail")
	}
}

func (f *GinHendler) SetupRoutes(svc *services.ServiceRepository) {
	//개인기록

	personalRepo, err := repository.NewPersonalRecordRepository(svc.DB)
	if err != nil {
		log.Fatal("could not personal repository create")
	}
	personalService, err := personal.NewPersonalRecordServices(personalRepo)
	if err != nil {
		log.Fatal("could not personal services create")
	}

	personalController := NewPersonalRecordController(personalService)
	personalController.SetupRoutes(f.App)

	//팀기록

}
