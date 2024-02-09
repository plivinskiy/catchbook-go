package auth

import (
	"catchbook/internal/config"
	"catchbook/internal/handlers"
	"catchbook/internal/usecase/auth"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	authUseCase auth.AuthorizeUseCaseInterface
	cfg         *config.Config
	logger      *slog.Logger
}

func NewHandler(authUseCase auth.AuthorizeUseCaseInterface, cfg *config.Config, l *slog.Logger) handlers.HandlerInterface {
	return &Handler{
		authUseCase: authUseCase,
		cfg:         cfg,
		logger:      l,
	}
}

func (h *Handler) Register(r *gin.Engine) {
	r.POST("/auth/token", h.token)
}
