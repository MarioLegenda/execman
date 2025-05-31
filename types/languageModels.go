package types

type Name string
type Text string
type Tag string
type InDevelopment bool
type InMaintenance bool

type Language struct {
	Name          Name          `json:"name"`
	Text          Text          `json:"text"`
	Tag           Tag           `json:"tag"`
	InDevelopment InDevelopment `json:"inDevelopment"`
	InMaintenance InMaintenance `json:"inMaintenance"`
	Language      string        `json:"language"`
	Extension     string        `json:"extension"`
}

var NodeLts = Language{
	Name:          "node_latest",
	Text:          "Javascript (Node latest)",
	Tag:           "node:node_latest",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "javascript",
	Extension:     "js",
}

var PerlLts = Language{
	Name:          "perl",
	Text:          "Perl (latest)",
	Tag:           "perl:perl",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "perl",
	Extension:     "pl",
}

var NodeEsm = Language{
	Name:          "node_latest_esm",
	Text:          "Javascript (Node ESM)",
	Tag:           "node:node_latest_esm",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "javascript",
	Extension:     "mjs",
}

var GoLang = Language{
	Name:          "go",
	Text:          "Go v1.18",
	Tag:           "go:go_v18",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "go",
	Extension:     "go",
}

var Python2 = Language{
	Name:          "python2",
	Text:          "Python2",
	Tag:           "python:python2",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "python",
	Extension:     "py",
}

var Python3 = Language{
	Name:          "python3",
	Text:          "Python3",
	Tag:           "python:python3",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "python",
	Extension:     "py",
}

var Lua = Language{
	Name:          "lua",
	Text:          "Lua",
	Tag:           "lua:lua",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "lua",
	Extension:     "lua",
}

var Ruby = Language{
	Name:          "ruby",
	Text:          "Ruby",
	Tag:           "ruby:ruby",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "ruby",
	Extension:     "rb",
}

var Php74 = Language{
	Name:          "php74",
	Text:          "PHP 7.4",
	Tag:           "php:php7.4",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "php",
	Extension:     "php",
}

var Rust = Language{
	Name:          "rust",
	Text:          "Rust",
	Tag:           "rust:rust",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "rust",
	Extension:     "rs",
}

var Haskell = Language{
	Name:          "haskell",
	Text:          "Haskell",
	Tag:           "haskell:haskell",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "haskell",
	Extension:     "hs",
}

var CLang = Language{
	Name:          "c",
	Text:          "C",
	Tag:           "c:c",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "c",
	Extension:     "c",
}

var CPlus = Language{
	Name:          "c++",
	Text:          "C++",
	Tag:           "c-plus:c-plus",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "c",
	Extension:     "cpp",
}

var CSharpMono = Language{
	Name:          "c_sharp_mono",
	Text:          "C# (Mono)",
	Tag:           "c_sharp_mono:c_sharp_mono",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "csharp",
	Extension:     "cs",
}

var Julia = Language{
	Name:          "julia",
	Text:          "Julia",
	Tag:           "julia:julia",
	InDevelopment: false,
	InMaintenance: false,
	Language:      "julia",
	Extension:     "jl",
}

type SingleFileBuildResult struct {
	ContainerName      string
	DirectoryName      string
	ExecutionDirectory string
	FileName           string
	Environment        *Language
	StateDirectory     string
	Timeout            int
	Args               []string
}
