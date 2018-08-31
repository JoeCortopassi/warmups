class TodoList {
  constructor () {
    this._items = [];
  }

  addItem (title) {
    const id = this._items.length;
    this._items.push({
      id,
      title,
      isCompleted: false 
    });
  }

  toggleItem (id) {
    this._items[id].isCompleted = !this._items[id].isCompleted
  }

  getAllItems() {
      return this._items
  }

  getCompletedItems () {
    return this._items.filter(item => item.isCompleted );
  }

  getUnCompletedItems () {
    return this._items.filter(item => !item.isCompleted );
  }
}


const output = (list) => {
  console.log("ALL: ", list.getAllItems().map(i => i.title))
  console.log("COMPLETED: ", list.getCompletedItems().map(i => i.title))
  console.log("UNCOMPLETED: ", list.getUnCompletedItems().map(i => i.title))
  console.log("\n\n");
}


const Todos = new TodoList();
Todos.addItem("First chore!");
output(Todos);
Todos.addItem("Second thing?");
output(Todos);
Todos.toggleItem(0);
output(Todos);
