## Vet Tool

You can use the go vet command to check for errors: `go vet error.go`.

The output:
```
# command-line-arguments
./error.go:9:2: fmt.Printf format %s has arg movie_year of wrong type int
```
