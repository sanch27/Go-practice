package main // entry point of app 
import "fmt" // importing the formatted output

func main () { // entry point of the project
	fmt.Println(" Hello World ")  // ln is used to print the data added in it directly

	//Assigment operator (variable should be declared to the value)
	var name string = "Sanchali"
 	var age int = 22
 	fmt.Printf("Name: %s\nAge: %d\n", name, age) // f is for formatted output 

	//Short declaration operator ( variable declation is not needed )
	batsman := "David Warner"
	runs := 34
	fmt.Printf("Batsmen: %s\nRuns: %d\n", batsman, runs)

}