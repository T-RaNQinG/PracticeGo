package main

type Account struct {
	FirstName string
	LastName string	
}

func (account *Account) ChangeName(firstName ,lastName string){

	account.FirstName = firstName
	account.LastName = lastName
}

type Employee struct {
	store float64
	account Account
}

func (employee *Employee) AddCredits(money float64){
	employee.store += money
}

func (employee *Employee) RemoveCredits(){
	employee.store = 0
}

func (employee *Employee) CheckoutCredits() float64 {
	return employee.store
}



