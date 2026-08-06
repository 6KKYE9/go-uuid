# go-uuid

零依赖的 UUID v4 生成器（仅用 `crypto/rand`，符合 RFC 4122）。

## 功能

- 生成单个 / 批量 UUID v4
- 支持每行前缀（适合订单号、追踪号等场景）

## 用法

```bash
go run .                  # 生成一个
go run . -n 5             # 批量生成 5 个
go run . -n 3 -prefix order_   # 带前缀
```

## 示例输出

```
$ go run . -n 2
f47ac10b-58cc-4372-a567-0e02b2c3d479
9a8b7c6d-5e4f-4a3b-2c1d-0e9f8a7b6c5d
```
