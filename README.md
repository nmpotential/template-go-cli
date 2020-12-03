# go-cli-template
Template that guides you through building a command line application which has the functionality to parse  CLI arguments 
and read keyboard inputs from the terminal in Go.

## Steb by Step Guide to: Building a CLI app that parses CLI arguments and reads keyboard inputs from the terminal in Go

1. - [ ] Briefly summaries how the apps supposed to work and put in **summary** of README.md - This will be updated as 
      functionality and implementation are refined:
      
      - User compiles and runs executable from command-line parsing [**valid argument**]
      - User is then prompted to enter [**input value**]
      - Then the app prints out [**desired outcome**]
      
      **Example:**
      - User compiles and runs executable from command-line parsing **the unit of temperature we want to convert from** 
      - User is then prompted to enter **input value**
      - Then the app prints out **converted temperature between Fahrenheit and Celsius**.
      
2. - [ ] Setup Environment
      
      - [ ] Create main.go file
      - [ ] Create [**test**] file
      - [ ] Compile executable with `go build -o [executable_name]` command using assigning relevant executable name - add command 
      for compiling to README under Usage.
      
      **Example**:
      - [ ] Create main.go file
      - [ ] Create **main_test.go** file
      - [ ] Compile executable with `go build -o temp` - add command for compiling to README under Usage.
      
3. - [ ] Declare the arguments(s) that will be parsed into the CLI and the keyboard input that will be read from the terminal

      **Example**:
      ```
      package main

      var originUnit string	// the arguments(s) that will be parsed into the CLI
      var originValue float64	// the keyboard input that will be read from the terminal

      var err error //Error output

      var errInvalidArguments = errors.New("Invalid arguments")	// Error output for invalid arguments
      var errReadingInput = errors.New("Error reading input")	// Error output for issue with reading input
      ```
      
3. - [ ] Declare the keyboard input that will be read from the terminal

      **Example**:
      ```
      package main

      var originUnit string	// the arguments(s) that will be parsed into the CLI
      var originValue float64	// the keyboard input that will be read from the terminal

      var shouldConvertAgain string //Expected Output

      var err error //Error output

      var errInvalidArguments = errors.New("Invalid arguments")	// Error output for invalid arguments
      var errReadingInput = errors.New("Error reading input")	// Error output for issue with reading input
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
