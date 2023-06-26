package analysistestrun

import (
	"testing"

	"github.com/stretchr/testify/require"
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
			results := analysistest.Run(t, testData, Analyzer, test.patterns...)
			for _, result := range results {
				require.NotEmpty(t, result.Diagnostics)
			}
		})
	}
}
