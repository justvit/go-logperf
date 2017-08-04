package main

import (
	"log"
	"os"
	"math/rand"
	"time"
)

var logger *log.Logger

const APP_NAME = "logperf"

func init() {
	logger = log.New(os.Stdout, APP_NAME + " ", log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

type ExampleApp struct {
}

func (app *ExampleApp) work(n int) {
	defer app.LogPerf()() // Here is the point!

	rand.Seed(time.Now().UnixNano())
	var r = 0.0
	for i := 0; i < n; i++ {
		r += rand.NormFloat64()
	}
	logger.Printf("### %f", r)
}

func main() {
	app := &ExampleApp{}
	app.work(10000)
}
