# go-cli-template
Template that guides you through building a command line application which has the functionality to parse  CLI arguments 
and read keyboard inputs from the terminal in Go.

## Steb by Step Guide to: Building a CLI app that parses CLI arguments and reads keyboard inputs from the terminal in Go

1. - [ ] Briefly summaries how the apps supposed to work and put in **summary** of README.md - This will be updated as 
      functionality and implementation are refined:
      
      - User compiles and runs executable from command-line parsing [**valid argument**]
      - User is then prompted to enter [**input value**]
      - Then the app reads **keyboard input** ([**input value**]) from terminal and prints out **desired outcome**
      - Finally, the app prompts user to enter [**input value**] to either restart or end the app
            
      **Example:**
      - User compiles and runs executable from command-line parsing **the unit of temperature we want to convert from** 
      - User is then prompted to enter **current temperature**
      - Then the app reads **keyboard input** (**current temperature**) from terminal,
      converts **current temperature** and prints out **converted temperature between Fahrenheit and Celsius**.
      - Finally, the app prompts user to enter **Y/N** to either restart or the end app

2. - [ ] Setup Environment
      
      - [ ] Create main.go file
      - [ ] Create [**test**] file
      - [ ] Compile executable with `go build -o [executable_name]` command using assigning relevant executable name - add command 
      for compiling to README under Usage.
      
      **Example**:
      - [ ] Create main.go file
      - [ ] Create **main_test.go** file
      - [ ] Compile executable with `go build -o temp` - add command for 
      compiling to README under Usage.
      
3. - [ ] Declare variables for the following (refer to summary) in main.go file:
      
      **Input**
      - [ ] The arguments(s) that will be parsed into the CLI
      - [ ] The keyboard input that will be read from the terminal
      - [ ] The input functionality that enables the app to either restart or end 
      
      **Output**
      - [ ] The error responsible for handling invalid arguments(s) that will 
      be parsed into the CLI - can simply be a message
      - [ ] The error responsible for handling issues with keyboard input that 
      will be read from the terminal - can simply be a message
      - [ ] The error responsible for handling all other potential errors the 
      app may encounter - can simply be a message

      **Example**:
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
      
4. - [ ] Validate Argument(s) in the ```func main()``` definition:

      - [ ] Create a function to check if arguments are valid
      - [ ] If argument invalid, invoke the print argument validation statement using 
      printError() function, passing **the error variable responsible for handling**
      **invalid arguments(s) that will be parsed into the CLI**
      
      **Example:**
 ```
      func main() {
	// We use the built-in len function to check if the length of os.Args is different than 2.
	if len(os.Args) != 2 {
		// If argument is invalid we invoke the printError() function, passing errInvalidArguments as the argument.
		printError(errInvalidArguments)
	}
      
```

      
      
      

5. - [ ] Read Argument, below argument validation statement:

      - [ ] Invoke suitable functions to ensure consistency when reading command line 
      arguments provided by the user, ```os.Args[1]``` passing as the argument - Note:
      index 0, ```os.Args[0]``` is assigned to the executable by default 
      - [ ] Assign the result to the previously defined variable that will 
      be parsed as an argument into the CLI.
      
    **Example:**
 ```
	//5. Read origin unit
	// Invoke strings.ToUpper function to ensure consistency when reading command line
	// originUnit value provided by the user
	// originUnit is the 2nd argument, os.Args[1], parsed into CLI 
	// the executable is 1st argument , os.Args[0], by default
	originUnit := strings.ToUpper(os.Args[1])
```

6. - [ ] Read Keyboard Input

      - [ ] Print statement that will prompt user to enter desired input that
      will be read from the terminal
      - [ ] Invoke the fmt.Scanln() function, passing the address to the keyboard 
      input variable that as the argument.
      - [ ] Assign the two return values to the variables _ and err respectively.
      - [ ] Next, create an if statement checking if err != nil, and if that's true, 
      invoke the printError() function, passing the error responsible for handling 
      issues with keyboard input that will be read from the terminal as its argument.
      
      **Example**
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
   
7. - [ ] Create function that handles runtime errors - Outside of ```func main()```

	**General Example** - Can be used for any app
	```
		func printError(err error) {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	```

     

8. - [ ] State each core functionality of app and create functions for each  - Outside of ```func main()```

	**Example:**
	
	- **Convert ***current temperature*** from Fahrenheit to Celsius**
	```
		func convertToCelsius(value float64) {
		convertedValue := (value - 32) * 5 / 9
		fmt.Printf("%v F = %.0f C\n", value, convertedValue)
	}
	```
	
	- **Convert ***current temperature*** from Celsius to Fahrenheit**
	```
		func convertToFahrenheit(value float64) {
		convertedValue := (value * 9 / 5) + 32
		fmt.Printf("%v C = %.0f F\n", value, convertedValue)
	}
	```
	
9. - [ ] Generate methods that will utilise core functionality functions

	**Example:**
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



9. - [ ] Prompt Restarting App


10. - [ ] Parse Prompt Answer


      
      
## Summary

      - User compiles and runs executable from command-line parsing [**valid argument**]
      - User is then prompted to enter [**input value**]
      - Then the app reads **keyboard input** ([**input value**]) from terminal and prints out **desired outcome**
      - Finally, the app prompts user to enter [**input value**] to either restart or end the app

## Usage

Compile with `go build -o [executable_name]`

Then, invoke the binary passing as argument [**valid_argument**].
For example:

`./[executable_name] [valid_argument]` to...[Brief decription of what it does]

To verify your work locally for this module, run the following command in a terminal window: 

`go test -v`
