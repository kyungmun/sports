package controller

import (
	"storm/services"
	"context"
)

type PersonalControllerIface interface {
	Listen(port string)
	SetupRoutes(svc *services.ServiceRepository)
}

type BookService interface {
	PrintBookTitle(ctx context.Context)
	TestBookService(ctx context.Context)
}