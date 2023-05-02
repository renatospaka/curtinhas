package chainOfResponsability

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
