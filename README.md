
# Linter

This is a simple LaTeX linter

## Compile
Os names are: windows = "windows", linux = "linux", mac = "darwin".

To cross compile you have to choose which os you want to compile to by writing: "go env -w GOOS='Os name'".  
You have to set your current os by writing: "go env -w GOHOSTOS='Os name'".  
Then just write : "go build", and it will be added in the directory as Linter.exe.  
The other option is to run : "go install",  
this will create an executable that will be usable independent of which file you're in with the name Linter.exe.  
To give custom names to these just add "-o newFile.exe" as a flag.  


## Custom Ruleset
Check rules to see structure of the rules within the packages/rulesReader/rules.json file.

#### Rules
* id - the word you'd like to do something with.
* spaceAfter - bool value that is set to true to add a space before and after the id.
* indentComingLines - a value where you add how many more or less tabs will be in coming lines.
* addBlankLinesBefore - a bool value that add an amount of newlines before the id.
* blankLines - set the amount of newlines before the id in addBlankLinesBefore.

#### Exceptions
* exception - the word you'd like to avoid applying rules to.
* after - bool value that is true if the exception is after the rule and false if not.


## Running the linter
You could either run an executable or write: "go run main.go".  
Run the program empty to see arguments, arguments within brackets are optional.  
Accepted extentions to -f and -n are ".tex", ".tikz" and ".bibz" and to -r ".json" and ".yaml"