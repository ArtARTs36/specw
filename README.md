# specw

```
go get github.com/artarts36/specw
```

specw - **w**rappers for yaml/json **spec**ifications for easy parsing and using.

## Wrappers

| Wrapper   | Description                                                          | Input examples                                          |
|-----------|----------------------------------------------------------------------|---------------------------------------------------------|
| Color     | Wrapper for _color.RGBA_                                             | `green`, `#eee`, `#cccccc`                              |
| Duration  | Wrapper for _time.Duration_                                          | `500`, `5s`                                             |
| Env       | Generic wrapper, which fetch value from environment variable         | `${MY_VAR}`, `my-value`                                 |
| IP        | Wrapper for _net.IP_                                                 | `192.168.0.1`                                           |
| OneOrMany | Generic wrapper that allows you to put one or more values into a key | Value or values: `key: {val: 1}` or `key: {val: [1,2]}` |
| SlogLevel | Wrapper for _slog.Level_                                             | `error`, `warn`, `info`, `debug`                        |
| URL       | Wrapper for _url.URL_                                                | `https://google.com`                                    |

## Usage

```go
package main

import (
	"fmt"
	"github.com/artarts36/specw"
	"gopkg.in/yaml.v3"
)

type Config struct {
	URL      specw.URL       `yaml:"url" json:"url"`
	IP       specw.IP        `yaml:"ip" json:"ip"`
	LogLevel specw.SlogLevel `yaml:"log_level" json:"log_level"`

	EnvString specw.Env[string]   `yaml:"env_string"`
	EnvIP     specw.Env[specw.IP] `yaml:"env_ip"`

	Slice specw.OneOrMany[string] `yaml:"slice"`
}

const content = `
url: http://google.com
ip: 192.168.0.1
log_level: info
env_string: "test"
env_ip: "${IP}"
slice: abcd
`

func main() {
	var cfg Config

	_ = yaml.Unmarshal([]byte(content), &cfg)

	fmt.Println(cfg)
}
```
