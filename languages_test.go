package execman

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorsWithContainers(t *testing.T) {
	_, err := New(Options{
		Ruby: Ruby{
			Workers:    0,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, InvalidOptions))
}

func TestErrorsWithWorkers(t *testing.T) {
	_, err := New(Options{
		Ruby: Ruby{
			Workers:    10,
			Containers: 0,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, InvalidOptions))
}

func TestMultipleExecmans(t *testing.T) {
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

	em, err = New(Options{
		Ruby: Ruby{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res = em.Run(RubyLang, `puts "Hello world"`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}

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

func TestHaskellLanguage(t *testing.T) {
	em, err := New(Options{
		Haskell: Haskell{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(HaskellLang, `main = putStrLn "Hello world"`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}

func TestCSharpLanguage(t *testing.T) {
	em, err := New(Options{
		CSharp: CSharp{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(CSharpLang, `
using System;

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Hello world");
    }
}
`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}

func TestPHP74Language(t *testing.T) {
	em, err := New(Options{
		Php74: Php74{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(PHP74Lang, `
<?php
echo "Hello world";
?>
`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "\nHello world")

	em.Close()
}

func TestRustLanguage(t *testing.T) {
	em, err := New(Options{
		Rust: Rust{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(RustLang, `
fn main() {
    println!("Hello world");
}
`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world\n")

	em.Close()
}

func TestJuliaLanguage(t *testing.T) {
	em, err := New(Options{
		Julia: Julia{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	assert.Nil(t, err)

	res := em.Run(JuliaLang, `print("Hello world")`)

	assert.Nil(t, res.Error)
	assert.True(t, res.Success)
	assert.Equal(t, res.Result, "Hello world")

	em.Close()
}
