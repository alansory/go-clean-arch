package config

import (
	"go-esb-test/internal/delivery/http/controller"
	"go-esb-test/internal/delivery/http/route"
	"go-esb-test/internal/repository"
	"go-esb-test/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	invoiceRepository := repository.NewInvoiceRepository(config.Log)
	itemRepository := repository.NewItemRepository(config.Log)
	userRepository := repository.NewUserRepository(config.Log)
	invoiceItemRepository := repository.NewInvoiceItemRepository(config.Log)

	// setup use cases
	invoiceUseCase := usecase.NewInvoiceUseCase(config.DB, config.Log, config.Validate, invoiceRepository, userRepository, invoiceItemRepository)
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	itemUseCase := usecase.NewItemUseCase(config.DB, config.Log, config.Validate, itemRepository)

	// setup controller
	invoiceController := controller.NewInvoiceController(invoiceUseCase, config.Log)
	userController := controller.NewUserController(userUseCase, config.Log)
	itemController := controller.NewItemController(itemUseCase, config.Log)

	// setup middleware

	routeConfig := route.RouteConfig{
		App:               config.App,
		InvoiceController: invoiceController,
		UserController:    userController,
		ItemController:    itemController,
	}

	routeConfig.Setup()
}
