package main

import (
	"fmt"
	vegeta "github.com/thierryhenrio/vegeta/lib"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var f = func(i int) io.Reader {
	format := `{"date":"2013-09-16T10:00:00.000Z","device":{"ip":"10.90.1.10"},"items":[{"id":"%s"}],"state":"RETAIL_SOLD","action":"RETAIL_SELLING"}`

	options := map[string]int{
		"company":   361253,
		"reference": 42,
		"sequence":  i,
	}

  body := fmt.Sprintf(format, sgtin(options))
	log.Printf("hit --> %d yields %s\n", i, body )
	return strings.NewReader( body )
}

func main() {
  header := http.Header{
		"content-type": []string{"application/json"},
	}

	t := vegeta.Target{
		Method: "POST",
		URL:    "http://localhost:9080/events",
		Fun:   f,
		Header: header,
	}
	rate := uint64(10) // per second
	duration := 1 * time.Second

	results := vegeta.Attack([]vegeta.Target{t}, rate, duration)

	log.Printf("Done! Writing results to stdout...")
	results.Encode(os.Stdout)
}
