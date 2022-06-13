package controllers

import (
	"backend/configs"
	"backend/models"
	"backend/responses"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")
var validate = validator.New()

func Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func Login(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"email": data["email"]}).Decode(&user)

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{

		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(os.Getenv("SecretKey")))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	// cookie := fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    token,
	// 	Expires:  time.Now().Add(time.Hour * 24),
	// 	HTTPOnly: true,
	// }

	// c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
		"token":   token,
	})

	// return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

// getting user with token
func User(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var token1 = c.Params("token")
	defer cancel()
	// cookie := c.Cookies("jwt")

	if token1 == "" {
		return c.JSON(fiber.Map{
			"message": "provide login details",
		})
	}

	token, err := jwt.ParseWithClaims(token1, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SecretKey")), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	doc := userCollection.FindOne(ctx, bson.M{"email": claims.Issuer}).Decode(&user)
	fmt.Println(doc)

	return c.JSON(fiber.Map{
		"data":    user,
		"message": "success",
		"token":   claims.Issuer,
	})
}

func Logout(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "logged out successfully",
		"token":   "",
		"user":    "",
	})
}

func CreateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	newProduct := models.Product{
		Pro_Id:      primitive.NewObjectID(),
		ProductName: product.ProductName,
		Cost:        product.Cost,
		UserMail:    product.UserMail,
	}

	result, err := productCollection.InsertOne(ctx, newProduct)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.ProductResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetProduct(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	val := c.Params("usermail")
	var products []models.Product
	defer cancel()

	// objId, _ := primitive.ObjectIDFromHex(userId)

	results, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "No Products",
		})
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleproduct models.Product
		if err = results.Decode(&singleproduct); err != nil {
			return c.JSON(fiber.Map{
				"message": "failed",
			})
		}
		if singleproduct.UserMail == val {
			products = append(products, singleproduct)
		}
	}
	return c.JSON(fiber.Map{
		"message":  "success",
		"products": products,
	})

}

func EditProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	productId := c.Params("ProductId")
	var product models.Product
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(productId)

	//validate the request body
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	update := bson.M{"productname": product.ProductName, "cost": product.Cost}

	result, err := productCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.JSON(fiber.Map{
		"message":         "successfully updated",
		"updated product": result,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	productId := c.Params("productId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(productId)

	result, err := productCollection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
	)
}
