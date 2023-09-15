package infrastructures

import (
	"golang-starter-project/configs"
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	Logger     zerolog.Logger
	loggerOnce sync.Once
)

func InitLogger() {
	loggerOnce.Do(func() {
		Logger = zerolog.New(os.Stdout).
			With().
			Timestamp().
			Str("app-name", configs.GetAppConfig("APP_NAME")).
			Logger()
	})
}
