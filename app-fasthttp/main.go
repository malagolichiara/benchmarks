package main

//CGO_ENABLED=0 GOOS=linux go build -a -o app .

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/valyala/fasthttp"
)

var (
	addr     = flag.String("addr", ":8000", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
	values   = make(map[string]string)
	m        sync.Mutex
)

func main() {
	flag.Parse()

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	fmt.Println("Serving...")

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/plain; charset=utf8")
	if bytes.Compare(ctx.Method(), []byte{'G', 'E', 'T'}) == 0 {
		if bytes.Compare(ctx.Path(), []byte{'/'}) == 0 {
			fmt.Fprintf(ctx, "{\"Hello\": \"World\"}")
		} else if bytes.Compare(ctx.Path()[:7], []byte{'/', 'i', 't', 'e', 'm', 's', '/'}) == 0 {
			idx := string(ctx.Path()[7:])
			m.Lock()
			value, ok := values[idx]
			m.Unlock()
			if ok {
				fmt.Fprintf(ctx, value)
			} else {
				fmt.Fprintf(ctx, "Key not found")
			}
		}
	} else if bytes.Compare(ctx.Method(), []byte{'P', 'O', 'S', 'T'}) == 0 {
		if bytes.Compare(ctx.Path()[:7], []byte{'/', 'i', 't', 'e', 'm', 's', '/'}) == 0 {
			idx := string(ctx.Path()[7:])
			// value := string(ctx.PostBody())
			value := string(ctx.FormValue("value"))
			m.Lock()
			values[idx] = value
			m.Unlock()
			fmt.Fprintf(ctx, "OK")
		}
	}
	fmt.Fprintf(ctx, "")
}
