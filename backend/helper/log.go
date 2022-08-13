package helper

import (
	"os"

	"github.com/rs/zerolog"
)

// Logger is the global logger.
// can be customized
var Logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()

func init() {
	zerolog.TimestampFieldName = "timestamp"
}
