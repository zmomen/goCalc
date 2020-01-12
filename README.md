# goCalc

#### Todo
- High priority
    - Add tests
    - Take command-line input for values
        - Convert values to int64
        - Loop for continous input
        - Limit to single op. (e.g. add)
    - Parse command-line expressions
        - Recognize valid expression
        - Message for invalid expression
- Low priority
    - Non-int arithmetic via cmd-line flag
    - Modulus
    - Check overflow behavior

#### Commits
- Prefix with `[Setup]`, `[Feat]`, `[Bug]`, `[Test]`, `[Refact]`, etc.

#### Install Go

1. Install [Go package](https://golang.org/doc/install?download=go1.13.6.darwin-amd64.pkg)

2. Restart shell

#### Run a Go program

1. Create a workspace `go` in `$HOME`

2. `$ go build -o <EXEC_NAME>`

3. `$ ./<EXEC_NAME>`
