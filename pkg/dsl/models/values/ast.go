package values

// String -
type String struct {
	value string
}

// NewString -
func NewString(value string) *String {
	return &String{
		value: value,
	}
}

// node -
func (sv *String) Node() {}

// statementNode -
func (sv *String) StatementNode() {}

// // Accept -
// func (sv *String) Accept(v Visitor) {
// 	v.VisitString(sv)
// }

// Value -
func (sv *String) Value() string {
	return sv.value
}
