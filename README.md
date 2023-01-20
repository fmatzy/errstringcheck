# errstringcheck

errstringcheck checks error message format.

## Install

```command
go install github.com/fmatzy/errstringcheck/cmd/errstringcheck@latest
```

## Usage

```command
errstringcheck ./...
```

## Examples

```go
// Bad
fmt.Errof("foo error=%v", err)
fmt.Errof("bar error=%w", err)
fmt.Errof("error %w message", err)

// Good
fmt.Errof("foo error: %v", err)
fmt.Errof("bar error: %w", err)
```

## Options

- `-wraponly`: Only allow `%w` verb to format error variable.
