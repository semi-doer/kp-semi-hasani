package.path = package.path .. ";../?.lua"
local TodoList = require("classes.todo_list_class")

describe("TodoList", function()
  
  it("sollte neue Items hinzufügen können", function()
    local tl = TodoList:new()
    tl:addItem("Test-Item 1", "01-01-2024")
    assert.are.equal(1, #tl.items)
    assert.are.equal("Test-Item 1", tl.items[1].description)
    assert.are.equal(false, tl.items[1].completed)
  end)

  it("sollte Items entfernen können", function()
    local tl = TodoList:new()
    tl:addItem("Item1", "01-01-2024")
    tl:addItem("Item2", "01-01-2024")
    
    tl:removeItem(1)  -- Erstes Element löschen
    assert.are.equal(1, #tl.items)
    assert.are.equal("Item2", tl.items[1].description)
  end)

  it("sollte den Toggle-Status setzen und das Item ggf. löschen", function()
    local tl = TodoList:new()
    tl:addItem("ToggleMe", "01-01-2024")
    
    -- Toggle führt von `false` -> `true` und löscht das Item direkt
    tl:toggleItem(1)
    assert.are.equal(0, #tl.items)
  end)

  it("sollte keinen Fehler werfen bei ungültigen Indizes", function()
    local tl = TodoList:new()
    -- Remove an leerer Liste
    tl:removeItem(1)
    -- Toggle an leerer Liste
    tl:toggleItem(1)
    assert.are.equal(0, #tl.items)
  end)

end)
