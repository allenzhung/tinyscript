package ast

var _ ASTNode = &Block{}

type Program struct {
	*Block
}

func MakeProgram() *Program {
	b := &Program{MakeBlock()}
	b.SetLabel("program")
	return b
}

func ProgramParse(stream *PeekTokenStream) ASTNode {
	p := MakeProgram()
	for stmt := StmtParse(stream); nil != stmt; {
		p.AddChild(stmt)
		stmt = StmtParse(stream)
	}

	return p
}
