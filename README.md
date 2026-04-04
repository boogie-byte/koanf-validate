# koanf-validate

Rule-based validation for [koanf](https://github.com/knadh/koanf) configuration.

## Installation

```sh
go get github.com/boogie-byte/koanf-validate
```

## Usage

```go
import (
    "fmt"

    "github.com/knadh/koanf/v2"
    validate "github.com/boogie-byte/koanf-validate"
)

k := koanf.New(".")
// ... load configuration ...

errs := validate.Validate(k,
    validate.Rule("db.host", validate.Required()),
    validate.Rule("db.port", validate.Required()),
)

for _, err := range errs {
    fmt.Println(err)
}
```

## Wildcard selectors

Selectors support `*` wildcards to validate multiple keys at once:

```go
errs := validate.Validate(k,
    validate.Rule("services.*.host", validate.Required()),
    validate.Rule("services.*.port", validate.Required()),
)
```

Multiple wildcards are supported: `regions.*.services.*.port`.

## Custom predicates

A `Predicate` is a `func(val any) error`. Write your own to extend validation:

```go
func PortRange() validate.Predicate {
    return func(val any) error {
        port, ok := val.(int64)
        if !ok {
            return fmt.Errorf("expected int64, got %T", val)
        }
        if port < 1 || port > 65535 {
            return fmt.Errorf("port %d out of range", port)
        }
        return nil
    }
}
```

## Error handling

`Validate` returns `[]error`. Each error is a `ValidationError` that can be inspected:

```go
for _, err := range errs {
    var ve validate.ValidationError
    if errors.As(err, &ve) {
        fmt.Println("field:", ve.FieldName())
    }
}
```
