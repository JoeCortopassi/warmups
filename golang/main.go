package main

import "fmt"

func main() {
	myItems := Items{}
	myItems.addItem("first thing")
	log(myItems)
	myItems.addItem("second thing")
	log(myItems)
	myItems.addItem("third thing")
	myItems.toggleItem(1)
	log(myItems)
}

func log(list Items) {
	fmt.Println("ALL: ", list.getAllItems())
	fmt.Println("UNCOMPLETED: ", list.getUnCompletedItems())
	fmt.Println("COMPLETED:", list.getCompletedItems())
}

type Item struct {
	id          int
	title       string
	isCompleted bool
}

type Items []Item

type TodoList interface {
	addItem(title string)
	toggleItem(id int)
	getAllItems() []Item
	getCompletedItems() []Item
	getUnCompletedItems() []Item
}

func (i *Items) addItem(title string) {
	item := Item{len(*i), title, false}
	*i = append(*i, item)
}

func (i Items) toggleItem(id int) {
	item := i[id]
	item.isCompleted = !item.isCompleted
	i[id] = item
}

func (i Items) getAllItems() Items {
	return i
}

func (i Items) getCompletedItems() Items {
	filteredItems := Items{}
	for _, val := range i {
		if val.isCompleted {
			filteredItems = append(filteredItems, val)
		}
	}

	return filteredItems
}

func (i Items) getUnCompletedItems() Items {
	filteredItems := Items{}
	for _, val := range i {
		if !val.isCompleted {
			filteredItems = append(filteredItems, val)
		}
	}

	return filteredItems
}
