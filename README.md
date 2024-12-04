# Historian Hysteria

* With each location checked for Chief Historian(CH), it will be marked on their list with a star.
* CH must be in one of the first 50 places they check.
* **You need to help them get fifty stars on their list**.
* Two puzzles made available each day. Second one unlocks when you finish the first one.
* Each puzzle grants 1 star.

## Repo. structure

* Common code used across multiple days is present in `global` package.
* Each day has a dedicated folder.
  + Common code used between the two puzzles is stored in `util.go`.
  + Each puzzle has a dedicated file.
  + The `main.go` invokes both the puzzles and prints the result.
  + `example.txt` contains the sample input provided in puzzle's description.
  + `input.txt` contains the test input that must be solved to earn a gold star.

## Running code

Go to a specific day's directory and run `go run . example.txt` or `go run . input.txt`

