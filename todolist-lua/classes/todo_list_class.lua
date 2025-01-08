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
-- @param date        Beliebiges Datumsformat
function TodoList:addItem(description, date)
    local item = TodoItem:new(description, date)
    table.insert(self.items, item)
end

--- Entfernt den Eintrag an Position `index` aus der Liste.
function TodoList:removeItem(index)
    if index >= 1 and index <= #self.items then
        table.remove(self.items, index)
    end
end

--- Wechselt den Status eines Items an Position `index` und entfernt es ggf.
--  Wenn ein offenes Item (completed=false) abgehakt wird (completed=true),
--  wird es direkt aus der Liste gelöscht.
function TodoList:toggleItem(index)
    if index >= 1 and index <= #self.items then
        local item = self.items[index]
        local wasCompleted = item.completed

        item:toggleCompleted()  -- schaltet von false->true oder true->false

        -- Entfernungs-Logik: Sobald ein Item von "offen" auf "erledigt" wechselt.
        if (not wasCompleted) and item.completed then
            table.remove(self.items, index)
        end
    end
end

return TodoList
