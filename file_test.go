package candy

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMustWithError(t *testing.T) {
	func() {
		defer func() {
			assert.NotNil(t, recover())
		}()
		Must(errors.New("An error"))
	}()

	func() {
		defer func() {
			assert.Nil(t, recover())
		}()
		Must(nil)
	}()
}

func TestCreateAndOpen(t *testing.T) {
	filename := path.Join(os.TempDir(),
		fmt.Sprintf("github.com-augmn-common-candy-%v", time.Now().UnixNano()))

	func() {
		defer func() {
			assert.NotNil(t, recover())
		}()
		mustOpen(filename)
	}()

	func() {
		defer func() {
			assert.Nil(t, recover())
		}()
		CreateAnd(filename, func(w io.Writer) {
			fmt.Fprintf(w, "Hello!")
		})
	}()

	func() {
		defer func() {
			assert.Nil(t, recover())
		}()
		assert.Equal(t, "Hello!",
			OpenAnd(filename, func(r io.Reader) interface{} {
				var x string
				fmt.Fscanf(r, "%s", &x)
				return x
			}).(string))
	}()

	func() {
		defer func() {
			assert.NotNil(t, recover())
		}()
		mustCreate("/not-exist-dir/not-exist-file")
	}()
}
