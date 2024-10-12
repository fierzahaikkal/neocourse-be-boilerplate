package handler

import (
	"log"

	"github.com/fierzahaikkal/neocourse-be-boilerplate-golang/internal/usecase"
	"github.com/fierzahaikkal/neocourse-be-boilerplate-golang/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	useCase   *usecase.AuthUseCase
	log       *log.Logger
	JWTSecret string
}

func NewAuthHandler(useCase *usecase.AuthUseCase, log *log.Logger, JWTSecret string) *AuthHandler {
	return &AuthHandler{
		useCase:   useCase,
		log:       log,
		JWTSecret: JWTSecret,
	}
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	var req 
}
