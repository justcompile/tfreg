package internal

import (
	"context"

	"github.com/justcompile/tfreg/internal/logging"
)

// Application contains properties which are shared throughout the codebase such
// as Configuration, & Database connection
type Application struct {
	Config *Config
	logger *logging.Logger
}

// Init runs startup tasks for the application such as connecting to the database
func (app *Application) Init() error {
	return nil
}

// LogWithContext permits the ability to write log messages which include the current request id for traceability
// This method returns a new instance
func (app *Application) LogWithContext(ctx context.Context) *logging.Logger {
	return app.logger.WithContext(ctx)
}

func (app *Application) SetLogger(stream logging.Stream) {
	app.logger = logging.New(stream)
}

// Shutdown should be called when the application exits in order to stop, close
// or simply tidy up anything that needs to be tidied
func (app *Application) Shutdown() error {
	return nil
}

var applicationSingleton *Application

// NewApplication uses a singleton instance of Application in order to be utilised throughout
// the app
func NewApplication(stream logging.Stream) (*Application, error) {
	if applicationSingleton == nil {
		cfg, err := NewConfig()
		if err != nil {
			return nil, err
		}

		applicationSingleton = &Application{
			Config: cfg,
			logger: logging.New(stream),
		}
	}

	return applicationSingleton, nil
}
