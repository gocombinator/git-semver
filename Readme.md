## Summary

`git-semver` command line tool helps to manage semantic versions in git repository.

Like `git` it works with repository in current directory or any subdirectories.

## Install

```sh
go install github.com/gocombinator/git-semver@latest
```

## Examples

We're in repo that looks like this:
```
$ git branch
* master

$ git tag
v0.1.0
v0.1.1
v0.2.0
v1.0.0
v2.0.0-foo
v3.0.0-foo+bar
v4.0.0
```

Print current highest semver:
```
$ git-semver
v4.0.0
```

Preview version bump:
```
$ git-semver -minor
v4.0.0 -> v4.1.0
```

Assert current branch name when trying to bump patch version:
```
$ git-semver -patch -branch main
git branch is master, not main

$ git-semver -patch -branch master
v4.0.0 -> v4.0.1
```

Create tag (using `-tag` flag):
```
$ git-semver -patch -tag -branch master
v4.0.0 -> v4.0.1

$ git tag
v0.1.0
v0.1.1
v0.2.0
v1.0.0
v2.0.0-foo
v3.0.0-foo+bar
v4.0.0
v4.0.1 # <- v4.0.1 has been created
```

Clean working tree is asserted:
```
$ touch bar # make working tree dirty
$ git-semver -patch -tag -push -branch master
git working tree not clean:
?? bar
```

Usually you'll just want to tag and push patch, minor or major bump while making sure
you're on the right branch (clean working tree will be asserted automatically) with:
```
$ git-semver -patch -tag -push -branch master
$ git-semver -minor -tag -push -branch master
$ git-semver -major -tag -push -branch master
```

## License

```
MIT License

Copyright 2022 Mirek Rusin

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
