package main
import "fmt"
type Account interface {
	depositMoney(float32)
	checkBalance()
	withDrawMoney(float32)
	getInterest()
}

//savings account

type savingsAccount struct {
	accountNo int
	interestRate float32
	balance float32
}
type currentAccount struct {
	accountNo int
	balance float32
}

func (currAccount *currentAccount) checkBalance(){
	fmt.Println("Balance of current account is ",currAccount.balance)
}
func (currAccount *currentAccount) withDrawMoney(amount float32){
	if currAccount.balance-amount < 500 {
		fmt.Println("Minimum balance of 500 is required")
	}else{
		currAccount.balance = currAccount.balance - amount
		fmt.Println("Amount Withdraw ",amount," current Balance: ",currAccount.balance )
	}
}
func (savingsAcc *savingsAccount ) checkBalance(){
	fmt.Println("Balance of Savings Account is ",savingsAcc.balance)
}
func (savingsAcc *savingsAccount) withDrawMoney(amount float32){
	if savingsAcc.balance-amount < 500 {
		fmt.Println("Minimum balance of 500 is required")
	}else{
		savingsAcc.balance = savingsAcc.balance - amount
		fmt.Println("Amount Withdraw ",amount," current Balance: ",savingsAcc.balance )
	}
}
func (savingsAcc *savingsAccount) getInterest(){
	fmt.Println("Your Intrest Amount is ",savingsAcc.balance*savingsAcc.interestRate/100," on balance ",savingsAcc.balance)
}
func (savingsAcc *savingsAccount) depositMoney(amount float32){
	savingsAcc.balance = savingsAcc.balance + amount;
	fmt.Println("Total balance : ",savingsAcc.balance)
}


func main(){
	var acc1 Account = &savingsAccount{accountNo:101,interestRate: 2,balance: 2000};
	acc1.checkBalance();
	acc1.depositMoney(2000)
	acc1.getInterest()
	acc1.withDrawMoney(500)


	}