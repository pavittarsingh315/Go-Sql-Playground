package routes

import (
	"errors"
	"fiber-gorm-tutorial/database"
	"fiber-gorm-tutorial/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

/*
   Serializer format:
   {
      id: 11,
      user: {
         id: 11,
         first_name: "pavittar"
         last_name: "singh"
      },
      product: {
         id: 22,
         name: "Macbook",
         serial_number: "128418468"
      }
   }
*/

type OrderSerializer struct {
	Id        uint              `json:"id"`
	User      UserSerializer    `json:"user"`
	Product   ProductSerializer `json:"product"`
	CreatedAt time.Time         `json:"order_date"`
}

func CreateResponseOrder(orderModel models.Order, user UserSerializer, product ProductSerializer) OrderSerializer {
	return OrderSerializer{Id: orderModel.Id, User: user, Product: product, CreatedAt: orderModel.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(fiber.StatusOK).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []OrderSerializer{}

	for _, order := range orders {
		var user models.User
		var product models.Product
		findUser(order.UserRefer, &user)
		findProduct(order.ProductRefer, &product)
		responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(fiber.StatusOK).JSON(responseOrders)
}

func findOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.Id == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var order models.Order

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := findOrder(id, &order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var user models.User
	var product models.Product
	findUser(order.UserRefer, &user)
	findProduct(order.ProductRefer, &product)
	responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))

	return c.Status(fiber.StatusOK).JSON(responseOrder)
}
