# specw

```
go get github.com/artarts36/specw
```

specw - **w**rappers for yaml/json **spec**ifications for easy parsing and using.

Usage examples
- [Unmarshal with YAML](#yaml)
- [Load from environment with caarlos0/env](#load-from-environment-with-caarlos0env)

## Wrappers

| Wrapper           | Description                                                                    | Input examples                                             |
|-------------------|--------------------------------------------------------------------------------|------------------------------------------------------------|
| BoolObject[T]     | Generic wrapper for boolean flag or object value                               | `true`, `false`, `{name: name, email: user@mail.ru}`       |
| Color             | Wrapper for _color.RGBA_                                                       | `green`, `#eee`, `#cccccc`                                 |
| Duration          | Wrapper for _time.Duration_                                                    | `500`, `5s`, `1m30s`                                       |
| Env[T]            | Generic wrapper that resolves environment expressions                          | `${MY_VAR}`, `my-value`                                    |
| EnvStrings        | Wrapper for strings list with env expressions                                  | `a,b`, `$A`, `[a,$B]`                                      |
| ExistsFilepath    | Wrapper for filepath with existence check                                      | `/etc/config.yaml`, `./config.yaml`                        |
| File              | Wrapper that reads file content by path                                        | `/run/secrets/token`, `./config.txt`                       |
| GitCommitter      | Wrapper for git commit author with name and email                              | `name <user@mail.ru>`, `{name: name, email: user@mail.ru}` |
| IP                | Wrapper for _net.IP_                                                           | `192.168.0.1`, `2001:db8::1`                               |
| JSONFile[T]       | Generic wrapper that reads JSON file by path and decodes it into _T_           | `./payload.json`, `/etc/app/config.json`                   |
| OneOrMany[T]      | Generic wrapper that accepts one value or many values and stores them as slice | `{val: 1}`, `[{val: 1}, {val: 2}]`                         |
| PositiveNumber[T] | Generic wrapper that validates number is positive                              | `1`, `42`, `0.5`                                           |
| SlogLevel         | Wrapper for _slog.Level_                                                       | `error`, `warn`, `info`, `debug`                           |
| URL               | Wrapper for _url.URL_                                                          | `https://google.com`, `http://localhost:8080`              |
| YAMLFile[T]       | Generic wrapper that reads YAML file by path and decodes it into _T_           | `./payload.yaml`, `/etc/app/config.yaml`                   |

### Usage

#### YAML

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

#### Load from environment with `caarlos0/env`

```go
package main

import (
	"fmt"
	"github.com/artarts36/specw"
	env "github.com/caarlos0/env/v11"
)

type Config struct {
	ServerIP specw.IP        `env:"SERVER_IP,required"`
	Timeout  specw.Duration  `env:"TIMEOUT" envDefault:"5s"`
	LogLevel specw.SlogLevel `env:"LOG_LEVEL" envDefault:"info"`
}

func main() {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	fmt.Println("server ip:", cfg.ServerIP.String())
	fmt.Println("timeout:", cfg.Timeout.Value)
	fmt.Println("log level:", cfg.LogLevel.String())
}
```
