package logger
import (
	"github.com/rs/zerolog"
	"github.com/midil-labs/core/pkg/config"
)


type NRHook struct{
	config *config.LoggingConfig
}

func NewNRHook(config *config.LoggingConfig) *NRHook {
	return &NRHook{config: config}
}


func (h *NRHook) NewRelicHook() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(h.config.NewRelicAppName),
		newrelic.ConfigLicense(h.config.NewRelicLicenseKey),
	)
	if err != nil {
		log.Fatalf("failed to create New Relic application: %v", err)
	}

	h.config.NewRelicApp = app
}

func (h *NRHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if h.config.NewRelicApp != nil {
		h.config.NewRelicApp.RecordCustomEvent("log", map[string]interface{}{
			"level": e.Str("level").String(),
			"message": e.Str("message").String(),
			"timestamp": e.Time("time").Time(),
		})
	}
}
