package controllers

import (
	"firstproject/models"
	"firstproject/services"
	"firstproject/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service}
}

func (pc *ProductController) AddProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse JSON",
			Error:   err.Error(),
		})
	}
	if err := pc.service.AddProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error adding product",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(utils.SuccessResponse{
		Status:  fiber.StatusCreated,
		Message: "Product added successfully",
		Data:    product,
	})
}

func (pc *ProductController) GetProducts(c *fiber.Ctx) error {
	products, err := pc.service.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error retrieving products",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Products retrieved successfully",
		Data:    products,
	})
}

func (pc *ProductController) UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse JSON",
			Error:   err.Error(),
		})
	}
	if err := pc.service.UpdateProduct(&product); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Product not found",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Product updated successfully",
		Data:    product,
	})
}

func (pc *ProductController) DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	if err := pc.service.DeleteProduct(productID); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: "Product not found",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Product deleted successfully",
		Data:    nil,
	})
}
