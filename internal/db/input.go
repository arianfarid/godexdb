package db

import (
	"strings"
)

type InputType int

const (
	InputUnknown InputType = iota
	InputEmpty
	InputMeta
	InputSQL
)

type Input struct {
	Type InputType
	Meta *MetaCommand
	SQL string
}

func ParseInput(input string) Input {
	trimmed := strings.TrimSpace(input[:len(input) - 1])

	if trimmed == "" {
		return Input{
			Type: InputEmpty,
		}
	}
	if strings.HasPrefix(trimmed, ".") {
		meta := ParseMeta(trimmed)
		return Input{
			Type: InputMeta,
			Meta: &meta,
		}
	}

	return Input{
		Type: InputSQL,
		SQL: trimmed,
	}
}
