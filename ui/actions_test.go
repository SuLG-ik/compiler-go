package ui

import (
	"testing"
)

func TestActionRegistry_RegisterAndGet(t *testing.T) {
	r := newActionRegistry()
	a := &Action{
		ID:      ActionNew,
		Label:   "New",
		Handler: func() {},
	}
	r.Register(a)

	got := r.Get(ActionNew)
	if got.ID != ActionNew {
		t.Errorf("expected ID %q, got %q", ActionNew, got.ID)
	}
	if got.Label != "New" {
		t.Errorf("expected Label %q, got %q", "New", got.Label)
	}
}

func TestActionRegistry_GetUnknown(t *testing.T) {
	r := newActionRegistry()
	got := r.Get("nonexistent")
	if got == nil {
		t.Fatal("expected non-nil fallback action")
	}
	if got.Label != "nonexistent" {
		t.Errorf("expected Label %q, got %q", "nonexistent", got.Label)
	}
	if got.Handler == nil {
		t.Error("expected non-nil Handler in fallback action")
	}
}

func TestActionRegistry_OverwriteAction(t *testing.T) {
	r := newActionRegistry()
	r.Register(&Action{ID: ActionSave, Label: "first", Handler: func() {}})
	r.Register(&Action{ID: ActionSave, Label: "second", Handler: func() {}})

	got := r.Get(ActionSave)
	if got.Label != "second" {
		t.Errorf("expected overwritten Label %q, got %q", "second", got.Label)
	}
}
