package model

import (
	"github.com/jasperstritzke/cubid/pkg/util/fileutil"
)

type VersionValue struct {
	BuildURL string `json:"url"`
	Display  string `json:"display"`
	Proxy    bool   `json:"proxy"`
	Mention  string `json:"mention"`
}

func (v *VersionValue) IsProxy() bool {
	return v.Proxy
}
func (v *VersionValue) IsServer() bool {
	return !v.Proxy
}

func (v *VersionValue) DownloadTo(pth string) error {
	return fileutil.DownloadFile(pth, v.BuildURL)
}

func (v VersionValue) ProxyAsString() string {
	var proxyString = "Server"
	if v.Proxy {
		proxyString = "Proxy"
	}

	return proxyString
}

type Versions struct {
	BungeeCord VersionValue
	Waterfall  VersionValue
	Paper17    VersionValue
	Paper16    VersionValue
	Empty      VersionValue
}

var Version = Versions{
	BungeeCord: VersionValue{
		BuildURL: "https://ci.md-5.net/job/BungeeCord/lastSuccessfulBuild/artifact/bootstrap/target/BungeeCord.jar",
		Display:  "BungeeCord",
		Mention:  "Last build",
		Proxy:    true,
	},
	Waterfall: VersionValue{
		BuildURL: "https://papermc.io/api/v2/projects/waterfall/versions/1.17/builds/451/downloads/waterfall-1.17-451.jar",
		Display:  "Waterfall-1.17",
		Mention:  "build 451 (probably outdated)",
		Proxy:    true,
	},
	Paper17: VersionValue{
		BuildURL: "https://papermc.io/api/v2/projects/paper/versions/1.17.1/builds/334/downloads/paper-1.17.1-334.jar",
		Display:  "Paper-1.17",
		Mention:  "build 334 (probably outdated)",
		Proxy:    false,
	},
	Paper16: VersionValue{
		BuildURL: "https://papermc.io/api/v2/projects/paper/versions/1.16.5/builds/788/downloads/paper-1.16.5-788.jar",
		Display:  "Paper-1.16",
		Mention:  "build 788 (probably outdated)",
		Proxy:    false,
	},
	Empty: VersionValue{
		Proxy: false,
	},
}

var VersionsAsArray = []VersionValue{Version.BungeeCord, Version.Waterfall, Version.Paper17, Version.Paper16}
