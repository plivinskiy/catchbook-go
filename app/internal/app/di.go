package app

import (
	"catchbook/internal/config"
	"catchbook/internal/handlers"
	ah "catchbook/internal/handlers/auth"
	fh "catchbook/internal/handlers/fish"
	mh "catchbook/internal/handlers/mainpage"
	uh "catchbook/internal/handlers/user"
	fr "catchbook/internal/repository/fish"
	ur "catchbook/internal/repository/user"
	fs "catchbook/internal/service/fish"
	us "catchbook/internal/service/user"
	aucase "catchbook/internal/usecase/auth"
	uucase "catchbook/internal/usecase/user"
	"catchbook/pkg/cache"
	"catchbook/pkg/cache/freecache"
	"catchbook/pkg/db"
	"catchbook/pkg/jwt"
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
	"os"
)

type Container struct {
	service struct {
		user us.ServiceInterface
		fish fs.ServiceInterface
		jwt  jwt.ServiceInterface
	}

	useCase struct {
		user struct {
			create uucase.CreateUserUseCaseInterface
			list   uucase.ListUserUseCaseInterface
			fetch  uucase.FetchUserUseCaseInterface
		}
		auth struct {
			authorize aucase.AuthorizeUseCaseInterface
		}
	}

	repository struct {
		user us.RepositoryInterface
		fish fs.RepositoryInterface
	}

	handler struct {
		user handlers.HandlerInterface
		main handlers.HandlerInterface
		fish handlers.HandlerInterface
		auth handlers.HandlerInterface
	}

	db     *sql.DB
	config *config.Config
	logger *slog.Logger
	ctx    context.Context
	cache  cache.CacheInterface
}

func NewContainer(ctx context.Context) *Container {
	return &Container{
		ctx: ctx,
	}
}

func (c *Container) GetConfig() *config.Config {
	if c.config == nil {
		c.config = config.CreateConfig()
	}
	return c.config
}

func (c *Container) GetLogger() *slog.Logger {
	if c.logger == nil {
		switch c.GetConfig().ENV {
		case config.EnvDev:
			c.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case config.EnvProd:
			c.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		case config.EnvLocal:
			c.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		}
	}
	return c.logger
}

// --------
// Services

func (c *Container) GetUserService() us.ServiceInterface {
	if c.service.user == nil {
		c.service.user = us.NewService(c.GetUserRepository())
	}
	return c.service.user
}

func (c *Container) GetFishService() fs.ServiceInterface {
	if c.service.fish == nil {
		c.service.fish = fs.NewService(c.GetFishRepository())
	}
	return c.service.fish
}

func (c *Container) GetJwtService() jwt.ServiceInterface {
	if c.service.jwt == nil {
		c.service.jwt = jwt.NewService(c.logger, c.GetCache())
	}
	return c.service.jwt
}

// ------------
// Repositories

func (c *Container) GetUserRepository() us.RepositoryInterface {
	if c.repository.user == nil {
		c.repository.user = ur.NewUserRepository(c.GetDatabaseConnection())
	}
	return c.repository.user
}

func (c *Container) GetFishRepository() fs.RepositoryInterface {
	if c.repository.fish == nil {
		c.repository.fish = fr.NewFishRepository(c.GetDatabaseConnection())
	}
	return c.repository.fish
}

// ------------
// Use Cases

func (c *Container) GetUserCreateUseCase() uucase.CreateUserUseCaseInterface {
	if c.useCase.user.create == nil {
		c.useCase.user.create = uucase.NewUseCaseCreateUser(c.GetUserService())
	}
	return c.useCase.user.create
}

func (c *Container) GetUserFetchUseCase() uucase.FetchUserUseCaseInterface {
	if c.useCase.user.fetch == nil {
		c.useCase.user.fetch = uucase.NewUseCaseFetchUser(c.GetUserService())
	}
	return c.useCase.user.fetch
}
func (c *Container) GetUserListUseCase() uucase.ListUserUseCaseInterface {
	if c.useCase.user.list == nil {
		c.useCase.user.list = uucase.NewUseCaseListUser(c.GetUserService())
	}
	return c.useCase.user.list
}

func (c *Container) GetAuthUseCase() aucase.AuthorizeUseCaseInterface {
	if c.useCase.auth.authorize == nil {
		c.useCase.auth.authorize = aucase.NewAuthorizeUseCase(c.GetUserService(), c.GetJwtService())
	}
	return c.useCase.auth.authorize
}

// ------------
// Handlers

func (c *Container) getUserHandler() handlers.HandlerInterface {
	if c.handler.user == nil {
		c.handler.user = uh.NewHandler(
			c.GetUserFetchUseCase(),
			c.GetUserCreateUseCase(),
			c.GetUserListUseCase(),
			c.GetConfig(),
			c.GetLogger(),
		)
	}
	return c.handler.user
}

func (c *Container) getFishHandler() handlers.HandlerInterface {
	if c.handler.fish == nil {
		c.handler.fish = fh.NewHandler(c.GetConfig(), c.GetLogger())
	}
	return c.handler.fish
}

func (c *Container) getAuthHandler() handlers.HandlerInterface {
	if c.handler.auth == nil {
		c.handler.auth = ah.NewHandler(c.GetAuthUseCase(), c.GetConfig(), c.GetLogger())
	}
	return c.handler.auth
}

func (c *Container) getMainHandler() handlers.HandlerInterface {
	if c.handler.main == nil {
		c.handler.main = mh.NewHandler(c.GetConfig(), c.GetLogger())
	}
	return c.handler.main
}

func (c *Container) Handlers() []handlers.HandlerInterface {
	return []handlers.HandlerInterface{
		c.getUserHandler(),
		c.getMainHandler(),
	}
}

// ------------
// Adapters

func (c *Container) GetDatabaseConnection() *sql.DB {
	if c.db == nil {
		c.db = db.NewMysqlClient(c.ctx, c.GetConfig().DatabaseDsn)
	}
	return c.db
}

func (c *Container) GetCache() cache.CacheInterface {
	if c.cache == nil {
		c.cache = freecache.NewCache(1024)
	}
	return c.cache
}
