# ahapattern

Aha! Pattern is another bad idea that overusing the power of reflection to involve pattern-matching into Go. Here is an example:

```go
package example

import (
    . "github.com/dannypsnl/ahapattern"
)

type Foo struct {
    A int
    B int
}

func main() {
    Match(Foo{A: 1, B: 2}).
        Of(Foo{A: 1, B: 2}, func(f Foo) int { return f.B }).
        Else(func(f Foo) int { return f.A })
}
```

### Problem

- cannot deconstruct structure
- cannot introduce new variable via pattern
- didn't support sub-pattern

### Solution

Use [Haskell](https://www.haskell.org/)