# specw

specw - **w**rappers for yaml/json **spec**ifications

## Usage

```go
package main

import (
	"fmt"
	"github.com/artarts36/specw"
	"gopkg.in/yaml.v3"
	"log/slog"
)

type Config struct {
	URL specw.URL `yaml:"url" json:"url"`
	IP  specw.IP  `yaml:"ip" json:"ip"`
	LogLevel specw.SlogLevel `yaml:"log_level" json:"log_level"`
}

const content = `
url: http://google.com
ip: 192.168.0.1
log_level: info
`

func main() {
	var cfg Config

	yaml.Unmarshal([]byte(content), &cfg)

	fmt.Println(cfg)
}
```