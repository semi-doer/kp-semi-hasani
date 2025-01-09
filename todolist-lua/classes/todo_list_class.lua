local TodoItem = require("classes.todo_item_class")

local TodoList = {}
TodoList.__index = TodoList

--- Konstruktor für die TodoList.
function TodoList:new()
    local instance = { items = {} }
    setmetatable(instance, TodoList)
    return instance
end

--- Fügt ein neues Item an das Ende der Liste an.
-- @param description Kurze Beschreibung
-- @param date        Datum in Format "DD-MM-YYYY"
function TodoList:addItem(description, date)
    if type(date) ~= "string" or not date:match("^%d%d%-%d%d%-%d%d%d%d$") then
        print("Ungültiges Datum: Das Datum muss im Format 'DD-MM-YYYY' angegeben werden.")
        return
    end

    -- Wenn das Format korrekt ist, das Item hinzufügen
    local item = TodoItem:new(description, date)
    table.insert(self.items, item)
end


function TodoList:removeItem(index)
    if type(index) ~= "number" then
        print("Ungültiger Index: Der Index muss eine Zahl sein.")
        return
    elseif index < 1 or index > #self.items then
        print("Ungültiger Index: Der Index liegt außerhalb des gültigen Bereichs (1 bis " .. #self.items .. ").")
        return
    end
    table.remove(self.items, index)
end

--- Wechselt den Status eines Items an Position `index` und entfernt es ggf.
--  Wenn ein offenes Item (completed=false) abgehakt wird (completed=true),
--  wird es direkt aus der Liste gelöscht.
function TodoList:toggleItem(index)
    if type(index) ~= "number" then
        print("Ungültiger Index: Der Index muss eine Zahl sein.")
        return
    elseif index < 1 or index > #self.items then
        print("Ungültiger Index: Der Index liegt außerhalb des gültigen Bereichs (1 bis " .. #self.items .. ").")
        return
    end

    local item = self.items[index]
    if not item or type(item.toggleCompleted) ~= "function" then
        print("Ungültiges Element an Index " .. index .. ": Erwartet wurde ein Objekt mit einer `toggleCompleted`-Methode.")
        return
    end

    local wasCompleted = item.completed

    item:toggleCompleted()  -- Schaltet von false->true oder true->false

    -- Entfernt das Item, wenn es von "offen" zu "erledigt" wechselt
    if (not wasCompleted) and item.completed then
        table.remove(self.items, index)
    end
end

return TodoList
