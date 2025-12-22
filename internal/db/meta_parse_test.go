package db

import "testing"

func TestParseMeta(t *testing.T) {
	tests := []struct {
		in string
		command_type MetaCommandType
	}{
		{".quit", MetaCommandQuit},
		{".exit", MetaCommandQuit},
		{".fake", MetaCommandQuit},
	}

	for _, tt := range tests {
		cmd := ParseMeta(tt.in)
		if cmd.Type != tt.command_type {
			t.Fatalf("input %q: expected %v, got %v", tt.in, tt.command_type, cmd.Type)
		}
	}
}
