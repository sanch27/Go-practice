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

// using for for do while loop
counts := 0

for {  // Infinite loop (runs at least once)
	fmt.Println("Count:", counts)
	counts++

	// Condition check after execution
	if counts >= 5 {
		break
	}
}

//stuct 

//Predefined data should be given 
type Person struct { // Structs are used to define custom data types.
    Name string
    Age  int
}

p := Person{"John Doe", 25} // Persn is object
fmt.Printf("Name: %s\nAge: %d\n", p.Name, p.Age) // Call by variable and attribute 


//calling a function
fmt.Println(add(5, 6, 7))
}

func add(a int, b int, c int) int {
	return a+b-c

}


