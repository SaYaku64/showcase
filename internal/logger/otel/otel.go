/* 
MOCK FOR THE OTEL

It can be go.opentelemetry.io/otel or github.com/tel-io/tel/v2, it doesn't matter
*/

package otel

import (
	li "github.com/SaYaku64/showcase/internal/logger"
)

type otelModule struct {}

func Init() li.ILogger {
	return &otelModule{}
}

// Info mock
func (l *otelModule) Info(format string, v ...any) {}

// Error mock
func (l *otelModule) Error(format string, v ...any) {}

// Warning mock
func (l *otelModule) Warning(format string, v ...any) {}

// Fatal mock
func (l *otelModule) Fatal(format string, v ...any) {
}
