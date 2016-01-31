package candy

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/wangkuiyi/fs"
)

// Must panics if e is not nil.
func Must(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func mustOpen(file string) io.ReadCloser {
	f, e := fs.Open(file)
	Must(e)
	return f
}

func mustCreate(fileanme string) io.WriteCloser {
	f, e := fs.Create(fileanme)
	Must(e)
	return f
}

// WithOpened calls fn if it can open file, or panics.  For example, if
// we want to decode a JSON struct from file /user/yi/hello on HDFS,
// we can do:
//
//    s := WithOpened("/hdfs/user/yi/hello", func (r io.Reader) interface{} {
//            s := new(AStruct)
//            Must(json.NewDecoder(r).Decode(s))
//            return s
//    }).(*AStruct)
//
func WithOpened(file string, fn func(r io.Reader) interface{}) interface{} {
	f := mustOpen(file)
	defer f.Close()
	return fn(f)
}

// WithCreated calls fn if it can create file, or panics.  For example,
// if we want to encode a struct in JSON format into a file
// /user/yi/hello located on HDFS, we can do:
//
//    var s AStruct
//    WithCreated("/hdfs/user/yi/hello", func (w io.Writer) {
//        Must(json.NewEncoder(w).Encode(s))
//    })
//
func WithCreated(file string, fn func(w io.Writer)) {
	f := mustCreate(file)
	defer f.Close()
	fn(f)
}

// ReadAll is preferred than WithOpened for small files whose content
// can be loaded into memory.  The example in the comment of
// WithOpened can be rewritten using RealAll:
//
//    var s AStruct
//    Must(json.NewDecoder(ReadAll("/hdfs/user/yi/hello")).Decode(&s))
//
func ReadAll(file string) io.Reader {
	f := mustOpen(file)
	defer f.Close()
	buf, e := ioutil.ReadAll(f)
	Must(e)
	return bytes.NewReader(buf)
}
