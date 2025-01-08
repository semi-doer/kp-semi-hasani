# Todo-Liste in Lua und Go

Dieses Projekt demonstriert eine **Todo-Liste** in **zwei Sprachen**:

1. **Lua**  
   - Als reine **Konsolenanwendung** (Eingabe und Ausgabe über das Terminal).  
   - Optionale **Love2D**-Variante für eine grafische Oberfläche (ausgelagert).  
   - **Unit-Tests** mit [Busted](https://lunarmodules.github.io/busted/) sind enthalten.

2. **Go**  
   - Ebenfalls eine **Konsolen-Anwendung**, die dieselben Kernkonzepte (Add, Remove, Toggle) nutzt.  
   - Veranschaulicht Objektorientierung in Go mithilfe von Structs und Methoden. 
   - **Unit-Tests** sind enthalten.

Ziel ist es, die **OOP-Konzepte** beider Sprachen zu vergleichen, indem man zwei ähnliche Todo-Listen mit **ähnlicher Funktionalität** gegenüberstellt.

## 1) Lua-Todo-Liste

### Anforderungen

- **Lua** installiert (oder eine Lua-Umgebung wie [LuaBinaries](https://luabinaries.sourceforge.net/) bzw. [Lua for Windows](https://github.com/rjpcomputing/luaforwindows)).
- Für **Unit-Tests**: [Busted](https://lunarmodules.github.io/busted/) (via LuaRocks oder manuell).

### Setup

1. **In den Projektordner wechseln**  
   Öffne ein Terminal (PowerShell oder CMD) und navigiere in den Ordner `todolist-lua/`.  

    Dort sollten folgende Dateien liegen:
- `classes/todo_item_class.lua`
- `classes/todo_list_class.lua`
- `main_console.lua`

2. **Optional: UTF-8 in Windows-Terminal**  
Unter Windows kann es sein, dass Umlaute oder Sonderzeichen (z. B. aus Busted) falsch dargestellt werden.  
Deshalb sollte die Codepage temporär auf UTF-8 umsgestellt werden mit folgendem Befehl:
    ```
    chcp 65001
    ```
    Anschließend zeigt das Terminal meist mehr Unicode-Zeichen korrekt an.

3. **Das Programm starten (Konsolen-Menü)**  
Um die Todo-Liste im Terminal zu starten, wird folgender Befehl benötigt:
    ```
    lua main.lua
    ```

    Daraufhin erscheint ein menügesteuertes Programm, indem man Todos hinzufügen, entfernen oder abhaken (toggle) kann.  

4. **Unit Tests ausführen (Busted)**  
- Es muss sichergestellt werden, dass [Busted](https://olivinelabs.com/busted/) installiert ist und der Befehl `busted` im `PATH` Verzeichnis liegt.  
- Als nächstes muss in das `test` Verzeichnis navigiert werden.
- Dort muss dann folgender Befehl ausgeführt werden um die Tests zu starten:
    ```
    busted test/
    ```

## 2) Go Todo-Liste

### Anforderungen
- **Go** installiert (Version 1.18 oder höher)
- Eine gültige `go.mod`-Datei im Projektverzeichnis (Go-Modul)

### Setup

1. **In den Projektordner wechseln**  
   Öffne ein Terminal (PowerShell oder CMD) und navigiere in den Ordner `todolist-go/`.  

   Dort sollten folgende Dateien liegen:
   - `classes/todo_item.go`
   - `classes/todo_list.go`
   - `test/todolist_test.go`
   - `main.go`
   - `go.mod`

2. **Optional: UTF-8 in Windows-Terminal**  
   Unter Windows kann es sein, dass Umlaute oder Sonderzeichen (z. B. aus Busted) falsch dargestellt werden.  
Deshalb sollte die Codepage temporär auf UTF-8 umsgestellt werden mit folgendem Befehl:
    ```
    chcp 65001
    ```

    Anschließend zeigt das Terminal viele Unicode-Zeichen korrekt an.

3. **Das Programm starten (Konsolen-Menü)**  
    Um die Todo-Liste im Terminal zu starten, wird folgender Befehl benötigt:
    ```
    go run .
    ```

    Daraufhin erscheint ein menügesteuertes Programm, in dem man Todos hinzufügen, entfernen oder abhaken (toggle) kann.  

4. **Unit Tests ausführen**  
- Als nächstes muss in das `test` Verzeichnis navigiert werden.
- Dort muss dann folgender Befehl ausgeführt werden um die Tests zu starten:
  ```
  go test -v
  ```
- Dabei werden alle Funktionen mit dem Präfix `Test` (z. B. `TestAddItem`) ausgeführt.



## 3) Kernfunktionen beider Listen

Egal ob Lua oder Go, die **Haupt-Methoden** sind identisch:

- **`addItem(description, date)`**  
  Erzeugt ein neues Todo-Objekt.
- **`removeItem(index)`**  
  Löscht einen Eintrag an der gegebenen Position.
- **`toggleItem(index)`**  
  Wechselt den Status zwischen `completed = false` und `true`. 
- **`listItems()`** (nur in der Konsolen-Lösung)  
  Gibt die Todos formatiert aus.

---

## 4) Vergleich


  - **Lua**: Keine nativen Klassen, sondern **Tabellen** + Metatables.
  - **Go**: Keine Klassen/Vererbung, aber **Structs** + Methoden (Receiver).


---
