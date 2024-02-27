package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseConfig(t *testing.T) {
	require := require.New(t)

	t.Run("json", func(t *testing.T) {
		type Config struct {
			Addr string `json:"addr"`
		}
		var cfg Config
		err := parseConfig([]byte(`{"addr": "example.com"}`), "json", &cfg)
		require.Nil(err)
		require.Equal(cfg.Addr, "example.com")
	})

	t.Run("yaml", func(t *testing.T) {
		type Config struct {
			Addr string `yaml:"addr"`
		}
		var cfg Config
		err := parseConfig([]byte(`addr: "example.com"`), "yaml", &cfg)
		require.Nil(err)
		require.Equal(cfg.Addr, "example.com")
	})

	t.Run("default value", func(t *testing.T) {
		type Config struct {
			Addr string `yaml:"addr"`
		}
		var cfg Config = Config{
			Addr: "abc",
		}

		err := parseConfig([]byte(``), "yaml", &cfg)
		require.Nil(err)
		require.Equal(cfg.Addr, "abc")
	})
}
