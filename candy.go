package candy

import (
	"io"
	"log"

	"github.com/wangkuiyi/fs"
)

// Must panics if e is not nil.
func Must(e error) {
	if e != nil {
		log.Panic(e)
	}
}

// MustOpen panics if it cannot open file.
func MustOpen(file string) io.ReadCloser {
	f, e := fs.Open(file)
	Must(e)
	return f
}

// MustCreate panics if it cannot create file.
func MustCreate(fileanme string) io.WriteCloser {
	f, e := fs.Create(fileanme)
	Must(e)
	return f
}

// OpenAnd calls fn if it can open file, or panics.  For example, if
// we want to decode a JSON struct from file /user/yi/hello on HDFS,
// we can do:
//
//    s := OpenAnd("/hdfs/user/yi/hello", func (r io.Reader) interface{} {
//            s := new(AStruct)
//            Must(json.NewDecoder(r).Decode(s))
//            return s
//    }).(*AStruct)
func OpenAnd(file string, fn func(r io.Reader) interface{}) interface{} {
	f := MustOpen(file)
	defer f.Close()
	return fn(f)
}

// CreateAnd calls fn if it can create file, or panics.  For example,
// if we want to encode a struct in JSON format into a file
// /user/yi/hello located on HDFS, we can do:
//
//    var s AStruct
//    CreateAnd("/hdfs/user/yi/hello", func (w io.Writer) {
//        Must(json.NewEncoder(w).Encode(s))
//    })
func CreateAnd(file string, fn func(w io.Writer)) {
	f := MustCreate(file)
	defer f.Close()
	fn(f)
}
