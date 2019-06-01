package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/aerogo/aero"
)

// Recover recovers from panics and shows them as the response body.
func Recover(next aero.Handler) aero.Handler {
	return func(ctx aero.Context) error {
		defer func() {
			r := recover()

			if r == nil {
				return
			}

			err, ok := r.(error)

			if !ok {
				err = fmt.Errorf("%v", r)
			}

			stack := make([]byte, 4096)
			length := runtime.Stack(stack, true)
			_ = ctx.Error(http.StatusInternalServerError, err, stack[:length])
		}()

		return next(ctx)
	}
}
