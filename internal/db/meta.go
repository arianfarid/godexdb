package db

import "strings"

type MetaCommand struct {
	Type MetaCommandType
	Raw string
}
type MetaCommandType int
func (k MetaCommandType) String() string {
	switch k {
	case MetaCommandQuit:
		return "quit"
	default:
		return "unknown"
	}
}
const (
	MetaCommandUnknown MetaCommandType = iota
	MetaCommandQuit
)

func ParseMeta(input string) MetaCommand {
	raw := strings.TrimSpace(input)
	if raw == "" {
		return MetaCommand{
			Type: MetaCommandUnknown,
			Raw: input,
		}
	}
	if raw[0] != '.' {
		return MetaCommand{
			Type: MetaCommandUnknown,
			Raw: input,
		}
	}

	switch raw {
	case ".quit", ".exit":
		return MetaCommand{
			Type: MetaCommandQuit,
			Raw: raw,
		}
	default: 
		return MetaCommand{
			Type: MetaCommandUnknown,
			Raw: input,
		}
	}
}
