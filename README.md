Requirements.
You may have to install the argparser,
this is simply done by running the command in the terminal: "go get github.com/akamensky/argparse"


If you want to compile.
Os names are: windows = "windows", linux = "linux", mac = "darwin".
To cross compile you have to choose which os you want to compile to by writing: "go env -w GOOS='Os name'".
You have to set your current os by writing: "go env -w GOHOSTOS='Os name'".
Then just write : "go build", and it will be added in the directory as Linter.exe.
The other option is to run : "go install",
this will create an executable that will be usable independent of which file you're in with the name Linter.exe.
To give custom names to these just add "-o newFile.exe" as a flag.


If you want to make a custom ruleset.
Check rules to see structure of the rules within the packages/rulesReader/rules.json file.

Rules:
id - the word you'd like to do something with.
spaceAfter - bool value that is set to true to add a space before and after the id.
indentComingLines - a value where you add how many more or less tabs will be in coming lines.
addBlankLinesBefore - a bool value that add an amount of newlines before the id.
blankLines - set the amount of newlines before the id in addBlankLinesBefore.

Exceptions:
exception - the word you'd like to avoid applying rules to.
after - bool value that is true if the exception is after the rule and false if not.


If you want to run the linter.
You could either run an executable or write: "go run main.go".
Run the program empty to see arguments, arguments inbetween brackets are optional.
In the arguments you have to specify the -f file and -n newFile extentions(acceptable excetions are: ".tex", ".bibz", ".tikz").