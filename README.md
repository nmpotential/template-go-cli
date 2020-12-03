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
      - Then the app reads **keyboard input** (**current temperature**) from terminal and prints out **converted temperature between Fahrenheit and Celsius**.
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

      - [ ] Create a validation rule to check if arguments are valid
      - [ ] If invalid, invoke the printError() function, passing **the error variable** 
      **responsible for handling invalid arguments(s) that will be parsed into the CLI**
      
      **Example:**
 ```
      func main() {
	// We use the built-in len function to check if the length of os.Args is different than 2.
	if len(os.Args) != 2 {
		// If argument is not valid we invoke the printError() function, passing errInvalidArguments as the argument.
		printError(errInvalidArguments)
	}
      
```

      
      
      

5. - [ ] Read Argument(s)

6. - [ ] Read Keyboard Inputs

7. - [ ] Func

8. - [ ] Prompt Restarting App

9. - [ ] Parse Prompt Answer


      
      
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
