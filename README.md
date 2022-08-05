# This issue has been resolved

Issue https://github.com/dolthub/vitess/issues/174 has been resolved by merging
https://github.com/dolthub/vitess/pull/175, and dolthub/go-mysql-server no
longer has any flag collision issues with Vitess.


# For posterity

Prior to commit `c7560b3fabf0d1fcdaf845662b921c5187eb9320`, the
`go-mysql-server` package defined flags that conflicted with Vitess.

```
go test src/my_test.go
```

fails with

```
$ go test src/my_test.go 
/tmp/go-build3353698422/b001/my.test flag redefined: log_err_stacks
panic: /tmp/go-build3353698422/b001/my.test flag redefined: log_err_stacks

goroutine 1 [running]:
flag.(*FlagSet).Var(0xc00011c120, {0xc44218, 0x10f2e7c}, {0xb2be81, 0xe}, {0xb374b2, 0x1b})
        /home/ryanpbrewster/opt/go/src/flag/flag.go:879 +0x2f4
flag.BoolVar(...)
        /home/ryanpbrewster/opt/go/src/flag/flag.go:638
vitess.io/vitess/go/vt/vterrors.init.0()
        /home/ryanpbrewster/go/pkg/mod/vitess.io/vitess@v0.14.1/go/vt/vterrors/vterrors.go:103 +0x4e
FAIL    command-line-arguments  0.006s
FAIL
```
