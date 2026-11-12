package main

import "fmt"

type Product struct {
	Name string
	Price float64
	Quantity int
}

func showWelcome(){
	fmt.Println("INVENTORY MANAGEMENT SYSTEM")
}
func showVersion(){
	fmt.Println("Version: 1.0.0")
}
func showMenu(){
	fmt.Printf("\n1. Add Product\n2. View Products\n3. Update Stock\n4. Exit\n")
}
func calculateStockValue(product Product) float64 {
	return product.Price * float64(product.Quantity)
}
func canRemoveStock(product Product, qt int) bool {
	if qt <= 0 ||  qt > product.Quantity{
		return false
	}
	return true
}
func validateStockRemoval(product Product, requestedQuantity int)(bool, string){
	if requestedQuantity <= 0{
		return false,  "Quantity must be greater than zero"
	}
	if requestedQuantity > product.Quantity {
		return false, "Insufficient stock"

	}
	return true, "Stock available"
}

func checkPurchase(product Product, requestedQuantity int) (bool,float64, string){
	setPrice := true
	var totalPrice float64
	if !setPrice{
		totalPrice = 0.0
	}else{
		totalPrice = float64(requestedQuantity) * product.Price
	}
	exist,value := validateStockRemoval(product, requestedQuantity)
	if !exist {
		setPrice = false
		return false, totalPrice, value
	}

	if !exist {
		setPrice = false
		return false, totalPrice, "Insufficient stock"
	}

	return true, totalPrice, "urchase allowed"

}

func removeProduct(products []Product, name string)([]Product, bool){

	for i, product := range products{
		if product.Name == name {
		products = append(products[:i], products[i+1:]...)
		return products, true

		}
		
	}
	return nil, false
}

func main(){
	showWelcome()
	showVersion()
	showMenu()

// 	products := []Product{
//     {"Laptop", 1200, 5},
//     {"Mouse", 25, 10},
//     {"Keyboard", 75, 8},
//     {"Monitor", 300, 3},
// }


product := Product{
    Name:     "Laptop",
    Price:    1200,
    Quantity: 5,
}
requestQuantity := 5
exist, price, value := checkPurchase(product, requestQuantity)
if exist{
	fmt.Printf("Purchase allowed: %t\n", exist)
}else{
	fmt.Printf("Purchase not allowed: %t\n", exist)

}	
fmt.Printf("Total Price: £%.2f\nMessage: %s\n", price, value)



}