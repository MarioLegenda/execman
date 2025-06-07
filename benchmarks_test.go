package execman

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Benchmark1Container10Workers(b *testing.B) {
	emulator, err := New(Options{
		CPlus: CPlus{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(b, err)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = emulator.Run(CPlusPlusLang, `
#include <iostream>

int main() {
    std::cout << "Hello world";
    return 0;
}`)
	}
	b.StopTimer()
}
