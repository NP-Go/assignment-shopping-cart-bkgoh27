package main

import (
	"fmt"
	"strconv"
	"strings"
)

type categoryList struct {
	name string
}

type shoppingList struct {
	category string
	quantity int
	cost     int
}

var m map[string]shoppingList
var c map[int]categoryList

func init() {
	m = make(map[string]shoppingList)
	m["Cup"] = shoppingList{"Household", 5, 3}
	m["Cake"] = shoppingList{"Food", 3, 1}
	m["Sprite"] = shoppingList{"Drinks", 5, 2}
	m["Fork"] = shoppingList{"Household", 4, 3}
	m["Bread"] = shoppingList{"Food", 2, 2}
	m["Plates"] = shoppingList{"Household", 4, 3}
	m["Coke"] = shoppingList{"Drinks", 5, 2}

	c = make(map[int]categoryList)
	c[0] = categoryList{"Household"}
	c[1] = categoryList{"Food"}
	c[2] = categoryList{"Drinks"}
}

func mainMenuChoice() {
	fmt.Println("Shopping List Application")
	fmt.Println(strings.Repeat("=", 25))
	fmt.Println("1. View entire shopping list.")
	fmt.Println("2. Generate Shopping List Report")
	fmt.Println("3. Add Items.")
	fmt.Println("4. Modify Items.")
	fmt.Println("5. Delete Item.")
	fmt.Println("6. Print Current Data.")
	fmt.Println("7. Add new Category Name.")
}

func generateReportChoice() {
	fmt.Println("Generate Report")
	fmt.Println("1. Total Cost of each category")
	fmt.Println("2. List of item by category")
	fmt.Println("3. Main Menu.")
}

func mainMenuView() int {
	mainMenuChoice()

	var choice int
	fmt.Println("Select your choice:")
	fmt.Scanln(&choice)
	return choice
}

func shoppingListView() {
	fmt.Println("Shopping List Contents:")
	getShoppingList()
}

func reportByCategoryTotalCostView() {
	fmt.Println("Total cost By Category.")
	getTotalCostByCategory()
}

func reportByCategoryListView() {
	// TODO: How to sort by category?
	fmt.Println("List by Category.")

	// fmt.Println("Category: Household - Item: Cups Quantity: 5 Unit Cost: 3")
	// fmt.Println("Category: Household - Item: Fork Quantity: 4 Unit Cost: 3")
	// fmt.Println("Category: Household - Item: Plates Quantity: 4 Unit Cost: 3")
	// fmt.Println("Category: Food - Item: Cake Quantity: 3 Unit Cost: 1")
	// fmt.Println("Category: Food - Item: Bread Quantity: 2 Unit Cost: 2")
	// fmt.Println("Category: Drinks - Item: Coke Quantity: 5 Unit Cost: 2")
	// fmt.Println("Category: Drinks - Item: Sprite Quantity: 5 Unit Cost: 2")
}

func getTotalCostByCategory() {
	type shoppingCategory struct {
		quantity int
		cost     int
	}

	var k map[string]shoppingCategory
	k = make(map[string]shoppingCategory)

	k["Household"] = shoppingCategory{0, 0}
	k["Food"] = shoppingCategory{0, 0}
	k["Drinks"] = shoppingCategory{0, 0}

	var totalQuantity, totalCost int
	for key := range m {
		totalQuantity = k[m[key].category].quantity + m[key].quantity
		totalCost = k[m[key].category].cost + m[key].quantity*m[key].cost
		k[m[key].category] = shoppingCategory{totalQuantity, totalCost}
	}

	fmt.Printf("Household cost: %v \n", k["Household"].cost)
	fmt.Printf("Food cost: %v \n", k["Food"].cost)
	fmt.Printf("Drink cost: %v \n", k["Drinks"].cost)
}

func generateReportView() {
	generateReportChoice()

	var choice int
	fmt.Println("\nChoose your report:")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		reportByCategoryTotalCostView()
	case 2:
		reportByCategoryListView()
	case 3:
		main()
	default:
		break
	}
}

