package errors_test

import (
	"testing"

	"github.com/Flo-Leroux/quanty-lang/internal/errors"
	. "github.com/smartystreets/goconvey/convey"
)

// TestToken -
func TestErrors(t *testing.T) {
	Convey("Create: LanguageError", t, func() {
		err := errors.New(errors.LanguageError)

		validateError := func(err errors.Error) {
			So(err.Error(), ShouldBeError)
			So(err.Error().Error(), ShouldEqual, err.String())
		}

		Convey("Kind is: LanguageError", func() {
			So(err.Kind(), ShouldEqual, errors.LanguageError)
			So(err.String(), ShouldEqual, "[Language error]")
			validateError(err)
		})

		Convey("Change Kind: with custom Kind = ValidationError", func() {
			err = err.SetKind("ValidationError")
			So(err.Kind(), ShouldEqual, "ValidationError")
			So(err.String(), ShouldEqual, "[ValidationError error]")
			validateError(err)
		})

		Convey("Add subKind: Parser", func() {
			err = err.SetSubKind("Parser")
			So(err.SubKind(), ShouldEqual, "Parser")
			So(err.String(), ShouldEqual, `[Language error] Parser`)
			validateError(err)
		})

		Convey("Add message: Unexpected token X", func() {
			err = err.SetMessage("Unexpected token X")
			So(err.Message(), ShouldEqual, "Unexpected token X")
			So(err.String(), ShouldEqual, "[Language error] Unexpected token X")
			validateError(err)
		})

		Convey("Add subKind: Parser; and message: Unexpected token X", func() {
			err = err.SetSubKind("Parser")
			err = err.SetMessage("Unexpected token X")
			So(err.Message(), ShouldEqual, "Unexpected token X")
			So(err.String(), ShouldEqual, "[Language error] Parser - Unexpected token X")
			validateError(err)
		})

		Convey("Add Param: info=arg", func() {
			err = err.AddParam("info", "arg")
			So(err.Params(), ShouldResemble, map[string]interface{}{"info": "arg"})
			So(err.String(), ShouldEqual, "[Language error] info=arg")
			validateError(err)
		})

		Convey("Set Params: info=arg debug=task-2", func() {
			err = err.SetParams(map[string]interface{}{
				"info":  "arg",
				"debug": "task-2",
			})
			So(err.Params(), ShouldResemble, map[string]interface{}{"info": "arg", "debug": "task-2"})
		})

		Convey("Add subKind: Parser; and params: info=arg", func() {
			err = err.SetSubKind("Parser")
			err = err.AddParam("info", "arg")

			So(err.SubKind(), ShouldEqual, "Parser")
			So(err.String(), ShouldEqual, "[Language error] Parser - info=arg")
			validateError(err)
		})

		Convey("Add message: Unexpected token X; and params: info=arg", func() {
			err = err.SetMessage("Unexpected token X")
			err = err.AddParam("info", "arg")

			So(err.Message(), ShouldEqual, "Unexpected token X")
			So(err.String(), ShouldEqual, "[Language error] Unexpected token X, info=arg")
			validateError(err)
		})

		Convey("Add subKind: Parser; and message: Unexpected token X; and params: info=arg", func() {
			err = err.SetSubKind("Parser")
			err = err.SetMessage("Unexpected token X")
			err = err.AddParam("info", "arg")

			So(err.SubKind(), ShouldEqual, "Parser")
			So(err.Message(), ShouldEqual, "Unexpected token X")
			So(err.String(), ShouldEqual, "[Language error] Parser - Unexpected token X, info=arg")
			validateError(err)
		})
	})
}
