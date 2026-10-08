package main

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

// สร้าง structure ของ product เพื่อจำลองข้อมูลที่จะเก็บใน database
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// สร้างตัวแปรแบบ struct Product เพื่อจำลองข้อมูลที่จะเก็บใน database
var products = []Product{}

func main() {
	//สร้าง instance ของ fiber
	app := fiber.New()

	//ใส่ mock data ลงในตัวแปร products เพื่อมาใช้ทดสอบ API (crud)
	products = append(products, Product{ID: 1, Name: "iPhon 15 Pro Max", Price: 35000})
	products = append(products, Product{ID: 2, Name: "Oppo Reno 11 5G", Price: 12000})
	products = append(products, Product{ID: 3, Name: "Samsung Galaxy S23 Ultra", Price: 55000})

	//ใช้ middleware สำหรับ Cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},                                                 //อนุญาตทุก origin
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},           //อนุญาติทุก method
		AllowHeaders: []string{"Content-Type", "Authorization", "origin", "accept"}, //อนุญาติทุก header
	}))

	//สร้าง API สำหรับ CURD ข้อมูล คือ กำหนด route และค่าต่างๆ ที่ต้อง response กลับไปให้ client-user
	app.Get("/api/products", getAllProduct)

	// (R-Read) ดึงข้อมูลแบบมีเงื่อนไข
	app.Get("/api/products/:id", getProductById)

	// (C-Create) เพิ่มข้อมูลใหม่
	app.Post("/api/products", createProduct)

	// (U-Update) อัพเดทข้อมูล
	app.Put("/api/products/:id", updateProduct)

	// (D-Delete) ลบข้อมูล
	app.Delete("/api/products/:id", deleteProduct)

	if err := app.Listen(":5599"); err != nil {
		panic(err)
	}

}

// -------------------------------------------------------------------------------------------
// ฟังก์ชัน handle สำหรับ CRUD ข้อมูล
func getAllProduct(c fiber.Ctx) error {
	return c.JSON(products)
}

// (R-Read) ดึงข้อมูลแบบมีเงื่อนไข
func getProductById(c fiber.Ctx) error {
	id := c.Params("id")

	//validate ข้อมูลที่ส่งมา ว่าเป็นตัวเลขหรือไม่
	if _, err := strconv.Atoi(id); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	//วนลูปหาข้อมูลตาม ID ที่ส่งมา
	for _, productResponse := range products {
		if strconv.Itoa(productResponse.ID) == id {
			return c.Status(200).JSON(productResponse)
		}
	}

	//ถ้าหาข้อมูลไม่พบ
	return c.Status(404).JSON(fiber.Map{
		"message": "Product not found",
	})
}

// ฟังก์ชัน handle สำหรับสร้างข้อมูลใหม่
func createProduct(c fiber.Ctx) error {
	//สร้างตัวแปรเก็บข้อมูลที่ส่งมาจาก client-user
	product := new(Product)

	//อ่านข้อมูลจาก body ของ request และเก็บลงในตัวแปร product
	if err := c.Bind().Body(product); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to parse product",
		})
	}

	//เพิ่มข้อมูลเข้าไปในตัวแปร products ที่เราเริ่มกำหนดไว้ด้านบน
	products = append(products, *product)
	//เพิ่มเสร็จแล้ว return กลับไปบอก client-user
	return c.Status(201).JSON(fiber.Map{
		"message": "Product created successfully",
	})

}

// ฟังก์ชัน handle สำหรับอัพเดทข้อมูล
func updateProduct(c fiber.Ctx) error {
	id := c.Params("id")

	//validate ข้อมูลที่ส่งมา ว่าเป็นตัวเลขหรือไม่
	if _, err := strconv.Atoi(id); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	//สร้างตัวแปรเก็บข้อมูลที่ส่งมาจาก client-user
	product := new(Product)

	//อ่านข้อมูลจาก body ของ request และเก็บลงในตัวแปร product
	if err := c.Bind().Body(product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid product data",
		})
	}

	//วนลูปหาข้อมูลตาม ID ที่ส่งมา
	for i, productData := range products {
		if strconv.Itoa(productData.ID) == id {
			products[i].Name = product.Name
			products[i].Price = product.Price

			return c.Status(200).JSON(fiber.Map{
				"message": "Product updated successfully",
			})
		}
	}

	//ถ้าหาข้อมูลไม่พบ
	return c.Status(404).JSON(fiber.Map{
		"message": "Product not found",
	})

}

// ฟังก์ชัน handle สำหรับลบข้อมูล
func deleteProduct(c fiber.Ctx) error {
	id := c.Params("id")

	//validate ข้อมูลที่ส่งมา ว่าเป็นตัวเลขหรือไม่
	if _, err := strconv.Atoi(id); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	//วนลูปหาข้อมูลตาม ID ที่ส่งมา
	for i, productData := range products {
		if strconv.Itoa(productData.ID) == id {
			products = append(products[:i], products[i+1:]...)
			return c.Status(200).JSON(fiber.Map{
				"message": "Product deleted successfully",
			})
		}
	}

	//ถ้าหาข้อมูลไม่พบ
	return c.Status(404).JSON(fiber.Map{
		"message": "Product not found",
	})

}
