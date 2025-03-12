package main // entry point of app 
import "fmt" // importing the formatted output

func main () { // entry point of the project
	fmt.Println(" Hello World ")  // ln is used to print the data added in it directly

	//Variables
	//Assigment operator (variable should be declared to the value)
	var name string = "Sanchali"
 	var age int = 22
 	fmt.Printf("Name: %s\nAge: %d\n", name, age) // f is for formatted output 

	//Short declaration operator ( variable declation is not needed )
	batsman := "David Warner"
	runs := 134
	fmt.Printf("Batsmen: %s\nRuns: %d\n", batsman, runs)

	//if else (condtional statements)
	if runs >= 50 {
		fmt.Println("Half-century!")
	} else { fmt.Println("Not a half-century!") }	

	//for (loop statement)
	//For is the only loop in go it hasscope only within the loop
	for i := 0; i < 5; i++ { 
		fmt.Println(i)
	}	

	//using for as while loop ( Go does not have while loop or dowhile loop)
	count := 0
    
    // while (count < 5)  in other languages
    for count < 5 {
        fmt.Println("Count:", count)
        count++
}
}