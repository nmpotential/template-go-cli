# go-cli-template
Template that guides you through building a command line application which has the functionality to parse  CLI arguments 
and read keyboard inputs from the terminal in Go.

## Steb by Step Guide to: Building a CLI app that parses CLI arguments and reads keyboard inputs from the terminal in Go

- [x] Step 1. Briefly summaries how the apps supposed to work and put in **summary** of README.md - This will be updated as 
      functionality and implementation are refined:

- [ ] Step 2. Setup Environment
   
  - [ ] Go mod???
  - [x] Create main.go file
  - [x] Create **main_test.go** file
  - [x] Compile executable with `go build -o temp` - add command for 
  compiling to README under Usage.
  - [ ] Docker???

      
- [x] Step 3. Declare variables for the following (refer to summary) in main.go file:
      
  **Input**
  - [x] The arguments(s) that will be parsed into the CLI
  - [x] The keyboard input that will be read from the terminal
  - [x] The input functionality that enables the app to either restart or end 
  
  **Output**
  - [x] The error responsible for handling invalid arguments(s) that will 
  be parsed into the CLI - can simply be a message
  - [x] The error responsible for handling issues with keyboard input that 
  will be read from the terminal - can simply be a message
  - [x] The error responsible for handling all other potential errors the 
  app may encounter - can simply be a message
      
      
      ```
      package main

      import (
            "errors"
      )

      var originUnit string   // Input: The arguments(s) that will be parsed into the CLI
      var originValue float64 // Input: The keyboard input that will be read from the terminal

      var shouldConvertAgain string // Input: The input functionality that enables the app to either restart or end 

      var err error //Output: The error responsible for handling all other potential errors the app may encounter

      var errInvalidArguments = errors.New("Invalid arguments") // Output: The error responsible for handling invalid arguments(s) that will be parsed into the CLI
      var errReadingInput = errors.New("Error reading input")   // Output: The error responsible for handling issues with keyboard input that will be read from the terminal

      ```
      
- [x] Step 4. Validate Argument(s) in the ```func main()``` definition:

  - [x] Create a function to check if arguments are valid
  - [x] If argument invalid, invoke the print argument validation statement using 
  printError() function, passing **the error variable responsible for handling**
  **invalid arguments(s) that will be parsed into the CLI**
      
 ```
      func main() {
	// We use the built-in len function to check if the length of os.Args is different than 2.
	if len(os.Args) != 2 {
		// If argument is invalid we invoke the printError() function, passing errInvalidArguments as the argument.
		printError(errInvalidArguments)
	}
      
```

- [x] Step 5. Read Argument, below argument validation statement:

  - [x] Invoke suitable functions to ensure consistency when reading command line 
  arguments provided by the user, ```os.Args[1]``` passing as the argument - Note:
  index 0, ```os.Args[0]``` is assigned to the executable by default 
  - [x] Assign the result to the previously defined variable that will 
  be parsed as an argument into the CLI.
  
 ```
	//5. Read origin unit
	// Invoke strings.ToUpper function to ensure consistency when reading command line
	// originUnit value provided by the user
	// originUnit is the 2nd argument, os.Args[1], parsed into CLI 
	// the executable is 1st argument , os.Args[0], by default
	originUnit := strings.ToUpper(os.Args[1])
```

- [x] Step 6.  Read Keyboard Input

  - [x] Print statement that will prompt user to enter desired input that
  will be read from the terminal
  - [x] Invoke the fmt.Scanln() function, passing the address to the keyboard 
  input variable that as the argument.
  - [x] Assign the two return values to the variables _ and err respectively.
  - [x] Next, create an if statement checking if err != nil, and if that's true, 
  invoke the printError() function, passing the error responsible for handling 
  issues with keyboard input that will be read from the terminal as its argument.
  
```
	// Read current temperature (originValue) Keyboard Input
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
```
   
- [x] Step 7. Create function that handles runtime errors - Outside of ```func main()```


	```
		func printError(err error) {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	```

     

- [x] Step 8.  Create functions for each core functionality of app  - Outside of ```func main()```
	
  - [x] **Convert ***current temperature*** from Fahrenheit to Celsius** and print
	out **converted temperature in Celsius**.
	
	```
		func convertToCelsius(value float64) {
		convertedValue := (value - 32) * 5 / 9
		fmt.Printf("%v F = %.0f C\n", value, convertedValue)
	}
	```
	
  - [x] **Convert ***current temperature*** from Celsius to Fahrenheit** and print
	out **converted temperature in Fahrenheit**.
	```
		func convertToFahrenheit(value float64) {
		convertedValue := (value * 9 / 5) + 32
		fmt.Printf("%v C = %.0f F\n", value, convertedValue)
	}
	```
	
- [x] Step 9. Generate methods that will utilise core functionality functions
        
```
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
```



- [ ] Step 10. Prompt Restarting App


- [ ] Step 11. Parse Prompt Answer


      
      
## Summary

  - User compiles and runs executable from command-line parsing **the unit of temperature we want to convert from** 
  - User is then prompted to enter **current temperature**
  - Then the app reads **keyboard input** (**current temperature**) from terminal
  - Core functionality occurs that:
    - **Converts current temperature from Fahrenheit to Celsius** and prints out **converted temperature in Celsius**. 
    - **Converts current temperature from Celsius to Fahrenheit** and prints out **converted temperature in Fahrenheit**.
  - Finally, the app prompts user to enter **Y/N** to either restart or the end app

## Usage

Compile with `go build -o [executable_name]`

Then, invoke the binary passing as argument [**valid_argument**].
For example:

`./[executable_name] [valid_argument]` to...[Brief decription of what it does]

To verify your work locally for this module, run the following command in a terminal window: 

`go test -v`
