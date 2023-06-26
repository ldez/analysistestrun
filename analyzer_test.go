package analysistestrun

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testData := analysistest.TestData()

	testCases := []struct {
		desc     string
		dir      string
		patterns []string
	}{
		{
			desc:     "simple",
			dir:      "simple",
			patterns: []string{"."},
		},
		{
			desc:     "module inside a workspace",
			dir:      "workspace",
			patterns: []string{"./hello/..."},
		},
		{
			desc:     "modules inside a workspace",
			dir:      "workspace",
			patterns: []string{"./hello/...", "./world/..."},
		},
		{
			desc:     "bad module design",
			dir:      "baddesign/submodule",
			patterns: []string{"."},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			// current implementation: fail with Go modules.
			// _ = analysistest.Run(t, testData, Analyzer, test.patterns...)

			// Proposal: use my fork
			_ = analysistest.RunWithOptions(t, Analyzer, analysistest.Options{
				GOPATHMode:     false,
				Dir:            filepath.Join(testData, test.dir),
				Patterns:       test.patterns,
				SuggestedFixes: false,
			})
		})
	}
}
