package analysistestrun

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "Example",
	Doc:  "Example",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if comment, ok := n.(*ast.Comment); ok {
				if strings.HasPrefix(comment.Text, "// BUG:") {
					pass.Report(analysis.Diagnostic{
						Pos:      comment.Pos(),
						End:      0,
						Category: "example",
						Message:  "There is a bug",
					})
				}
			}

			return true
		})
	}

	return nil, nil
}
