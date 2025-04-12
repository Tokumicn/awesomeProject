package servicegroup

import (
	"context"
	"github.com/sagikazarmark/slog-shim"
	"runtime/debug"
)

// Recover is used with defer to do cleanup on panics.
// Use it like:
//
//	defer Recover(func() {})
func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		slog.Error("Recover panic recovered", slog.Any("stack", string(debug.Stack())))
	}
}

// RecoverCtx is used with defer to do cleanup on panics.
func RecoverCtx(ctx context.Context, cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		slog.ErrorContext(ctx, "%+v\n%s", p, debug.Stack())
	}
}
