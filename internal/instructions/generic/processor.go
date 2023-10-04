package generic

// Processor 指令处理器
type Processor interface {
	GetToken() (token *string)
	Execute()
}
