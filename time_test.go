package rss
import (
	"testing"
	"log"
)

func Test_timeParsing(t *testing.T) {
	test := map[string]string{
		`Thu, 09 July 2015 11:37:32 +0600`: `2015-07-09 11:37:32 +0600 +0600`,
		`08 Jul 2015 19:04:04 +0300` : `2015-07-01 19:04:04 +0300 MSK`,
		`Thu, 09 Jul 15 10:29:29 +0400`: `2015-07-09 10:29:29 +0400 +0400`,
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