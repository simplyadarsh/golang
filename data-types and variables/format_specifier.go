package main

import ("fmt")

func main() {
	var(
		name string = "Tom"
		city string = "Delhi"
		country string = "India"
		code int = 91
	)
	fmt.Printf("%s, is a capital of %s whose country code is %d \n",city,country,code)
 // string and decimal format specifier
     fmt.Printf("Nice to see you here,%v.",name ) 
}