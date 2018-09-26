# Comparison package
Provide a comparison method between different types with implicit conversion.
Below types are supported:
* string
* bool
* int8
* int16
* int32
* int
* int64
* float32
* float64

## Examples

### Int & String

```golang
"1" == 1
"111" == 111
"111" >= 111
"0" <= 111
```

### Int & Float
```golang
int64(1) == float32(1.0)
int64(100000000) > float32(1.0)
```

### Float & String
```golang
float32(1) == "1"
```

### Bool & String
```golang
bool(true) == "true"
bool(true) == 1
bool(false) == "false"
bool(false) == 0
```
