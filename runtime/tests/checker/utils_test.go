package checker

import (
	"testing"

	"github.com/dapperlabs/cadence/runtime/common"
	"github.com/dapperlabs/cadence/runtime/sema"
	"github.com/dapperlabs/cadence/runtime/stdlib"
	. "github.com/dapperlabs/cadence/runtime/tests/utils"
)

func ParseAndCheckWithPanic(t *testing.T, code string) (*sema.Checker, error) {
	return ParseAndCheckWithOptions(t,
		code,
		ParseAndCheckOptions{
			Options: []sema.Option{
				sema.WithPredeclaredValues(
					stdlib.StandardLibraryFunctions{
						stdlib.PanicFunction,
					}.ToValueDeclarations(),
				),
			},
		},
	)
}
func ParseAndCheckWithAny(t *testing.T, code string) (*sema.Checker, error) {
	return ParseAndCheckWithOptions(t,
		code,
		ParseAndCheckOptions{
			Options: []sema.Option{
				sema.WithPredeclaredTypes(map[string]sema.TypeDeclaration{
					"Any": stdlib.StandardLibraryType{
						Name: "Any",
						Type: &sema.AnyType{},
						Kind: common.DeclarationKindType,
					},
				}),
			},
		},
	)
}
