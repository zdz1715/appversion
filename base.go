package appversion

import (
	"runtime/debug"
	"time"
)

// Base version information.
//
// This is the fallback data used when version information from git is not
// provided via go ldflags. It provides an approximation of the version for
// ad-hoc builds (e.g. `go build`) that cannot get the version
// information from git.
var (
	// if no version is set, get it from git tag,
	// output of $(git describe --tags --abbrev=0 --exact-match)
	version = ""

	// sha1 from git, output of $(git rev-parse HEAD)
	gitCommit = "" // sha1 from git, output of $(git rev-parse HEAD)

	// state of git tree, either "clean" or "dirty"
	// output of $(test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
	gitTreeState = "" // state of git tree, either "clean" or "dirty"

	// build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	buildDate = ""

	// app name
	name = ""
)

func SetVersion(ver string) {
	version = ver
}

func SetName(n string) {
	name = n
}

func init() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	var buildTime time.Time

	if buildDate == "" {
		buildDate, buildTime = getBuildDate(bi)
	}

	if gitCommit == "" {
		gitCommit = getCommit(bi)
	}

	if version == "" {
		version = getGitVersion(bi)
		if version == "" {
			version = "v0.0.0"
			if !buildTime.IsZero() {
				version += "-" + buildTime.Format("20060102150405")
			}
			if len(gitCommit) >= 12 {
				version += "-" + gitCommit[:12]
			}
		}
	}

	if gitTreeState == "" {
		gitTreeState = getDirty(bi)
	}

}

func getGitVersion(bi *debug.BuildInfo) string {
	if bi.Main.Version == "(devel)" || bi.Main.Version == "" {
		return ""
	}

	return bi.Main.Version
}

func getCommit(bi *debug.BuildInfo) string {
	return getKey(bi, "vcs.revision")
}

func getBuildDate(bi *debug.BuildInfo) (string, time.Time) {
	buildTime := getKey(bi, "vcs.time")
	t, err := time.Parse(time.RFC3339, buildTime)
	if err != nil {
		return "", t
	}
	return t.Format(time.RFC3339), t
}

func getKey(bi *debug.BuildInfo, key string) string {
	if bi == nil {
		return ""
	}
	for _, iter := range bi.Settings {
		if iter.Key == key {
			return iter.Value
		}
	}
	return ""
}

func getDirty(bi *debug.BuildInfo) string {
	modified := getKey(bi, "vcs.modified")
	if modified == "true" {
		return "dirty"
	}
	if modified == "false" {
		return "clean"
	}
	return ""
}
