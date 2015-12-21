package rss

import (
	"strings"
	"time"
	"strconv"
)

func parseTime(s string) (time.Time, error) {
	formats := []string{
		`2006-01-02 15:04:05 -07`,
		`02 Jan 2006 15:04:05 -0700`,
		"Mon, 02 Jan 06 15:04:05 -0700",
		"Mon, 02 January 2006 15:04:05 -0700",
		"Mon, _2 Jan 2006 15:04:05 MST",
		"Mon, _2 Jan 2006 15:04:05 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		`2006-01-02 15:04:05`,
		`2006-01-02T15:04:05`,
		`02/01/2006 - 15:04`,
		`Mon, 02 Jan 2006 15:04:05`,
	}

	s = strings.TrimSpace(s)
	
	var e error
	var t time.Time

	ts, e := strconv.ParseInt(s, 10, 64)
	if e == nil {
		return time.Unix(ts, 0), e
	}

	for _, format := range formats {
		t, e = time.Parse(format, s)
		if e == nil {
			return t, e
		}
	}
	
	return time.Time{}, e
}
