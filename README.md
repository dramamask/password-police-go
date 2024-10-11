# WORK IN PROGRESS #

# Password Police
In development. Not ready for use.

## Example password policy file
```json
{
    "min-length": 4,
    "max-length": 8
}
```

## Running the code
Run the code with `go run .`
Run the tests with `go test -v ./...`

## Developing
Run `go mod tidy` on the command line when using a new third-party package that you haven't used before. This imports the package.

## TODO
Implement all the rules
