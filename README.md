# go-uuid

生成 UUID v4，用 `crypto/rand`，符合 RFC 4122。

```bash
go run .                  # 生成一个
go run . -n 5             # 批量生成 5 个
go run . -n 3 -prefix order_   # 带前缀，适合订单号、追踪号
```

加前缀时会逐行输出。

```
$ go run . -n 2
f47ac10b-58cc-4372-a567-0e02b2c3d479
9a8b7c6d-5e4f-4a3b-2c1d-0e9f8a7b6c5d
```
