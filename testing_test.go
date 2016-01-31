package candy

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestData(t *testing.T) {
	assert.Equal(t,
		path.Join(GoPath(), "src/github.com/topicai/candy/testdata"),
		TestData())
}
