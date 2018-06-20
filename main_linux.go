//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/moby/sys/reexec"

	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
	hostlocal "github.com/containernetworking/plugins/plugins/ipam/host-local"
	"github.com/containernetworking/plugins/plugins/main/bridge"
	"github.com/containernetworking/plugins/plugins/main/loopback"
	"github.com/containernetworking/plugins/plugins/meta/bandwidth"
	"github.com/containernetworking/plugins/plugins/meta/firewall"
	"github.com/containernetworking/plugins/plugins/meta/flannel"
	"github.com/containernetworking/plugins/plugins/meta/portmap"
)

func main() {
	os.Args[0] = filepath.Base(os.Args[0])
	reexec.Register("bandwidth", bandwidth.Main)
	reexec.Register("bridge", bridge.Main)
	reexec.Register("firewall", firewall.Main)
	reexec.Register("flannel", flannel.Main)
	reexec.Register("host-local", hostlocal.Main)
	reexec.Register("loopback", loopback.Main)
	reexec.Register("portmap", portmap.Main)
	if !reexec.Init() {
		_, _ = fmt.Fprintln(os.Stderr, bv.BuildString("plugins"))
		_, _ = fmt.Fprintf(os.Stderr, "CNI protocol versions supported: %s\n", strings.Join(version.All.SupportedVersions(), ", "))
	}
}
