package zerologadapter

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
	"github.com/rs/zerolog"
)

func newTestSuite() *loggertesting.LoggerTestSuite {
	return &loggertesting.LoggerTestSuite{
		TraceFallbackToDebug: true,
		LoggerFactory: func(level Level) (Logger, func() []LogEvent) {
			var buf bytes.Buffer
			logger := zerolog.New(&buf).Level(zerolog.Level(level))

			return New(logger), func() []LogEvent {
				lines := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")

				events := make([]LogEvent, len(lines))

				for key, line := range lines {
					var event map[string]interface{}

					err := json.Unmarshal([]byte(line), &event)
					if err != nil {
						panic(err)
					}

					level, _ := ParseLevel(strings.ToLower(event["level"].(string)))
					msg := event["message"].(string)

					delete(event, "level")
					delete(event, "message")

					events[key] = LogEvent{
						Line:   msg,
						Level:  level,
						Fields: event,
					}
				}

				return events
			}
		},
	}
}

func TestLoggerSuite(t *testing.T) {
	newTestSuite().Execute(t)
}
