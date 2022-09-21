package ui

import (
	"gobe/api/model"
	"gobe/api/services"

	"github.com/gofiber/fiber/v2"
)

func Home() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Home")
	}
}

// func About() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		return c.SendString("About")
// 	}
// }

func ListKategori(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req model.Halaman
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		data := s.PanggilkategoriX(req)
		return c.JSON(data)
	}

}

func ListKategoriTables(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req model.Halaman
		var output model.Pagedata
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		output.Page.Totalelements = s.BanyakKategoriX(req)
		output.Page.Pagenumber = req.Pagenumber
		output.Page.Size = req.Size
		output.Data = s.PanggilkategoriX(req)
		return c.JSON(output)
	}		
}


func ListProduk(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := s.PanggilProdukdbx()
		return c.JSON(data)
	}
}


func ListDept(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := s.PanggilDeptdbx()
		return c.JSON(data)
	}
}

func PostDept(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req model.DepartmentRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		data := s.InsertDeptdbx(req)
		return c.JSON(data)
	}

}





