package db

import "strings"

type MetaCommand struct {
	Kind MetaCommandKind
	Raw string
}
type MetaCommandKind int
func (k MetaCommandKind) String() string {
	switch k {
	case MetaCommandQuit:
		return "quit"
	default:
		return "unknown"
	}
}
const (
	MetaCommandUnknown MetaCommandKind = iota
	MetaCommandQuit
)

func ParseMeta(input string) MetaCommand {
	raw := strings.TrimSpace(input)
	if raw == "" {
		return MetaCommand{
			Kind: MetaCommandUnknown,
			Raw: input,
		}
	}
	if raw[0] != '.' {
		return MetaCommand{
			Kind: MetaCommandUnknown,
			Raw: input,
		}
	}

	switch raw {
	case ".quit", ".exit":
		return MetaCommand{
			Kind: MetaCommandQuit,
			Raw: raw,
		}
	default: 
		return MetaCommand{
			Kind: MetaCommandUnknown,
			Raw: input,
		}
	}
}
