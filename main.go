package main

import (
	"fmt"
	"gobe/api/routes"
	"gobe/api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "rahasia123", "bootcampgo")
	dbx, err := sqlx.Connect("postgres", connection)
	if err != nil {
		fmt.Print(err.Error())	
	}
	s := services.InitService(dbx)
	routes.BeRoutes(app, s)
	app.Listen(":26")

}
