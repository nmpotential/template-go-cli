package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)
//3. Declare variables
var originUnit string   // Input: The arguments(s) that will be parsed into the CLI
var originValue float64 // Input: The keyboard input that will be read from the terminal

var shouldConvertAgain string // Input: The input functionality that enables the app to either restart or end

var err error //Output: The error responsible for handling all other potential errors the app may encounter

var errInvalidArguments = errors.New("Invalid arguments") // Output: The error responsible for handling invalid arguments(s) that will be parsed into the CLI
var errReadingInput = errors.New("Error reading input")   // Output: The error responsible for handling issues with keyboard input that will be read from the terminal

func main() {
	//4. Validate Argument(s)
	// We use the built-in len function to check if the length of os.Args is different than 2.
	if len(os.Args) != 2 {
		// If argument invalid we invoke the printError() function, passing errInvalidArguments as the argument.
		printError(errInvalidArguments)
	}
	//5. Read origin unit
	// Invoke strings.ToUpper function to ensure consistency when reading command line
	// originUnit value provided by the user
	// originUnit is the 2nd argument, os.Args[1], parsed into CLI
	// the executable is 1st argument , os.Args[0], by default
	originUnit := strings.ToUpper(os.Args[1])
	//6. Read current temperature (originValue) Keyboard Input
	for {
		// Print statement that will prompt user to enter current temperature
		fmt.Print("What is the current temperature in " + originUnit + " ? ")
		// Invoke the fmt.Scanln() function, passing &originValue as the argument.
		// Assign the two return values to the variables _ and err respectively.
		_, err := fmt.Scanln(&originValue)
 		// If statement checking if err != nil, and if that's true,
 		// invoke the printError() function, passing errReadingInput as its argument.
		if err != nil {
			printError(errReadingInput)
		}
		//9. Generate methods that will utilise core functionality functions

		// If statement to check if originUnit is equal to "C" (notice the capital casing).
		// If so, invoke the convertToFahrenheit() function passing originValue as the argument.
		if originUnit == "C" {
			// Convert current temperature from Celsius to Fahrenheit
			convertToFahrenheit(originValue)
			// Otherwise, create an else block and invoke the convertToCelsius() function,
			// passing originValue as the argument.
		} else {
			// Convert current temperature from Fahrenheit to Celsius
			convertToCelsius(originValue)
		}
		fmt.Print("Would you like to convert another temperature ? (y/n) ")
		//10. Prompt to convert again
		// Invoke the fmt.Scanln() function passing &shouldConvertAgain as its argument.
		// Assign the two return values to the previously defined variables _ and err respectively.
		_, err = fmt.Scanln(&shouldConvertAgain)
		// On the following line, create an if statement checking if err != nil.
		// If that condition is true, invoke printError() passing errReadingInput as its argument.
		if err != nil {
			printError(errReadingInput)
		}
		//11. Parse prompt answer
		// To ensure consistency when reading command line, we invoke a combination
		// of strings.ToUpper function passing it the result of calling the
		// strings.TrimSpace() function with shouldConvertAgain as its argument.
		// If result is NOT equal to "Y", then the existing if block containing then "Good bye!" followed by break
		if strings.ToUpper(strings.TrimSpace(shouldConvertAgain)) != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}
//7. Create function that handles runtime errors
func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

// 8. State each core functionality of app and create functions for each
// Convert current temperature from Fahrenheit to Celsius
func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}
// Convert current temperature from Celsius to Fahrenheit
func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}