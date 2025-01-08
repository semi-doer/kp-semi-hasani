local TodoItem = {}
TodoItem.__index = TodoItem

--- Konstruktor für ein einzelnes Todo-Item.
-- @param description Kurze Beschreibung des Tasks
-- @param date        Fälligkeit oder beliebiges Datumsfeld
function TodoItem:new(description, date)
    local instance = {
        description = description or "",
        date        = date or "",
        completed   = false
    }
    setmetatable(instance, TodoItem)
    return instance
end

--- Wechselt den `completed`-Status dieses Items zwischen true/false.
function TodoItem:toggleCompleted()
    self.completed = not self.completed
end

return TodoItem
