package routes

import (
	"gobe/api/services"
	"gobe/api/ui"

	"github.com/gofiber/fiber/v2"
)

func BeRoutes(app fiber.Router, s services.Service) {
	app.Get("/",ui.Home())
	// app.Get("/about",ui.About())
	app.Get("/api/kategorilist",ui.ListKategori(s))
	app.Get("/api/produklist",ui.ListProduk(s))
	app.Post("/api/kategoridt" ,ui.ListKategoriTables(s))
	app.Get("/api/deptlist",ui.ListDept(s))
	app.Post("/api/deptpost",ui.PostDept(s))
}