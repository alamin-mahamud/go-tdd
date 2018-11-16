# Learn Go With Tests
![](https://travis-ci.org/alamin-mahamud/go-tdd.svg?branch=master)

``` bash
# Debugger
go get -u github.com/derekparker/delve/cmd/dlv

# lint
go get -u github.com/alecthomas/gometalinter
gometalinter --install
```

## Workflow
1. Write a Test.
2. Make the compiler PASS.
3. Run the test, see that it fails and check the ERROR msg is meaningful.
4. Write Enough Code to make the test pass.
5. Refactor.
