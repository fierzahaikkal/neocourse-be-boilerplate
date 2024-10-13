package handler

import (
	"log"

	"github.com/fierzahaikkal/neocourse-be-boilerplate-golang/internal/entity"
	"github.com/fierzahaikkal/neocourse-be-boilerplate-golang/internal/model/user"
	"github.com/fierzahaikkal/neocourse-be-boilerplate-golang/internal/usecase"
	"github.com/fierzahaikkal/neocourse-be-boilerplate-golang/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	useCase   *usecase.AuthUseCase
	log       *log.Logger
	JWTSecret string
}

func NewAuthHandler(useCase *usecase.AuthUseCase, JWTSecret string) *AuthHandler {
	return &AuthHandler{
		useCase:   useCase,
		JWTSecret: JWTSecret,
	}
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	var req user.SignUpRequest
	var user entity.User

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := h.useCase.Register(&req)
	if err != nil {
		if err == utils.ErrUserExists {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
		}
	}

	token, err := utils.GenerateJWT(&user, h.JWTSecret)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return utils.SuccessResponse(ctx, token, fiber.StatusOK)
}

func (h *AuthHandler) SignIn(ctx *fiber.Ctx) error {
	var req user.SignInRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	userFromDB, err := h.useCase.SignIn(&req)
	if err != nil {
		if err == utils.ErrRecordNotFound {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "An error occurred"})
	}

	token, err := utils.GenerateJWT(userFromDB, h.JWTSecret)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
