package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestSyncedBlockCodec_Parse(t *testing.T) {
	c := &SyncedBlockCodec{}
	tests := []struct {
		name   string
		line   string
		valid  bool
		wantID string
	}{
		{"with id", `<synced id="abc-123">`, true, "abc-123"},
		{"without id", `<synced>`, true, ""},
		{"invalid", `synced id="abc-123"`, false, ""},
		{"empty", ``, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if !tt.valid {
				return
			}
			b, ok := block.(*blocks.SyncedBlock)
			if !ok {
				t.Fatal("expected *blocks.SyncedBlock")
			}
			if b.SyncedBlock.ID != tt.wantID {
				t.Errorf("expected id=%q, got %q", tt.wantID, b.SyncedBlock.ID)
			}
		})
	}
}

func TestSyncedBlockCodec_Render(t *testing.T) {
	c := &SyncedBlockCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with id",
			&blocks.SyncedBlock{SyncedBlock: blocks.SyncedBlockData{ID: "abc-123"}},
			`<synced id="abc-123">`,
			true,
		},
		{
			"without id",
			&blocks.SyncedBlock{SyncedBlock: blocks.SyncedBlockData{}},
			`<synced>`,
			true,
		},
		{"wrong type", &blocks.ParagraphBlock{}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := c.Render(tt.block)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid && got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}