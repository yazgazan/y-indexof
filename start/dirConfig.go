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

type DirConfig struct {
	View string
}

type DirConfigError struct {
	What string
}

func (e *DirConfigError) Error() string {
	return e.What
}

func (c *DirConfig) CheckDirConfig() error {
	return nil
}

func MakeDirConfig(config Config) *DirConfig {
	return &DirConfig{
		View: config.IndexView,
	}
}

func ReadDirConfig(path string, config Config) (*DirConfig, error) {
	conf := MakeDirConfig(config)

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
