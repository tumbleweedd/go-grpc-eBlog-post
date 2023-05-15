package main

import (
	_ "github.com/lib/pq"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/internal/app"
)

func main() {
	app.Run()
}
