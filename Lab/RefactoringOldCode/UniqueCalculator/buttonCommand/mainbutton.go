package mbutton

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

type calculator interface {
	add()
	multiply()
	divide()
	sub()
}
