package mainpage

import (
	"catchbook/internal/config"
	"catchbook/internal/handlers"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type Handler struct {
	cfg    *config.Config
	logger *slog.Logger
}

func NewHandler(cfg *config.Config, l *slog.Logger) handlers.HandlerInterface {
	return &Handler{
		cfg:    cfg,
		logger: l,
	}
}

func (h *Handler) Register(r *gin.Engine) {
	r.GET("/", h.Main)
}

func (h *Handler) Main(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "welcome to the mainpage page"})
}
