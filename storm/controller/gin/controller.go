package controller

// 각 기능 HTTP API 요청을 Gin 프레임워크를 사용해서 서비스와 연결 설정

import (
	"log"
	"storm/services"

	"github.com/gofiber/fiber/v2"
)

type FiberHendler struct {
	App *fiber.App
}

func NewFiber() *FiberHendler {
	app := fiber.New()
	return &FiberHendler{App: app}
}

func (f *FiberHendler) Listen(port string) {
	err := f.App.Listen(port)
	if err != nil {
		log.Panic("service start fail")
	}
}

func (f *FiberHendler) SetupRoutes(svc *services.ServiceRepository) {
	//개인기록
	/*
		personalRepo, err := repository.NewPersonalRecordRepository(svc.DB)
		if err != nil {
			log.Fatal("could not personal repository create")
		}
		personalService, err := personal.NewPersonalRecordServices(personalRepo)
		if err != nil {
			log.Fatal("could not personal services create")
		}

		personalCtrl := NewPersonalRecordController(personalService)
		personalCtrl.SetupRoutes(f.App)
	*/
	//팀기록

}
