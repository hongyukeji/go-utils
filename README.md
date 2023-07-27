# go-utils

### 安装

To download the specific tagged release, run:
```bash
go get -u github.com/hongyukeji/go-utils@v1.0.0
```
Import it in your program as:
```go
import "github.com/hongyukeji/go-utils"
```
It requires Go 1.11 or later due to usage of Go Modules.

Refer to the documentation here:
https://github.com/hongyukeji/go-utils

The rest of this document describes the the advances in v1 and a list of
breaking changes for users that wish to upgrade from an earlier version.

### 使用

#### string 字符串

##### TextTemplateFormat()

```go
package main

import (
	"fmt"
	"github.com/hongyukeji/go-utils/utils"
)

func main() {
	text := "您的验证码为: ${code}"
	params := map[string]string{"code": "1234"}
	template := utils.TextTemplateFormat(text, params)
	fmt.Println(template)
}

```
