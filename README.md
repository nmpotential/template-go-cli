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
      - [ ] Compile executable with `go build -o temp` - add command for compiling to README under Usage.
      
3. - [ ] Declare variables for the following:
      **Input**
      - The arguments(s) that will be parsed into the CLI
      - The keyboard input that will be read from the terminal
      - The input functionality that enables the app to either restart or end 
      
      **Output**
      - The error responsible for handling invalid arguments(s) that will be parsed into the CLI
      - The error responsible for handling issues with keyboard input that will be read from the terminal
      - The error responsible for handling all other potential errors the app may encounter 

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
      var errReadingInput = errors.New("Error reading input")   // Output: The error responsible for handling issues with keyboard input that will be read from the        terminal

      ```
      
      
      
## Summary

      - User compiles and runs executable from command-line parsing [**valid argument(s)**]
      - User is then prompted to enter [**input value**]
      - Then the app prints out [**desired outcome**]

## Usage

Compile with `go build -o [executable_name]`

Then, invoke the binary passing as argument [**valid_argument**].
For example:

`./[executable_name] [valid_argument]` to...[Brief decription of what it does]

To verify your work locally for this module, run the following command in a terminal window: 

`go test -v`
