package errors

import (
	err "errors"
	"fmt"
	"strings"
)

var (
	LanguageError Kind = "Language"
)

type (
	Kind string

	Error interface {
		Error() error
		String() string
		Kind() Kind
		SetKind(Kind) Error
		SubKind() Kind
		SetSubKind(Kind) Error
		Message() string
		SetMessage(string) Error
		Params() map[string]interface{}
		AddParam(name string, value interface{}) Error
		SetParams(map[string]interface{}) Error
	}

	ErrorObject struct {
		kind    Kind
		subKind Kind
		message string
		params  map[string]interface{}
	}

	Errors map[string]error
)

// String -
func (k Kind) String() string {
	return string(k)
}

// New -
func New(kind Kind) Error {
	return ErrorObject{
		kind:   kind,
		params: make(map[string]interface{}),
	}
}

func (e ErrorObject) String() string {
	var msg strings.Builder

	msg.WriteString("[")
	msg.WriteString(e.kind.String())
	msg.WriteString(" error")
	msg.WriteString("]")

	if e.subKind != "" {
		msg.WriteString(" ")
		msg.WriteString(e.subKind.String())
	}

	if e.message != "" {

		if e.subKind != "" {
			msg.WriteString(" -")
		}

		msg.WriteString(" ")

		msg.WriteString(e.message)
	}

	if len(e.params) == 0 {
		return msg.String()
	} else {

		if e.message != "" {
			msg.WriteString(", ")
		} else {
			if e.subKind != "" {
				msg.WriteString(" -")
			}

			msg.WriteString(" ")
		}

		params := make([]string, len(e.params)/2)
		for key, value := range e.params {
			formatted := fmt.Sprintf("%v=%v", key, value)
			params = append(params, formatted)
		}

		msg.WriteString(strings.Join(params, " "))
	}
	return msg.String()
}

func (e ErrorObject) Error() error {
	str := e.String()
	return err.New(str)
}

func (e ErrorObject) Kind() Kind {
	return e.kind
}

func (e ErrorObject) SetKind(k Kind) Error {
	e.kind = k
	return e
}

func (e ErrorObject) SubKind() Kind {
	return e.subKind
}

func (e ErrorObject) SetSubKind(k Kind) Error {
	e.subKind = k
	return e
}

func (e ErrorObject) Message() string {
	return e.message
}

func (e ErrorObject) SetMessage(msg string) Error {
	e.message = msg
	return e
}

func (e ErrorObject) Params() map[string]interface{} {
	return e.params
}

func (e ErrorObject) AddParam(name string, value interface{}) Error {
	e.params[name] = value
	return e
}

func (e ErrorObject) SetParams(m map[string]interface{}) Error {
	for k, v := range m {
		e.AddParam(k, v)
	}
	return e
}
