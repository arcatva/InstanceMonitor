package handler

import (
	"backend/model"
	"backend/repository"
	"backend/util"
	"github.com/gofiber/fiber/v2"
	"log"
	"regexp"
	"strconv"
	"time"
)

// Register used for admin registration
func Register(ctx *fiber.Ctx) error { // this is used for admin registration
	data := make(map[string]any)
	// parse the context into data map
	if err := ctx.BodyParser(&data); err != nil {
		log.Println(err)
		ctx.Status(fiber.StatusServiceUnavailable)
		return ctx.JSON(fiber.Map{"message": "Can not parse the message"})
	}

	// input all params needed
	if data["email"] == nil ||
		data["first_name"] == nil ||
		data["last_name"] == nil ||
		data["password"] == nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{"message": "Please input all params"})
	}

	// check email is formatted
	if !validateEmail(data["email"].(string)) {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{"message": "Invalid email"})
	}

	// check password length
	if len(data["password"].(string)) <= 6 {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{"message": "Password must be more than 6 characters"})
	}

	// check if user exists in the repository
	userSearched := &model.User{}
	repository.PostgresDB.Where("email=?", data["email"].(string)).First(userSearched)
	if userSearched.Id != 0 {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{"message": "Email already exists"})
	}

	// create user
	user := &model.User{
		Email:     data["email"].(string),
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
	}
	if err := user.SetPassword(data["password"].(string)); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{"message": "Failed to set password"})
	}
	repository.PostgresDB.Create(user)
	ctx.Status(fiber.StatusOK)
	return ctx.JSON(user)

}

// use regex to match email format
func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

// Login used for admin logging
func Login(ctx *fiber.Ctx) error {
	data := make(map[string]any)

	// parse json data from request body
	if err := ctx.BodyParser(&data); err != nil {
		log.Println(err)
	}

	// check email
	if data["email"] == nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{"message": "Please input email address"})
	}

	// check if it is formatted email address
	if !validateEmail(data["email"].(string)) {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{"message": "Please input valid email address"})
	}

	userSearched := &model.User{}

	// find user with email matched
	repository.PostgresDB.Where("email=?", data["email"].(string)).Find(userSearched)
	if userSearched.Id == 0 {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{"message": "No existing user with this email address"})
	}

	// check if password inputted
	if data["password"] == nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{"message": "Please input password"})
	}

	// check if password in valid length
	if len(data["password"].(string)) <= 6 {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{"message": "Password must be more than 6 characters"})
	}

	// check if password matched
	password := data["password"].(string)
	if err := userSearched.ComparePassword(password); err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{"message": "Wrong password"})
	}

	// generate jwt token
	token, err := util.GenerateJwt(strconv.Itoa(int(userSearched.Id)))
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		log.Fatalln("Jwt token generation failed")
	}

	// return token, in production env should be encrypted with HTTPS
	ctx.Status(fiber.StatusOK)
	log.Println("Logged in as " + data["email"].(string) + " at " + time.Now().String())
	return ctx.JSON(fiber.Map{"token": token})

}
