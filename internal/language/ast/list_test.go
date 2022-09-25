package ast_test

import (
	"testing"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	. "github.com/smartystreets/goconvey/convey"
)

// TestList -
func TestList(t *testing.T) {
	FocusConvey("Subject: List Utility - check with T = string ", t, func() {

		var strList *ast.List[string]

		Convey("New List", func() {
			strList := ast.NewList[string]()
			So(strList, ShouldHaveSameTypeAs, ast.List[string]{})
		})

		Convey("Length", func() {
			So(strList.Len(), ShouldEqual, 0)
		})

		Convey("IsEmpty", func() {
			So(strList.IsEmpty(), ShouldBeTrue)
		})

		Convey("Is not Empty", func() {
			strList.Append("Hello")
			So(strList.IsEmpty(), ShouldBeFalse)
		})

		Convey("Append item", func() {
			strList.Append("Hello")

			So(strList.Len(), ShouldEqual, 1)
			So(strList.Items(), ShouldAlmostEqual, []string{"Hello"})
		})

	})
}
