local TodoList = require("classes.todo_list_class")

local myTodoList = TodoList:new()

-- Hilfsfunktionen:
local function parseNumber(str)
    local num = tonumber(str)
    if not num then
        return 0
    end
    return num
end

local function listItems()
    if #myTodoList.items == 0 then
        print("Keine Todos vorhanden.")
        return
    end
    for i, item in ipairs(myTodoList.items) do
        local checkbox = item.completed and "[X]" or "[ ]"
        print(string.format("%d) %s %s (%s)", i, checkbox, item.description, item.date))
    end
end

local function addItem()
    io.write("Beschreibung eingeben: ")
    local desc = io.read("*l")
    if not desc or desc == "" then
        print("Fehler: Beschreibung darf nicht leer sein!\n")
        return
    end
    io.write("Datum eingeben (z.B. 2024-01-10): ")
    local dat = io.read("*l")
    myTodoList:addItem(desc, dat or "")
    print("Todo hinzugefuegt!\n")
end

local function removeItem()
    listItems()
    io.write("Welches Item soll entfernt werden? (Zahl): ")
    local idx = parseNumber(io.read("*l"))
    if idx < 1 or idx > #myTodoList.items then
        print("Ungueltiger Index!\n")
        return
    end
    myTodoList:removeItem(idx)
    print("Item entfernt.\n")
end

local function toggleItem()
    listItems()
    io.write("Welches Item soll getoggelt werden? (Zahl): ")
    local idx = parseNumber(io.read("*l"))
    if idx < 1 or idx > #myTodoList.items then
        print("Ungueltiger Index!\n")
        return
    end
    myTodoList:toggleItem(idx)
    print("Toggle ausgefuehrt.\n")
end

local function main()
    -- Füge ein paar Test-Einträge hinzu
    myTodoList:addItem("Praesentation vorbereiten", "2024-01-10")
    myTodoList:addItem("Kapitel lesen", "2024-01-15")

    while true do
        print("Menue:")
        print("[1] Neues Todo hinzufuegen")
        print("[2] Todo entfernen")
        print("[3] Todo abhaken (toggle)")
        print("[4] Liste anzeigen")
        print("[5] Beenden")
        io.write("Waehle eine Option: ")
        local choice = io.read("*l")

        if choice == "1" then
            addItem()
        elseif choice == "2" then
            removeItem()
        elseif choice == "3" then
            toggleItem()
        elseif choice == "4" then
            listItems()
        elseif choice == "5" or choice == "q" then
            print("Auf Wiedersehen!")
            break
        else
            print("Ungueltige Auswahl, bitte 1-5 eingeben.\n")
        end
    end
end

main()
