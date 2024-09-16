package main

import ("fmt")

var a string = "Country is India"  // Global Varibale

func main() {
    var b string = "First city is Delhi " // Local Varibale
	fmt.Println(a)
   fmt.Println(b)
    {
        var c string = "Second city is Mumbai" // Local Variable
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
    }
}