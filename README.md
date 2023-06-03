# go_run_in_terminal
Simple wrapper to execute golang program from another golang program

## Parameters
`string` **execName** - name of the executable

`[]string` **args** - list of parameters

`string` **dirPath** - location of executable

`bool` **noWait** - trigger waiting for execution to complete


## Examples

### Execute with "go run"
```
params := "param1=1 param2=tst"

res, err := Run("go", []string{
	"run",
	"main.go",
	"functions.go",
	"param0=0",
	params,
}, "../test_program", true)
```
> cd ../test_program ; go run main.go functions.go param0=0 param1=1 param2=tst &

### Execute binary
```
params := "param1=1 param2=tst"

res, err := Run("./search_tw_accounts", []string{
	"./test_program",
    "param0=0",
	params,
}, "../test_program", false)
```
> ./../test_program/test_program param0=0 param1=1 param2=tst
