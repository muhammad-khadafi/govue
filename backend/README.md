# POC-GO
Project for learning and exploring Go Language

## Code Convention
### Files
1. Go follows a convention where source files are all lower case with underscore separating multiple words.
2. Compound file names are separated with _
3. File names that begin with “.” or “_” are ignored by the go tool
4. Files with the suffix _test.go are only compiled and run by the go test tool.

### Functions and Methods
1. Use camel case 
2. exported functions should start with uppercase

### Constants
Constant should use all capital letters and use underscore _ to separate words.

### Variables
1. Generally, use relatively simple (short) name.
2. Consistent naming style should be used the entire source code
3. If variable type is bool, its name should start with Has, Is, Can or Allow, etc.
4. Single letter represents index: i, j, k

### The Unwritten Rules in Go
1. Shorter variable names
2. Acronyms such as API, HTTP, etc. or names like ID and DB. Conventionally, we keep these words in their original form : 
```
userID instead of userId 
productAPI instead of productApi
```