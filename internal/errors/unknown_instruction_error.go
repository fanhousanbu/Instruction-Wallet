package errors

import "fmt"

// UnknownInstructionError 未知指令
type UnknownInstructionError struct {
	InstructionName string
}

func (e UnknownInstructionError) Error() string {
	return fmt.Sprintf("unknown instruction: %s", e.InstructionName)
}
