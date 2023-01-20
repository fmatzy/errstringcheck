package errstringcheck

import (
	"log"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestErrorf(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), NewAnalyzer(), "errorf")
}

func TestWraponly(t *testing.T) {
	analyzer := NewAnalyzer()
	err := analyzer.Flags.Set("wraponly", "true")
	if err != nil {
		log.Fatal(err)
	}
	analysistest.Run(t, analysistest.TestData(), analyzer, "wraponly")
}
