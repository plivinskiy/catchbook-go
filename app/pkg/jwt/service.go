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
	GetUserId() string
	GetEmail() string
}

type RToken struct {
	Token string `json:"refresh_token"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	Email string
}

type ServiceInterface interface {
	GenerateAccessToken(u UserInterface, secret string) ([]byte, error)
	RefreshToken(rt RToken, secret string) ([]byte, error)
}

type Service struct {
	logger *slog.Logger
	cache  cache.CacheInterface
}

func NewService(l *slog.Logger, c cache.CacheInterface) ServiceInterface {
	return &Service{logger: l, cache: c}
}

func (s *Service) GenerateAccessToken(u UserInterface, secret string) ([]byte, error) {
	signer, err := jwt.NewSignerHS(jwt.HS256, []byte(secret))
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
	result, err := json.Marshal(map[string]string{
		"token":         token.String(),
		"refresh_token": refreshTokenUuid.String(),
	})
	return result, nil
}

func (s *Service) claims(u UserInterface) UserClaims {
	return UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        u.GetUserId(),
			Audience:  []string{"user"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
		},
		Email: u.GetEmail(),
	}
}

func (s *Service) RefreshToken(rt RToken, secret string) ([]byte, error) {
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
