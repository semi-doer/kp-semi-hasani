package classes

import (
	"fmt"
	"regexp"
	"sync"
)

// TodoList verwaltet eine Sammlung von TodoItems.
type TodoList struct {
	mu    sync.Mutex // Mutex für threadsicheren Zugriff
	Items []TodoItem
}

// AddItem fügt ein neues TodoItem zur Liste hinzu.
func (tl *TodoList) AddItem(description, date string) {
	// Datum überprüfen: Format "DD-MM-YYYY"
	match, _ := regexp.MatchString(`^\d{2}-\d{2}-\d{4}$`, date)
	if !match {
		fmt.Println("Ungültiges Datum: Das Datum muss im Format 'DD-MM-YYYY' sein.")
		return
	}

	// Element hinzufügen, wenn das Datum gültig ist
	tl.mu.Lock()
	defer tl.mu.Unlock()

	item := TodoItem{
		Description: description,
		Date:        date,
		Completed:   false,
	}
	tl.Items = append(tl.Items, item)
}

// RemoveItem entfernt ein TodoItem anhand des Indexes (1-basiert).
func (tl *TodoList) RemoveItem(index interface{}) {
	tl.mu.Lock()
	defer tl.mu.Unlock()

	// Überprüfen, ob der Index eine Zahl ist
	idx, ok := index.(int)
	if !ok {
		fmt.Println("Ungültiger Index: Der Index muss eine Zahl sein.")
		return
	}

	// Prüfen, ob der Index im gültigen Bereich liegt
	if idx < 1 || idx > len(tl.Items) {
		fmt.Printf("Ungültiger Index: Der Index liegt außerhalb des gültigen Bereichs (1 bis %d).\n", len(tl.Items))
		return
	}

	// Element entfernen (Index 1-basiert zu 0-basiert umwandeln)
	tl.Items = append(tl.Items[:idx-1], tl.Items[idx:]...)
	fmt.Println("Element erfolgreich entfernt.")
}

// ToggleItem schaltet den Status eines Eintrags um und entfernt ihn, falls er als erledigt markiert wird.
func (tl *TodoList) ToggleItem(i interface{}) {
	tl.mu.Lock()
	defer tl.mu.Unlock()

	// Überprüfen, ob der übergebene Index eine Zahl ist
	index, ok := i.(int)
	if !ok {
		fmt.Println("Ungültiger Index: Der Index muss eine Zahl sein.")
		return
	}

	// Prüfen, ob die Liste leer ist
	if len(tl.Items) == 0 {
		fmt.Println("Die Liste ist leer.")
		return
	}

	// Index von 1-basiert auf 0-basiert umwandeln
	index--

	// Prüfen, ob der Index im gültigen Bereich liegt
	if index < 0 || index >= len(tl.Items) {
		fmt.Printf("Ungültiger Index: Der Index liegt außerhalb des gültigen Bereichs (1 bis %d).\n", len(tl.Items))
		return
	}

	// Element holen und toggeln
	item := &tl.Items[index]
	item.Toggle()

	// Entfernen, wenn es erledigt wurde
	if item.Completed {
		tl.Items = append(tl.Items[:index], tl.Items[index+1:]...)
		fmt.Println("Das Element wurde erledigt und aus der Liste entfernt.")
	} else {
		fmt.Println("Das Element wurde erfolgreich aktualisiert.")
	}
}

// ListItems gibt alle Einträge mit Index, Status, Beschreibung und Datum aus.
func (tl *TodoList) ListItems() {
	tl.mu.Lock()
	defer tl.mu.Unlock()

	if len(tl.Items) == 0 {
		fmt.Println("Keine Todos vorhanden.")
		return
	}

	for idx, item := range tl.Items {
		checkbox := "[ ]"
		if item.Completed {
			checkbox = "[X]"
		}
		fmt.Printf("%d) %s %s (%s)\n", idx+1, checkbox, item.Description, item.Date)
	}
}
