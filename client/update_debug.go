// +build !release,!autoupdate

package client

import (
	"github.com/agurha/tunnel/client/mvc"
)

// no auto-updating in debug mode
func autoUpdate(state mvc.State, token string) {
}
