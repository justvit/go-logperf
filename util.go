package main

import (
	"runtime"
	"time"
)

func GetCallerInfo() (file string, line int, function string) {
	fpcs := make([]uintptr, 1)
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return "", -1, "" // proper error here would be better
	}
	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return "", -1, ""  // proper error here would be better
	}

	file, line = fun.FileLine(fun.Entry())
	return file, line, fun.Name()
}

// Usage: defer app.LogPerf()()
// Explanation: the function is called at the start of the outer function's call,
//              memorises the moment
//              and returns a function that is called (by `defer`) before the finish of the outer function's call;
//              the latter function calculates elapsed time
func (app *ExampleApp) LogPerf() func() {
	start := time.Now()
	_, _, caller := GetCallerInfo()

	logger.Printf("started %s(...)", caller)
	return func() {
		finish := time.Now()
		elapsed := finish.UnixNano() - start.UnixNano()
		logger.Printf("finished %s(...), elapsed: %d ns", caller, elapsed)
	}
}


