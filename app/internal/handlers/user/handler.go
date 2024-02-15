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
	authorized := r.Group("/api/user/")
	authorized.Use(jwt.Middleware(h.cfg.GetSecret()))
	authorized.GET("/:id", h.profile)
	authorized.GET("/list", h.list)
	authorized.POST("/create", h.create)
}
