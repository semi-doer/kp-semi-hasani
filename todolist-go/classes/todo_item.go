package classes

// TodoItem repräsentiert einen einzelnen Eintrag in der Todo-Liste.
type TodoItem struct {
	Description string
	Date        string
	Completed   bool
}

// Toggle wechselt den Status Completed (erledigt / nicht erledigt).
func (t *TodoItem) Toggle() {
	t.Completed = !t.Completed
}