func addItem() {
	var name, category string
	var quantity, cost int
	fmt.Println("What is the name of your item?")
	fmt.Scanln(&name)
	fmt.Println("What category does it belong to?")
	fmt.Scanln(&category)
	fmt.Println("How many units are there?")
	fmt.Scanln(&quantity)
	fmt.Println("How much does it cost per unit?")
	fmt.Scanln(&cost)

	addItemAction(name, category, quantity, cost)
}

func addItemAction(name, category string, quantity, cost int) {
	m[name] = shoppingList{category, quantity, cost}
}

func modifyItem() {
	fmt.Println("Modify Items.")

	var name, newName, category string
	var quantity, cost int
	fmt.Println("Which item would you wish to modify?")
	fmt.Scanln(&name)
	fmt.Printf("Current item name is %v - Category is %v - Quantity is %v - Unit Cost %v \n", name, m[name].category, m[name].quantity, m[name].cost)

	fmt.Println("Enter new name. Enter for no change.")
	fmt.Scanln(&newName)
	fmt.Println("Enter new Category. Enter for no change.")
	fmt.Scanln(&category)
	fmt.Println("Enter new Quantity. Enter for no change.")
	fmt.Scanln(&quantity)
	fmt.Println("Enter new Unit cost. Enter for no change.")
	fmt.Scanln(&cost)
	if category == "" {
		fmt.Println("No changes to category made.")
		category = m[name].category
	}
	if strconv.Itoa(quantity) == "" {
		fmt.Println("No changes to quantity made.")
		quantity = m[name].quantity
	}
	if strconv.Itoa(cost) == "" {
		fmt.Println("No changes to unit cost made.")
		cost = m[name].cost
	}
	if newName == "" {
		fmt.Println("No changes to item name made.")
	}

	modifyItemAction(name, newName, category, quantity, cost)
}

func modifyItemAction(name, newName, category string, quantity, cost int) {
	if newName == "" {
		m[name] = shoppingList{category, quantity, cost}
	} else {
		m[newName] = shoppingList{category, quantity, cost}
		deleteItemAction(name)
	}
}

func deleteItem() {
	fmt.Println("Delete Items.")

	var name string
	fmt.Println("Enter item name to delete:")
	fmt.Scanln(&name)

	deleteItemAction(name)
}

func deleteItemAction(name string) {
	if _, found := m[name]; found {
		delete(m, name)
		fmt.Println("Deleted", name)
	} else {
		fmt.Println("Item not found. Nothing to delete!")
	}
}

func printCurrentDataFields() {
	fmt.Println("Print Current Data.")

	if len(m) == 0 {
		fmt.Println("No data found!")
	} else {
		for key := range m {
			// TODO: What is the first element suppose to be?
			// It should be an assigned index...
			// fmt.Println("Bread - {1 1 2}")
			// fmt.Println("Chips - {3 11 14}")

			fmt.Println(key, "-", m[key])
		}
	}
}

func addNewCategoryName() {
	fmt.Println("Add New Category name.")

	var newCategory string
	fmt.Println("What is the New Category Name to add?")
	fmt.Scanln(&newCategory)

	addNewCategoryNameAction(newCategory)
}

func addNewCategoryNameAction(newCategory string) {
	if newCategory == "" {
		fmt.Println("No Input Found!")
		main()
	}

	flag := true
	for i := 0; i < len(c); i++ {
		if c[i].name == newCategory {
			fmt.Printf("Category %v already exist in index %v !", newCategory, i)
			flag = false
			break
		}
	}

	if flag == true {
		new_index := len(c) + 1
		c[new_index] = categoryList{newCategory}
		fmt.Printf("New category: %v added at index %v", newCategory, new_index)
	}
}

func getShoppingList() {
	for key := range m {
		fmt.Printf("Category: %v - Item: %v Quantity: %v Unit Cost: %v \n", m[key].category, key, m[key].quantity, m[key].cost)
	}
}

func main() {
	choice := mainMenuView()
	fmt.Println(choice)
	switch choice {
	case 1:
		shoppingListView()
	case 2:
		generateReportView()
	case 3:
		addItem()
		main()
	case 4:
		modifyItem()
		main()
	case 5:
		deleteItem()
	case 6:
		printCurrentDataFields()
	case 7:
		addNewCategoryName()
	default:
		break
	}
}
