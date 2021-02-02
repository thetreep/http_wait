package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"net/http"
	"os"
	"time"
)

func main() {
	var flagAlias = map[string]string{
		"uri":           "u",
		"port":          "p",
		"timeout":       "t",
		"interval":      "i",
		"response-code": "rc",
	}

	var endpoint = flag.String("uri", "http://localhost", "URI, with protocol and no port.")
	var port = flag.Int("port", 80, "Port to use.")
	var timeout = flag.Int("timeout", 10000, "Timeout before returning a non 0 exit code, in milliseconds.")
	var interval = flag.Int("interval", 1000, "Polling interval, in milliseconds.")
	var responseCode = flag.Int("response-code", 200, "Response code to wait for.")

	for from, to := range flagAlias {
		flagSet := flag.Lookup(from)
		flag.Var(flagSet.Value, to, fmt.Sprintf("alias to %s", flagSet.Name))
	}

	flag.Parse()

	startTime := time.Now()
	timeoutDuration := time.Duration(*timeout) * time.Millisecond
	sleepDuration := time.Duration(*interval) * time.Millisecond

	ticker := time.NewTicker(sleepDuration)

	for {
		res, err := http.Head(fmt.Sprintf("%s:%d", *endpoint, *port))
		if err == nil && res.StatusCode == *responseCode {
			os.Exit(0)
		}
		timeElapsed := time.Now().Sub(startTime)
		if timeElapsed > timeoutDuration {
			os.Exit(1)
		}
		<-ticker.C
	}
}
