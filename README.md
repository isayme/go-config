## Config
Simple configure loader

## Features
- support json/yaml
- default config path from env `CONF_FILE_PATH`;

## Example
```
package conf

import "github.com/isayme/go-config"

type Config struct {
	Addr string `json:"addr" yaml:"addr"`
}

var cfg Config

func Get() *Config {
	config.Parse(&cfg)

	return &cfg
}
```