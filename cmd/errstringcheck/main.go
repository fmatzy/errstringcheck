package main

import (
	"github.com/fmatzy/errstringcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(errstringcheck.NewAnalyzer())
}
