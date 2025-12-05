package auth

import (
	"github.com/AlexGo12311/TwitterClone/internal/common/cache"
	"github.com/AlexGo12311/TwitterClone/internal/common/database"
	"github.com/AlexGo12311/TwitterClone/internal/common/middleware"
	"github.com/AlexGo12311/TwitterClone/internal/module/auth/action"
	"github.com/AlexGo12311/TwitterClone/internal/module/auth/service"
	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router, db database.Database, cache cache.Cache) {
	authMiddleware := middleware.NewAuthMiddleware()
	r.Post("/login", buildLoginHandler(db))
	r.Get("/me", authMiddleware.Execute(), buildMeHandler(db))
	r.Get("/token", buildTokenHandler(db, cache))
	r.Post("/logout", buildLogoutHandler(cache))
}

func buildLoginHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewLoginService(db)
		action := action.NewLoginAction(service)

		return action.Execute(c)
	}
}

func buildMeHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewMeService(db)
		action := action.NewMeAction(service)

		return action.Execute(c)
	}
}

func buildTokenHandler(db database.Database, cache cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewTokenService(db, cache)
		action := action.NewTokenAction(service)

		return action.Execute(c)
	}
}

func buildLogoutHandler(cache cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewLogoutService(cache)
		action := action.NewLogoutAction(service)

		return action.Execute(c)
	}
}
