# Hacking

## Release

- Update the version in _version.go_
- Commit, tag and push the changes
- Create a new release on GitHub: https://github.com/yuzutech/kroki-go/releases/new
- Prepare the next version, edit _version.go_ and update the version number with `-SNAPSHOT` suffix
- Commit and push the changes

For instance, if we want to release version 2.1.3:

**version.go**
```go
const Version = "2.1.3"
```

    $ git commit -m "Version 2.1.3"
    $ git tag v2.1.3
    $ git push origin master --tags

Then prepare the next version 2.2.0:

**version.go**
```go
const Version = "2.2.0-SNAPSHOT"
```

    $ git commit -m "Prepare next version 2.2.0"
    $ git push origin master
