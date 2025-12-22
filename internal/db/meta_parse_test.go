package db

import "testing"

func TestParseMeta(t *testing.T) {
	tests := []struct {
		in string
		kind MetaCommandKind
	}{
		{".quit", MetaCommandQuit},
		{".exit", MetaCommandQuit},
		{".fake", MetaCommandQuit},
	}

	for _, tt := range tests {
		cmd := ParseMeta(tt.in)
		if cmd.Kind != tt.kind {
			t.Fatalf("input %q: expected %v, got %v", tt.in, tt.kind, cmd.Kind)
		}
	}
}
