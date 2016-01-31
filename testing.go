package candy

import (
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

// GoPath returns the first path in environment GOPATH.
func GoPath() string {
	return strings.Split(os.Getenv("GOPATH"), ":")[0]
}

// TestData returns the full path of the ./testdata sub-directory.  It
// is supposed to be used in unit tests and works like:
//
//    func TestSomethingUsingData(t *testing.T) {
//        OpenAnd(path.Join(TestData(), "dataset.txt"), func(f io.Reader) {
//            // f is the file ./testdata/dataset.txt
//        })
//    }
//
func TestData() string {
	_, f, _, ok := runtime.Caller(1)
	if !ok {
		log.Panic("runtime.Caller(1) failed")
	}
	return path.Join(path.Dir(f), "testdata")
}
