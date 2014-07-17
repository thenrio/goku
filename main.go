package main

import (
  vegeta "github.com/thierryhenrio/vegeta/lib"
  "time"
  "strings"
  "io"
  "log"
  "os"
)

func main() {
  t := vegeta.Target{
    Method: "POST",
    URL: "http://localhost:9080/events",
    Pipe: func(i int) io.Reader { return strings.NewReader(string(i)) },
  }
  rate := uint64(10) // per second
  duration := 1 * time.Second

  results := vegeta.Attack([]vegeta.Target{ t }, rate, duration)

	log.Printf("Done! Writing results to stdout...")
  results.Encode(os.Stdout)
}
