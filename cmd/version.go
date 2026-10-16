package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/Michal-Nithesh/MichalTerm/internal/consts"
)

func buildVersion() string {
	ver := "v0.0.0-dev"
	if info, ok := debug.ReadBuildInfo(); ok {
		switch info.Main.Version {
		case "", "(devel)":
			ver = "v0.0.0-dev"
		default:
			ver = info.Main.Version
		}
	}
	return fmt.Sprintf("%s %s", consts.DisplayName, ver)
}
