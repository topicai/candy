package candy

import (
	"fmt"
	"net"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHttpGet(t *testing.T) {
	ln, e := net.Listen("tcp", ":0")
	Must(e)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})
	go func() { Must(http.Serve(ln, router)) }()

	b, e := HTTPGet("http://"+ln.Addr().String(), 0)
	assert.Nil(t, e)
	assert.Equal(t, []byte("Hello!"), b)
}
