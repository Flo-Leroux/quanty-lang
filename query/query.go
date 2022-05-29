package query

func Register(root string) *Program {

	program := NewProgram(root)

	program.look()

	return program
}
