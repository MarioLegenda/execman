package types

type Language struct {
	Name      string
	Tag       string
	Extension string
}

var NodeLts = Language{
	Name:      "node_latest",
	Tag:       "node:node_latest",
	Extension: "js",
}

var PerlLts = Language{
	Name:      "perl",
	Tag:       "perl:perl",
	Extension: "pl",
}

var NodeEsm = Language{
	Name:      "node_latest_esm",
	Tag:       "node:node_latest_esm",
	Extension: "mjs",
}

var GoLang = Language{
	Name:      "go",
	Tag:       "go:go_latest",
	Extension: "go",
}

var Python2 = Language{
	Name:      "python2",
	Tag:       "python:python2",
	Extension: "py",
}

var Python3 = Language{
	Name:      "python3",
	Tag:       "python:python3",
	Extension: "py",
}

var Lua = Language{
	Name:      "lua",
	Tag:       "lua:lua",
	Extension: "lua",
}

var Ruby = Language{
	Name:      "ruby",
	Tag:       "ruby:ruby",
	Extension: "rb",
}

var Php74 = Language{
	Name:      "php74",
	Tag:       "php:php7.4",
	Extension: "php",
}

var JavaLts = Language{
	Name:      "java",
	Tag:       "java:java_latest",
	Extension: "java",
}

var Rust = Language{
	Name:      "rust",
	Tag:       "rust:rust",
	Extension: "rs",
}

var Haskell = Language{
	Name:      "haskell",
	Tag:       "haskell:haskell",
	Extension: "hs",
}

var CLang = Language{
	Name:      "c",
	Tag:       "c:c",
	Extension: "c",
}

var CPlus = Language{
	Name:      "c++",
	Tag:       "c-plus:c-plus",
	Extension: "cpp",
}

var CSharpMono = Language{
	Name:      "c_sharp_mono",
	Tag:       "c_sharp_mono:c_sharp_mono",
	Extension: "cs",
}

var Julia = Language{
	Name:      "julia",
	Tag:       "julia:julia",
	Extension: "jl",
}

var Dart = Language{
	Name:      "dart",
	Tag:       "dart:dart",
	Extension: "dart",
}

var KotlinLts = Language{
	Name:      "kotlin",
	Tag:       "kotlin:kotlin",
	Extension: "kt",
}

var ZigLts = Language{
	Name:      "zig",
	Tag:       "zig:zig",
	Extension: "zig",
}

var Bash = Language{
	Name:      "bash",
	Tag:       "bash:bash",
	Extension: "sh",
}

var Php8 = Language{
	Name:      "php8",
	Tag:       "php8:php8",
	Extension: "php",
}
