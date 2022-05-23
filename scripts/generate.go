package main

import (
	"fmt"
	"go/build"
	"os"

	"github.com/princjef/gomarkdoc"
	"github.com/princjef/gomarkdoc/lang"
	"github.com/princjef/gomarkdoc/logger"
)

func main() {
	// Create a renderer to output data
	out, err := gomarkdoc.NewRenderer()
	if err != nil {
		panic(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	wd = wd + "/pkg/thirdweb/"
	fmt.Println(wd)

	buildPkg, err := build.ImportDir(wd, build.ImportComment)
	if err != nil {
		panic(err)
	}

	// Create a documentation package from the build representation of our
	// package.
	log := logger.New(logger.DebugLevel)
	pkg, err := lang.NewPackageFromBuild(log, buildPkg)

	file := lang.NewFile("", "", []*lang.Package{pkg})
	if err != nil {
		panic(err)
	}

	// Write the documentation out to console.
	output, err := out.File(file)
	if err != nil {
		panic(err)
	}

	os.WriteFile("docs/test.md", []byte(output), 0644)
}
