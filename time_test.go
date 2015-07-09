package rss
import (
	"testing"
	"log"
)

func Test_timeParsing(t *testing.T) {
	test := map[string]string{
		`Thu, 09 July 2015 11:37:32 +0600`: `2015-07-09 11:37:32 +0600 +0600`,
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