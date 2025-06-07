package execman

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRubyLanguage(t *testing.T) {
	em, err := New(Options{
		Ruby: Ruby{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(RubyLang, `puts "Hello world"`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}

func TestCLanguage(t *testing.T) {
	em, err := New(Options{
		CLang: CLang{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(C, `
#include <stdio.h>

int main() {
	printf("Hello world");

    return 0;
}`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world")

	em.Close()
}

func TestCPlusPlusLanguage(t *testing.T) {
	t.Skip("temporary")
	em, err := New(Options{
		CPlus: CPlus{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(CPlusPlusLang, `
#include <iostream>

int main() {
    std::cout << "Hello world";
    return 0;
}`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world")

	em.Close()
}

func TestNodeLatestLanguage(t *testing.T) {
	em, err := New(Options{
		NodeLts: NodeLts{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(NodeLatestLang, `console.log('Hello world')`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}

func TestPerlLanguage(t *testing.T) {
	em, err := New(Options{
		Perl: Perl{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(PerlLtsLang, `
#!/usr/bin/perl
use warnings;
print("Hello world\n");
`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}

func TestNodeEsmLanguage(t *testing.T) {
	em, err := New(Options{
		NodeEsm: NodeEsm{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(NodeEsmLtsLang, `console.log('Hello world')`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}

func TestGoLanguage(t *testing.T) {
	em, err := New(Options{
		GoLang: GoLang{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(Golang, `
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}

func TestPython2Language(t *testing.T) {
	em, err := New(Options{
		Python2: Python2{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(Python2Lang, `print("Hello world")`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}

func TestPython3Language(t *testing.T) {
	em, err := New(Options{
		Python3: Python3{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(Python3Lang, `print("Hello world")`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}

func TestLuaLanguage(t *testing.T) {
	em, err := New(Options{
		Lua: Lua{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(LuaLang, `print("Hello world")`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}
