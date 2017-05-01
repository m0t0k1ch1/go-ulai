# go-ulai

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/go-ulai?status.svg)](https://godoc.org/github.com/m0t0k1ch1/go-ulai) [![wercker status](https://app.wercker.com/status/dc058d4a4600a7008b41a55a44f2cc15/s/master "wercker status")](https://app.wercker.com/project/byKey/dc058d4a4600a7008b41a55a44f2cc15)

a [User Local AI API](http://ai.userlocal.jp/document/free/top) client for golang

## Example

``` go
package main

import (
	"context"
	"fmt"
	"log"

	ulai "github.com/m0t0k1ch1/go-ulai"
)

func main() {
	client := ulai.NewClient()
	client.SetKey("your API key")

	res, err := client.Chat(context.Background(), "あなたは人工知能？")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
```

```
人工知能というより人工無脳かな
```
