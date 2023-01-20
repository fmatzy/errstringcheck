package errstringcheck

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/ssa"
)

var errType = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

func NewAnalyzer() *analysis.Analyzer {
	r := &runner{}
	a := &analysis.Analyzer{
		Name: "errstringcheck",
		Doc:  "errstringcheck check error message format",
		Run:  r.run,
		Requires: []*analysis.Analyzer{
			buildssa.Analyzer,
		},
	}
	a.Flags.BoolVar(&r.wrapOnly, "wraponly", false, "only allow use of %w verb for formatting errors")

	return a
}

type runner struct {
	wrapOnly bool
}

func (r *runner) run(pass *analysis.Pass) (interface{}, error) {
	funcs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs

	for _, f := range funcs {
		for _, b := range f.Blocks {
			for _, inst := range b.Instrs {
				if isInvalidErrorf(pass, inst, r.wrapOnly) {
					msg := `invalid format for fmt.Errorf. Use "...: %%v" or "...: %%w" to format errors`
					if r.wrapOnly {
						msg = `invalid format for fmt.Errorf. Use "...: %%w" to format errors`
					}
					pass.Reportf(inst.Pos(), msg)
				}
			}
		}
	}

	return nil, nil
}

func isInvalidErrorf(pass *analysis.Pass, inst ssa.Instruction, wrapOnly bool) bool {
	call, ok := inst.(*ssa.Call)
	if !ok {
		return false
	}

	if !isCallFmtErrorf(call) {
		return false
	}

	if len(call.Call.Args) != 2 {
		return false
	}

	formatStr, ok := getFormatStr(call.Call.Args[0])
	if !ok {
		return false
	}

	if !hasErrArg(pass, call.Pos()) {
		return false
	}

	if strings.HasSuffix(formatStr, ": %w") {
		return false
	}
	if !wrapOnly && strings.HasSuffix(formatStr, ": %v") {
		return false
	}
	return true
}

func isCallFmtErrorf(call *ssa.Call) bool {
	f := call.Common().StaticCallee()
	if f == nil {
		return false
	}

	return f.Pkg.Pkg.Path() == "fmt" && f.Name() == "Errorf"
}

func getFormatStr(v ssa.Value) (string, bool) {
	format, ok := v.(*ssa.Const)
	if !ok {
		return "", false
	}

	if format.Value.Kind() != constant.String {
		return "", false
	}

	return constant.StringVal(format.Value), true
}

func hasErrArg(pass *analysis.Pass, pos token.Pos) bool {
	file := getFile(pass.Files, pos)
	if file == nil {
		return false
	}

	path, exact := astutil.PathEnclosingInterval(file, pos, pos)
	if !exact || len(path) == 0 {
		return false
	}

	callExpr, ok := path[0].(*ast.CallExpr)
	if !ok {
		return false
	}

	if callExpr.Ellipsis != token.NoPos {
		return false
	}

	if len(callExpr.Args) < 2 {
		return false
	}

	for _, arg := range callExpr.Args[1:] {
		typ := pass.TypesInfo.TypeOf(arg)
		if types.Implements(typ, errType) {
			return true
		}
	}
	return false
}

func getFile(fs []*ast.File, pos token.Pos) *ast.File {
	for i := range fs {
		if fs[i].Pos() <= pos && pos <= fs[i].End() {
			return fs[i]
		}
	}
	return nil
}
