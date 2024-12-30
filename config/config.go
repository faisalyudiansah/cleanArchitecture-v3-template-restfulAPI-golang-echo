package config

import (
	httproute "server/internal/adapter/http"
	"server/internal/adapter/http/controller"
	"server/internal/adapter/repository"
	"server/internal/adapter/validator"
	"server/internal/infrastructure/database"
	"server/internal/infrastructure/logger"
	"server/internal/usecase"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB        *database.Kind[*gorm.DB]
	App       *echo.Echo
	Logger    *logger.Logger
	Validator *validator.CustomValidator
	Config    *viper.Viper
	// Publisher        *messaging.Publisher[any, any]
}

func (c *BootstrapConfig) SetDefaultConfigs() {
	setDefaultCoreConfigs(c.Config)
	setDefaultDatabaseConfig(c.Config)
	setMySQLDefaultConfig(c.Config)
}

func (c *BootstrapConfig) WatchConfig() {
	c.Config.OnConfigChange(func(fsnotify.Event) {
		c.SetDefaultConfigs()
	})
	c.Config.WatchConfig()
}

// Bootstrap bootstrap app
func Bootstrap(config *BootstrapConfig) error {
	// config.SetDefaultConfigs()
	// config.Config.WatchConfig()

	// Init repositories
	repo := repository.NewRepository(config.DB)
	exampleRepo := repository.NewExampleRepository(config.Logger, config.DB)

	// init domain validator
	// SOItemValidator := domain_validator.NewSOItemValidator()

	// Init usecases
	healtcheckUseCase := usecase.NewHealthCheckUseCase(
		config.Logger,
		exampleRepo,
	)
	exampleUsecase := usecase.NewExampleUsecase(config.Logger, config.Validator, repo, exampleRepo)

	// Init controllers
	healtcheckController := controller.NewHealthCheckController(config.Logger, healtcheckUseCase)
	logController := controller.NewLogController(config.Logger, config.Validator)
	exampleController := controller.NewExampleController(config.Logger, exampleUsecase)

	route := httproute.Route{
		App:                   config.App,
		HealthCheckController: healtcheckController,
		LogController:         logController,
		ExampleController:     exampleController,
	}

	route.Setup()

	return nil
}
