package validate

import (
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/animenotifier/notify.moe/arn/autocorrect"
)

const (
	// DateFormat is the format used for short dates that don't include the time.
	DateFormat = "2006-01-02"

	// YearMonthFormat is the format used for validating dates that include the year and month.
	YearMonthFormat = "2006-01"

	// DateTimeFormat is the format used for long dates that include the time.
	DateTimeFormat = time.RFC3339
)

var (
	discordNickRegex = regexp.MustCompile(`^([^#]{2,32})#(\d{4})$`)
	emailRegex       = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// Nick tests if the given nickname is valid.
func Nick(nick string) bool {
	if len(nick) < 2 {
		return false
	}

	if !utf8.ValidString(nick) {
		return false
	}

	return nick == autocorrect.UserNick(nick)
}

// DiscordNick tests if the given Discord nickname is valid.
func DiscordNick(nick string) bool {
	return discordNickRegex.MatchString(nick)
}

// DateTime tells you whether the datetime is valid.
func DateTime(date string) bool {
	if date == "" || strings.HasPrefix(date, "0001") {
		return false
	}

	_, err := time.Parse(DateTimeFormat, date)
	return err == nil
}

// Date tells you whether the datetime is valid.
func Date(date string) bool {
	if date == "" || strings.HasPrefix(date, "0001") {
		return false
	}

	_, err := time.Parse(DateFormat, date)
	return err == nil
}

// YearMonth tells you whether the date contains only the year and the month.
func YearMonth(date string) bool {
	if len(date) != len(YearMonthFormat) || strings.HasPrefix(date, "0001") {
		return false
	}

	_, err := time.Parse(YearMonthFormat, date)
	return err == nil
}

// Email tests if the given email address is valid.
func Email(email string) bool {
	return emailRegex.MatchString(email)
}

// URI validates a URI.
func URI(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	return err == nil
}
