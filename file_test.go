package candy

import (
	"encoding/json"
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

func TestWithCreateAndOpened(t *testing.T) {
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
		WithCreated(filename, func(w io.Writer) {
			fmt.Fprintf(w, "Hello!")
		})
	}()

	func() {
		defer func() {
			assert.Nil(t, recover())
		}()
		assert.Equal(t, "Hello!",
			WithOpened(filename, func(r io.Reader) interface{} {
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

func TestReadAll(t *testing.T) {
	m := make(map[string]string)
	Must(json.NewDecoder(ReadAll(TestData("example.json"))).Decode(&m))
	assert.Equal(t, "Yi", m["Name"])
	assert.Equal(t, "36", m["Age"])
}
