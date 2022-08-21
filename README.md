# pushx-go

`pushx-go` is a simple wrapper library for [pushx](https://github.com/robertlestak/pushx). `pushx-go` exports a single function, `Pushx`, which accepts two arguments, `input`, and `args`. `input` is the input `io.Reader`, and `args` is an array containing `pushx` configuration arguments. 

You must have `pushx` installed in order to use `pushx-go`.

See `examples/main.go` for an example of using one codebase and dynamically pushing to multiple services with a single configuration array.

Here is a basic example:

```golang
data := ExampleData{
    Name: "John Doe",
    Age:  42,
    Work: "Software Engineer",
}
jd, err := json.Marshal(data)
if err != nil {
    panic(err)
}
args := []string{
    "-driver",
    "redis-list",
    "-redis-host",
    "localhost",
    "-redis-port",
    "6379",
    "-redis-key",
    "test-redis-list",
}
if err := pushx.Pushx(bytes.NewReader(jd), args); err != nil {
    panic(err)
}
```

Since the entire data layer configuration is contained within the `args` string slice, this can be moved to a configuration layer such as Vault or Consul. If you ever need to change the data layer configuration, you can simply update the `args` configuration, and your code will remain entirely the same.