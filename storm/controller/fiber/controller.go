package controller

// 각 기능 HTTP API 요청을 Fiber 프레임워크를 사용해서 서비스와 연결 설정

import (
	"log"
	"storm/repository"
	"storm/services"
	"storm/services/personal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type FiberHendler struct {
	App *fiber.App
}

func NewFiber() *FiberHendler {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
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
