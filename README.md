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