package ast_test

import (
	"testing"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewComponentStatement(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *ast.ComponentStatement
	}{
		// TODO: Add test cases.
	}

	SkipConvey("NewComponentStatement()", t, func() {
		for _, tt := range tests {
			Convey(tt.name, func() {
				c := ast.NewComponentStatement(tt.args.name)

				So(c, ShouldResemble, tt.want)
			})
		}
	})

}
