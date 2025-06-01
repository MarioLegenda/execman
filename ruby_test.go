package execman

import (
	"github.com/MarioLegenda/execman/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRubyLanguage(t *testing.T) {
	em := New(Options{
		Ruby: Ruby{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})

	res := em.RunJob(string(types.Ruby.Name), `puts "Hello world"`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}
