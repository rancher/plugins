//go:build windows
// +build windows

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
	winbridge "github.com/containernetworking/plugins/plugins/main/windows/win-bridge"
	winoverlay "github.com/containernetworking/plugins/plugins/main/windows/win-overlay"
	"github.com/containernetworking/plugins/plugins/meta/flannel"
)

func main() {
	os.Args[0] = filepath.Base(os.Args[0])
	reexec.Register("flannel", flannel.Main)
	reexec.Register("host-local", hostlocal.Main)
	reexec.Register("win-bridge", winbridge.Main)
	reexec.Register("win-overlay", winoverlay.Main)
	reexec.Register("flannel.exe", flannel.Main)
	reexec.Register("host-local.exe", hostlocal.Main)
	reexec.Register("win-bridge.exe", winbridge.Main)
	reexec.Register("win-overlay.exe", winoverlay.Main)
	if !reexec.Init() {
		_, _ = fmt.Fprintln(os.Stderr, bv.BuildString("plugins"))
		_, _ = fmt.Fprintf(os.Stderr, "CNI protocol versions supported: %s\n", strings.Join(version.All.SupportedVersions(), ", "))
	}
}
