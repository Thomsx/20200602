package main

import (
	"io"
)

// URL: /
// 网站首页
func indexHandler(handler *Handler) {
	io.WriteString(handler.ResponseWriter, "working")
}
