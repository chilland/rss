package rss
import (
	"testing"
	"log"
)

func Test_timeParsing(t *testing.T) {
	test := map[string]string{
		`Thu, 09 July 2015 11:37:32 +0600`: `2015-07-09 11:37:32 +0600 +0600`,
		`08 Jul 2015 19:04:04 +0300` : `2015-07-08 19:04:04 +0300 MSK`,
		`Thu, 09 Jul 15 10:29:29 +0400`: `2015-07-09 10:29:29 +0400 +0400`,
		`2015-12-21 00:00:00`: `2015-12-21 00:00:00 +0000 UTC`,
		`2015-12-21T07:57:18`: `2015-12-21 07:57:18 +0000 UTC`,
		`20 Dec 2015 15:49:51 +0300`: `2015-12-20 15:49:51 +0300 MSK`,
		`21/12/2015 - 10:58`: `2015-12-21 10:58:00 +0000 UTC`,
		`Mon, 21 Dec 2015 10:58:00`: `2015-12-21 10:58:00 +0000 UTC`,
		`1450641341`: `2015-12-20 22:55:41 +0300 MSK`,
	}

	for str, parsedTime := range test {
		tmp, err := parseTime(str)
		if err != nil {
			log.Println("error parsing time", err)
			t.Fail()
		} else {
			if tmp.String() != parsedTime {
				log.Println("error parsing time", str, `expected`, parsedTime, `getted`, tmp)
				t.Fail()
			}
		}
	}
}