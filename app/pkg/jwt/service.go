package jwt

import (
	"catchbook/pkg/cache"
	"encoding/json"
	"github.com/cristalhq/jwt/v3"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type UserInterface interface {
	GetUserId() uint
	GetEmail() string
}

type RToken struct {
	Token string `json:"refresh_token"`
}

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserClaims struct {
	ID        uint             `json:"id,omitempty"`
	Email     string           `json:"email,omitempty"`
	ExpiresAt *jwt.NumericDate `json:"expires_at,omitempty"`
}

type ServiceInterface interface {
	GenerateAccessToken(u UserInterface, secret []byte) (*TokenResponse, error)
	RefreshToken(rt RToken, secret []byte) (*TokenResponse, error)
}

type Service struct {
	logger *slog.Logger
	cache  cache.CacheInterface
}

func NewService(l *slog.Logger, c cache.CacheInterface) ServiceInterface {
	return &Service{logger: l, cache: c}
}

func (s *Service) GenerateAccessToken(u UserInterface, secret []byte) (*TokenResponse, error) {
	signer, err := jwt.NewSignerHS(jwt.HS256, secret)
	if err != nil {
		return nil, err
	}
	builder := jwt.NewBuilder(signer)
	token, err := builder.Build(s.claims(u))
	if err != nil {
		return nil, err
	}
	s.logger.Info("token successfully generated")
	refreshTokenUuid := uuid.New()
	userBytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	err = s.cache.Set([]byte(refreshTokenUuid.String()), userBytes, 0)
	result := &TokenResponse{
		Token:        token.String(),
		RefreshToken: refreshTokenUuid.String(),
	}
	return result, nil
}

func (s *Service) claims(u UserInterface) UserClaims {
	return UserClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
		ID:        u.GetUserId(),
		Email:     u.GetEmail(),
	}
}

func (s *Service) RefreshToken(rt RToken, secret []byte) (*TokenResponse, error) {
	defer s.cache.Del([]byte(rt.Token))
	userBytes, err := s.cache.Get([]byte(rt.Token))
	if err != nil {
		return nil, err
	}
	var u UserInterface
	err = json.Unmarshal(userBytes, u)
	if err != nil {
		return nil, err
	}
	return s.GenerateAccessToken(u, secret)
}
