package gvite_plugins

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"gopkg.in/urfave/cli.v1"

	"github.com/vitelabs/go-vite/cmd/utils"
	"github.com/vitelabs/go-vite/version"
)

var (
	versionCommand = cli.Command{
		Action:    utils.MigrateFlags(versionAction),
		Name:      "version",
		Usage:     "Print version numbers",
		ArgsUsage: " ",
		Category:  "MISCELLANEOUS COMMANDS",
		Description: `
The output of this command is supposed to be machine-readable.
`,
	}
	licenseCommand = cli.Command{
		Action:    utils.MigrateFlags(licenseAction),
		Name:      "license",
		Usage:     "Display license information",
		ArgsUsage: " ",
		Category:  "MISCELLANEOUS COMMANDS",
	}
)

func versionAction(ctx *cli.Context) error {
	fmt.Println(strings.Title("gvite"))
	fmt.Println("Version:", version.VITE_BUILD_VERSION)
	fmt.Println("Git Commit:", version.VITE_COMMIT_VERSION)
	fmt.Println("Architecture:", runtime.GOARCH)
	//fmt.Println("Network Id:", ctx.GlobalInt())
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Printf("GOPATH=%s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT=%s\n", runtime.GOROOT())
	return nil
}

func licenseAction(_ *cli.Context) error {
	fmt.Println(`GVite is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GVite is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received chain copy of the GNU General Public License
along with gvite. If not, see <http://www.gnu.org/licenses/>.`)
	return nil
}
