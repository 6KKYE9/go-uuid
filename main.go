// go-uuid 生成 RFC 4122 版本 4 的 UUID，零依赖（只用标准库 crypto/rand）。
// 子命令：
//   (无参数)        生成一个 UUID v4
//   -n <数量>       批量生成
//   -prefix <字符串>  每行前缀
// 例：
//   go run .
//   go run . -n 5
//   go run . -n 3 -prefix order_
package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	n := flag.Int("n", 1, "生成数量")
	prefix := flag.String("prefix", "", "每行前缀")
	flag.Parse()

	if *n < 1 {
		*n = 1
	}
	for i := 0; i < *n; i++ {
		id, err := newUUIDv4()
		if err != nil {
			fmt.Fprintf(os.Stderr, "生成失败: %v\n", err)
			os.Exit(1)
		}
		if *prefix != "" {
			fmt.Printf("%s%s\n", *prefix, id)
		} else {
			fmt.Println(id)
		}
	}
}

// newUUIDv4 生成一个版本 4 的 UUID 字符串，形如 xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
func newUUIDv4() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	// 设置版本(4)和变体位。
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hex.EncodeToString(b[0:4]),
		hex.EncodeToString(b[4:6]),
		hex.EncodeToString(b[6:8]),
		hex.EncodeToString(b[8:10]),
		hex.EncodeToString(b[10:16]),
	), nil
}
