package prototype

type NodeInterface interface {
	Print(string)
	Clone() NodeInterface
}
