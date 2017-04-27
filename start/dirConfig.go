/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * <yazgazan@gmail.com> wrote this file. As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return.
 * ----------------------------------------------------------------------------
 */

package start

import (
	"github.com/BurntSushi/toml"

	"io/ioutil"
)

type dirConfig struct {
	View string
}

func (c *dirConfig) CheckDirConfig() error {
	return nil
}

func newDirConfig(config Config) *dirConfig {
	return &dirConfig{
		View: config.IndexView,
	}
}

func readDirConfig(path string, config Config) (*dirConfig, error) {
	conf := newDirConfig(config)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	_, err = toml.Decode(string(content), &conf)
	if err != nil {
		return nil, err
	}
	err = conf.CheckDirConfig()
	if err != nil {
		return nil, err
	}

	return conf, nil
}
