# go-beast-mode
Experimenting with Go build
I worked on this experiment with Go on GitHub Codespace.

Have a look! Here is our statement of problem.

``` zsh
go test -bench main_test.go -v
```

You will notice that
``` zsh
=== RUN   Test_computeEuler
=== RUN   Test_computeEuler/empty
--- PASS: Test_computeEuler (4.19s)
    --- PASS: Test_computeEuler/empty (4.19s)
=== RUN   Test_computeBeastEuler
=== RUN   Test_computeBeastEuler/empty
--- PASS: Test_computeBeastEuler (2.51s)
    --- PASS: Test_computeBeastEuler/empty (2.51s)
PASS
ok      go-beast-demo   6.705s
```

The beast Euler is twice faster than the normal function.
The trick is that we reduce the number of iterations to calculate e^5.
The solution is just move it up to the outermost iteration and pass it through the innermost loop.

Anyway, this should be detected by the compiler.
You can build it like this
```
go build main.go
time ./main 
{133 110 84 27 144}

real    0m4.371s
...
```

We want the following binary to run as fast like the beast Euler one.

Here is what it looks
``` zsh
go build -gcflags=-m main.go
# command-line-arguments
./main.go:18:6: can inline intPow
./main.go:32:19: inlining call to intPow
./main.go:34:20: inlining call to intPow
./main.go:36:21: inlining call to intPow
./main.go:38:22: inlining call to intPow
./main.go:39:22: inlining call to intPow
./main.go:58:36: inlining call to errors.New
./main.go:63:18: inlining call to intPow
./main.go:65:19: inlining call to intPow
./main.go:67:20: inlining call to intPow
./main.go:69:21: inlining call to intPow
./main.go:71:22: inlining call to intPow
./main.go:89:36: inlining call to errors.New
./main.go:95:14: inlining call to fmt.Println
./main.go:99:13: inlining call to fmt.Println
./main.go:58:36: &errors.errorString{...} escapes to heap
./main.go:89:36: &errors.errorString{...} escapes to heap
./main.go:95:14: ... argument does not escape
./main.go:99:13: ... argument does not escape
./main.go:99:13: solution escapes to heap
```