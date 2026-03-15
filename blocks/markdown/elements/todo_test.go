package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestToDoCodec_Parse(t *testing.T) {
	c := &ToDoCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"unchecked", `- [ ] Task to do`, true},
		{"checked should not match", `- [x] Done task`, false},
		{"invalid", `[ ] Task`, false},
		{"empty", ``, false},
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
			b, ok := block.(*blocks.ToDoBlock)
			if !ok {
				t.Fatal("expected *blocks.ToDoBlock")
			}
			if b.ToDo.Checked {
				t.Error("expected Checked=false")
			}
		})
	}
}

func TestToDoCodec_Render(t *testing.T) {
	c := &ToDoCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"unchecked",
			&blocks.ToDoBlock{ToDo: blocks.ToDoData{RichText: toRichText("Task to do"), Checked: false}},
			"- [ ] Task to do",
			true,
		},
		{
			"checked should not render",
			&blocks.ToDoBlock{ToDo: blocks.ToDoData{RichText: toRichText("Done"), Checked: true}},
			"",
			false,
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

func TestToDoDoneCodec_Parse(t *testing.T) {
	c := &ToDoDoneCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"checked lowercase", `- [x] Done task`, true},
		{"checked uppercase", `- [X] Done task`, true},
		{"unchecked should not match", `- [ ] Task`, false},
		{"empty", ``, false},
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
			b, ok := block.(*blocks.ToDoBlock)
			if !ok {
				t.Fatal("expected *blocks.ToDoBlock")
			}
			if !b.ToDo.Checked {
				t.Error("expected Checked=true")
			}
		})
	}
}

func TestToDoDoneCodec_Render(t *testing.T) {
	c := &ToDoDoneCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"checked",
			&blocks.ToDoBlock{ToDo: blocks.ToDoData{RichText: toRichText("Done task"), Checked: true}},
			"- [x] Done task",
			true,
		},
		{
			"unchecked should not render",
			&blocks.ToDoBlock{ToDo: blocks.ToDoData{RichText: toRichText("Task"), Checked: false}},
			"",
			false,
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