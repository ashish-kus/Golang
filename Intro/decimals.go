package main 

import "fmt"

var smsSendingLimit int
var costPerSMS float64
var hasPermission bool
var username string


func main(){
	fmt.Printf("%v %f %v %q\n",
	smsSendingLimit, 
	costPerSMS, 
	hasPermission, 
	username)
}
