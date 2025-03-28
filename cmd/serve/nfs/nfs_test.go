// The serving is tested in cmd/nfsmount - here we test anything else
package nfs

import (
	"testing"

	_ "github.com/rclone/rclone/backend/local"
	"github.com/rclone/rclone/cmd/serve/servetest"
	"github.com/rclone/rclone/fs/rc"
)

func TestRc(t *testing.T) {
	servetest.TestRc(t, rc.Params{
		"type": "nfs",
		"opt": rc.Params{
			"ListenAddr": "localhost:0",
		},
		"vfsOpt": rc.Params{
			"CacheMode": "off",
		},
	})
}
