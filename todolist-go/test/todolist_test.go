package classes_test

import (
	"testing"
	"example.com/TODOLIST-GO-OOP/classes"
)

func TestAddItem(t *testing.T) {
	tl := classes.TodoList{}
	tl.AddItem("Test1", "2024-01-01")

	if len(tl.Items) != 1 {
		t.Errorf("AddItem: Erwartet 1 Item, aber es sind %d", len(tl.Items))
	}
	if tl.Items[0].Description != "Test1" {
		t.Errorf("AddItem: Description stimmt nicht. Erwartet 'Test1', aber: %s", tl.Items[0].Description)
	}
}

func TestRemoveItem(t *testing.T) {
	tl := classes.TodoList{}
	tl.AddItem("A", "x")
	tl.AddItem("B", "x")

	tl.RemoveItem(1) // Entfernt erstes Item
	if len(tl.Items) != 1 {
		t.Errorf("RemoveItem: Erwartet 1 verbleibendes Item, aber es sind %d", len(tl.Items))
	}
	if tl.Items[0].Description != "B" {
		t.Errorf("RemoveItem: Falsches Item gelöscht? Erwartet 'B', aber: %s", tl.Items[0].Description)
	}
}

func TestToggleItem(t *testing.T) {
	tl := classes.TodoList{}
	tl.AddItem("ToggleMe", "xxx")

	// Index = 1 (User-Eingabe, 1-basiert)
	tl.ToggleItem(1)
	if len(tl.Items) != 0 {
		t.Errorf("ToggleItem: Erwartet, dass das Item sofort gelöscht wird, aber Items = %d", len(tl.Items))
	}
}

func TestInvalidIndex(t *testing.T) {
	tl := classes.TodoList{}

	// Sollte nicht abstürzen, wenn Liste leer ist
	tl.RemoveItem(1)
	tl.ToggleItem(1)

	if len(tl.Items) != 0 {
		t.Errorf("Erwartet 0, aber Items = %d", len(tl.Items))
	}
}
