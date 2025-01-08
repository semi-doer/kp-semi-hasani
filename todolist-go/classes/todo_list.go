package classes

import "fmt"

// TodoList verwaltet eine Sammlung von TodoItems.
type TodoList struct {
	Items []TodoItem
}

// AddItem erzeugt einen neuen Eintrag (TodoItem) und hängt ihn an.
func (tl *TodoList) AddItem(description, date string) {
	item := TodoItem{
		Description: description,
		Date:        date,
		Completed:   false,
	}
	tl.Items = append(tl.Items, item)
}

// RemoveItem entfernt den Eintrag an Index i (1-basiert oder 0-basiert je nach Konvention).
// Hier nehmen wir an, i ist 1-basiert, so wie der User tippen würde.
func (tl *TodoList) RemoveItem(i int) {

	i-- // => Damit "1" auf Index 0 abgebildet wird

	if i >= 0 && i < len(tl.Items) {
		tl.Items = append(tl.Items[:i], tl.Items[i+1:]...)
	} else {
		fmt.Println("Ungültiger Index!")
	}
}

// ToggleItem schaltet den Status eines Eintrags (von offen zu erledigt oder umgekehrt).
func (tl *TodoList) ToggleItem(i int) {
	i--
	if i >= 0 && i < len(tl.Items) {
		wasCompleted := tl.Items[i].Completed
		tl.Items[i].Toggle()

		if !wasCompleted && tl.Items[i].Completed {
			// Sobald wir von false -> true wechseln, entfernen wir es
			tl.Items = append(tl.Items[:i], tl.Items[i+1:]...)
		}
	} else {
		fmt.Println("Ungültiger Index!")
	}
}

// ListItems gibt alle Einträge mit Index, Status, Beschreibung und Datum aus.
func (tl *TodoList) ListItems() {
	if len(tl.Items) == 0 {
		fmt.Println("Keine Todos vorhanden.")
		return
	}
	for idx, item := range tl.Items {
		checkbox := "[ ]"
		if item.Completed {
			checkbox = "[X]"
		}
		// Index +1, weil wir dem User "1" als erstes Element zeigen
		fmt.Printf("%d) %s %s (%s)\n", idx+1, checkbox, item.Description, item.Date)
	}
}
