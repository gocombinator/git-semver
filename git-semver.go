package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gocombinator/sh"
	"github.com/gocombinator/sh/git"
)

func main() {

	var major = flag.Bool("major", false, "bump major version")
	var minor = flag.Bool("minor", false, "bump minor version")
	var patch = flag.Bool("patch", false, "bump patch version")
	var prerelease = flag.String("prerelease", "", "use prerelease version")
	var build = flag.String("build", "", "use build version")
	var branch = flag.String("branch", "", "assert branch name")
	var tag = flag.Bool("tag", false, "create git tag")
	var push = flag.Bool("push", false, "perform git push")
	flag.Parse()

	var latestSemver = git.LatestSemver()

	// Maybe bump.
	var nextSemver = latestSemver
	if *major {
		nextSemver = nextSemver.NextMajor()
	}
	if *minor {
		nextSemver = nextSemver.NextMinor()
	}
	if *patch {
		nextSemver = nextSemver.NextPatch()
	}
	nextSemver.Prerelease = *prerelease
	nextSemver.Build = *build

	// If we're not bumping anything, just print and exit.
	if nextSemver == latestSemver {
		fmt.Println(nextSemver)
		os.Exit(0)
	}

	// Maybe confirm branch name we're on.
	if *branch != "" {
		var gitBranch = git.Branch()
		if gitBranch != *branch {
			fmt.Fprintf(os.Stderr, "git branch is %s, not %s\n", gitBranch, *branch)
			os.Exit(1)
		}
	}

	// Maybe tag and push.
	if *tag {

		// Make sure working tree is clean.
		if status := git.Status(); status != "" {
			fmt.Fprintf(os.Stderr, "git working tree not clean:\n%s\n", status)
			os.Exit(1)
		}

		// Create tag.
		sh.Run("git", "tag", nextSemver.String())

		// Maybe push.
		if *push {
			sh.Run("git", "push")
			sh.Run("git", "push", "--tags")
		}

	}

	// Print semver bump.
	fmt.Printf("%s -> %s\n", latestSemver.String(), nextSemver.String())
}
