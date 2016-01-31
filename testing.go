package candy

import (
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

// When we do regression test, we often need to build a package
func GoPath() string {
	return strings.Split(os.Getenv("GOPATH"), ":")[0]
}

func TestData() string {
	_, f, _, ok := runtime.Caller(1)
	if !ok {
		log.Panic("runtime.Caller(1) failed")
	}
	return path.Join(path.Dir(f), "testdata")
}
