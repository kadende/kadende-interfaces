package plugin_interface

import (
	"fmt"
)

type PluginType string

func (p PluginType) Validate() (bool) {
	if p == "provider" || p == "flavour" {
		return true
	}
	return false
}

func (p PluginType) ToString() string{
	return fmt.Sprintf("%s",p)
}
