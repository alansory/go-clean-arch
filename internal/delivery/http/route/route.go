package route

import (
	"go-esb-test/internal/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	InvoiceController *controller.InvoiceController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/invoices", c.InvoiceController.Create)
	c.App.Get("/invoices/:id", c.InvoiceController.Get)
	c.App.Delete("/invoices/:id", c.InvoiceController.Delete)
}

func (c *RouteConfig) SetupAuthRoute() {

}
