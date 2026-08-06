package main

import (
	"strings"
	"testing"
)

func TestNewUUIDv4(t *testing.T) {
	for i := 0; i < 100; i++ {
		id, err := newUUIDv4()
		if err != nil {
			t.Fatal(err)
		}
		// 版本位在第 14 个字符（索引 14，即第三段的第一个字符）应为 '4'
		if len(id) != 36 || id[14] != '4' {
			t.Fatalf("版本位错误: %q", id)
		}
		// 变体位第三段首字符应为 8/9/a/b
		c := id[19]
		if c != '8' && c != '9' && c != 'a' && c != 'b' {
			t.Fatalf("变体位错误: %q", id)
		}
		// 分隔符数量应为 4
		if strings.Count(id, "-") != 4 {
			t.Fatalf("分隔符数量异常: %q", id)
		}
	}
}

func TestFormatUUID(t *testing.T) {
	base := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	if got := formatUUID(base, true, "-"); got != strings.ToUpper(base) {
		t.Errorf("upper 失败: %q", got)
	}
	if got := formatUUID(base, false, "none"); strings.Contains(got, "-") {
		t.Errorf("none 应去掉分隔符: %q", got)
	}
	if got := formatUUID(base, false, ""); strings.Contains(got, "-") {
		t.Errorf("空分隔符应去掉: %q", got)
	}
	if got := formatUUID(base, false, ":"); got != strings.ReplaceAll(base, "-", ":") {
		t.Errorf("自定义分隔符失败: %q", got)
	}
}
