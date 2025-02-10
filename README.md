### Installation
- For the first time
```
go mod init mutation-demo
# go install github.com/zimmski/go-mutesting/cmd/go-mutesting@latest # Fucking out date
go install github.com/go-gremlins/gremlins/cmd/gremlins@latest

```
- After installed (clone), only need to run: `go mod tidy`

### Mutation test - Killed
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


### Mutation test - lived
- I create some example in folder `calculator.mutants`
- Output expected of cmd: `gremlins unleash calculator.mutants/`
```
Starting...
Gathering coverage... go: no module dependencies to download
done in 91.215846ms
 NOT COVERED ARITHMETIC_BASE at calculator.go:9:16
 NOT COVERED ARITHMETIC_BASE at calculator.go:9:25
 NOT COVERED INVERT_NEGATIVES at calculator.go:9:16
 NOT COVERED ARITHMETIC_BASE at calculator.go:20:23
       LIVED CONDITIONALS_BOUNDARY at calculator.go:8:10
       LIVED CONDITIONALS_NEGATION at calculator.go:8:19
 NOT COVERED CONDITIONALS_BOUNDARY at calculator.go:28:14
 NOT COVERED CONDITIONALS_NEGATION at calculator.go:28:14
 NOT COVERED INVERT_NEGATIVES at calculator.go:29:20
 NOT COVERED ARITHMETIC_BASE at calculator.go:29:20
       LIVED CONDITIONALS_BOUNDARY at calculator.go:19:19
 NOT COVERED CONDITIONALS_BOUNDARY at calculator.go:37:10
 NOT COVERED CONDITIONALS_NEGATION at calculator.go:37:10
 NOT COVERED INVERT_NEGATIVES at calculator.go:38:16
 NOT COVERED ARITHMETIC_BASE at calculator.go:38:16
       LIVED CONDITIONALS_NEGATION at calculator.go:8:10
       LIVED CONDITIONALS_NEGATION at calculator.go:16:10
       LIVED CONDITIONALS_NEGATION at calculator.go:16:20
       LIVED CONDITIONALS_NEGATION at calculator.go:19:19
       LIVED CONDITIONALS_BOUNDARY at calculator.go:8:19
       LIVED ARITHMETIC_BASE at calculator.go:22:14
       LIVED CONDITIONALS_BOUNDARY at calculator.go:19:10
       LIVED ARITHMETIC_BASE at calculator.go:11:14
       LIVED CONDITIONALS_NEGATION at calculator.go:19:10
       LIVED CONDITIONALS_NEGATION at calculator.go:27:10
       LIVED ARITHMETIC_BASE at calculator.go:33:14

Mutation testing completed in 195 milliseconds 868 microseconds
Killed: 0, Lived: 14, Not covered: 12
Timed out: 0, Not viable: 0, Skipped: 0
Test efficacy: 0.00%
Mutator coverage: 53.85%
```