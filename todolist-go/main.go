package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/TODOLIST-GO-OOP/classes"
)

// parseIndex wandelt eine String-Eingabe (z. B. "2") in int um (0 bei Fehler).
func parseIndex(s string) int {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	if err != nil {
		fmt.Println("Ungültige Zahl!")
		return 0
	}
	return i
}

func main() {
	var todoList classes.TodoList

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Menü ---")
		fmt.Println("[1] Neues Todo hinzufügen")
		fmt.Println("[2] Todo entfernen")
		fmt.Println("[3] Todo togglen (abhaken/öffnen)")
		fmt.Println("[4] Liste anzeigen")
		fmt.Println("[5] Beenden")

		fmt.Print("Wähle eine Option: ")
		if !scanner.Scan() {
			break // EOF oder Fehler
		}
		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			// Neues Todo hinzufügen
			fmt.Print("Beschreibung: ")
			if !scanner.Scan() {
				continue
			}
			desc := strings.TrimSpace(scanner.Text())

			fmt.Print("Datum (z.B. 10-01-2024): ")
			if !scanner.Scan() {
				continue
			}
			date := strings.TrimSpace(scanner.Text())

			todoList.AddItem(desc, date)
			fmt.Println("Todo hinzugefügt!")

		case "2":
			// Todo entfernen
			todoList.ListItems()
			fmt.Print("Welches Todo soll entfernt werden? (Zahl): ")
			if !scanner.Scan() {
				continue
			}
			idxStr := strings.TrimSpace(scanner.Text())
			idx := parseIndex(idxStr)
			if idx > 0 {
				todoList.RemoveItem(idx)
			}

		case "3":
			// Todo togglen
			todoList.ListItems()
			fmt.Print("Welches Todo soll getoggelt werden? (Zahl): ")
			if !scanner.Scan() {
				continue
			}
			idxStr := strings.TrimSpace(scanner.Text())
			idx := parseIndex(idxStr)
			if idx > 0 {
				todoList.ToggleItem(idx)
			}

		case "4":
			// Liste anzeigen
			todoList.ListItems()

		case "5", "exit", "q":
			fmt.Println("Auf Wiedersehen!")
			return

		default:
			fmt.Println("Ungültige Auswahl. Bitte 1-5 eingeben.")
		}
	}
}
