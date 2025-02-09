### Installation
- For the first time
```
go mod init mutation-demo
# go install github.com/zimmski/go-mutesting/cmd/go-mutesting@latest # Fucking out date
go install github.com/go-gremlins/gremlins/cmd/gremlins@latest

```
- After installed (clone), only need to run: `go mod tidy`

### Mutation test
- Before test: `go test ./calculator/`
```
ok      calculator-project/calculator   0.001s`
```
- Export path first in `bash.rc`
```
export PATH=$PATH:$(go env GOPATH)/bin
```

- Run:  `gremlins unleash ./calculator`
```
Starting...
Gathering coverage... go: no module dependencies to download
done in 289.348841ms
      KILLED ARITHMETIC_BASE at calculator.go:23:14
      KILLED ARITHMETIC_BASE at calculator.go:14:14
      KILLED ARITHMETIC_BASE at calculator.go:9:14
      KILLED CONDITIONALS_NEGATION at calculator.go:20:10

Mutation testing completed in 230 milliseconds 990 microseconds
Killed: 4, Lived: 0, Not covered: 0
Timed out: 0, Not viable: 0, Skipped: 0
Test efficacy: 100.00%
Mutator coverage: 100.00%
```

- Clean cache: `go clean -testcache`
- 