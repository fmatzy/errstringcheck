package errstringcheck

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestErrorsAs(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "errorf")
}
