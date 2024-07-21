package route

import (
	"go-esb-test/internal/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	InvoiceController *controller.InvoiceController
	UserController    *controller.UserController
	ItemController    *controller.ItemController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	// invoices routes
	invoices := c.App.Group("/invoices")
	{
		invoices.Get("", c.InvoiceController.List)
		invoices.Get("/:id", c.InvoiceController.Get)
		invoices.Post("", c.InvoiceController.Create)
		invoices.Put("/:id", c.InvoiceController.Update)
		invoices.Delete("/:id", c.InvoiceController.Delete)
	}

	// users routes
	c.App.Get("/users", c.UserController.List)

	// items routes
	c.App.Get("/items", c.ItemController.List)

}

func (c *RouteConfig) SetupAuthRoute() {

}
