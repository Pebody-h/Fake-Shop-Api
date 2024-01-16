package handlers

import (
	"fmt"

	"github.com/Pebody-h/Fake-Shop-Api/database"
	"github.com/Pebody-h/Fake-Shop-Api/models"
	"github.com/gofiber/fiber/v2"
)

func ConnectionTest(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func GetProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []models.Product

	db.Find(&products)
	return c.JSON(products)
}

func GetProductById(c *fiber.Ctx) error {
	db := database.DB
	productId := c.Params("id")
	var product models.Product

	res := db.Where("ID = ?", productId).Limit(1).Find(&product)
	if res.Error != nil {
		return c.Status(400).JSON(res.Error)
	}

	if res.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	return c.JSON(product)
}

func PostProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []models.Product

	if err := c.BodyParser(products); err != nil {
		fmt.Println(products)
		return c.Status(400).JSON(err)
	}

	for _, p := range products {
		db.Create(&p)
	}
	return c.JSON(products)
}

func PostProduct(c *fiber.Ctx) error {
	db := database.DB
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(err)
	}

	db.Create(&product)
	return c.JSON(product)
}
