package main 

import "fmt"

func main(){
	const name = "saul Goodman"
	const openRate = 30.5

	msg :=	fmt.Sprintf("Hi %s, yout open rate is %.1f Percent", name, openRate) 

	fmt.Println(msg)
}
