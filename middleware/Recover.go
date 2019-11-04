package middleware

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
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
			length := runtime.Stack(stack, false)
			stackString := string(stack[:length])
			fmt.Fprint(os.Stderr, stackString)

			// Save crash in database
			crash := &arn.Crash{
				Error: err.Error(),
				Stack: stackString,
				Path:  ctx.Path(),
			}

			crash.ID = arn.GenerateID("Crash")
			crash.Created = arn.DateTimeUTC()
			user := arn.GetUserFromContext(ctx)

			if user != nil {
				crash.CreatedBy = user.ID
			}

			crash.Save()

			// Send HTML
			message := "<div class='crash'>" + err.Error() + "<br><br>" + strings.ReplaceAll(stackString, "\n", "<br>") + "</div>"
			_ = ctx.Error(http.StatusInternalServerError, message)
		}()

		return next(ctx)
	}
}
