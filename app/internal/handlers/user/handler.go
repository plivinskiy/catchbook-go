package user

import (
	"catchbook/internal/config"
	"catchbook/internal/handlers"
	uc "catchbook/internal/usecase/user"
	"catchbook/pkg/jwt"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	fetchUserUseCase  uc.FetchUserUseCaseInterface
	createUserUseCase uc.CreateUserUseCaseInterface
	listUserUseCase   uc.ListUserUseCaseInterface
	cfg               *config.Config
	logger            *slog.Logger
}

func NewHandler(
	fetchUserUseCase uc.FetchUserUseCaseInterface,
	createUserUseCase uc.CreateUserUseCaseInterface,
	listUserUseCase uc.ListUserUseCaseInterface,
	cfg *config.Config,
	l *slog.Logger,
) handlers.HandlerInterface {
	return &Handler{
		fetchUserUseCase:  fetchUserUseCase,
		createUserUseCase: createUserUseCase,
		listUserUseCase:   listUserUseCase,
		cfg:               cfg,
		logger:            l,
	}
}

func (h *Handler) Register(r *gin.Engine) {
	r.Use(jwt.Middleware())
	r.GET("/api/user/:id", h.profile)
	r.GET("/api/users", h.list)
	r.POST("/api/user/create", h.create)
}
