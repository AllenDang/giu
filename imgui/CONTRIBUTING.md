## Contributing Guidelines

Thank you for considering to contribute to this library!

The following text lists guidelines for contributions.
These guidelines don't have legal status, so use them as a reference and common sense - and feel free to update them as well!


### "I just want to know..."

For questions, or general usage-information about **Dear ImGui**, please refer to the [homepage](https://github.com/ocornut/imgui), or look in the detailed documentation of the C++ source.
This wrapper houses next to no functionality. As long as it is not a Go-specific issue, help will rather be there.

### Scope

This wrapper exposes minimal functionality of **Dear ImGui**. Ideally, this functionality is that of the common minimum that someone would want. This wrapper does not strife for full configurability, like the original library. This is not even possible in some cases, as it requires compilation flags.

### Extensions
At the moment, this library is primarily used by **InkyBlackness**. If you can and want to make use of this library in your own projects, you are happy to do so. Pull-requests with extensions are happily accepted, provided that they uphold the following minimum requirements:
* Code is properly formatted & linted (use [golangci-lint](https://github.com/golangci/golangci-lint) for a full check)
* Public Go API is documented. Copied documentation from **Dear ImGui** is acceptable and recommended, assuming it is adapted regarding type names. If there is no documentation in the original, try to spend some time figuring it out. In any case, please make the comments readable as complete English sentences, as recommended by Go.
* API and version philosophies are respected (see README.md)

#### Clarification on API naming and signatures

If an **Dear ImGui** function has the signature of

```
SomeControl(const char *label, int value, int optArg1 = 0, const char *optArg2 = "stuff");
```

then the wrapper functions should be

```
// SomeControl calls SomeControlV(label, value, 0, "stuff"). 
SomeControl(label string, value int32) {
    SomeControlV(label, value, 0, "stuff")
}

// SomeControlV does things (text possibly copied from imgui.h).
SomeControlV(label string, value int32, optArg1 int32, optArg2 string) {
    // ...
}
```

The "idiomatic" function should have only the required parameters of the underlying function, and its comment specifies all the defaults, matching that of `imgui.h`.
The "verbose" variant should require all the parameters of the underlying function.

### Code Style

Please make sure code is formatted according to `go fmt`, and use the following linter: [golangci-lint](https://github.com/golangci/golangci-lint).

> If there are linter errors that you didn't introduce, you don't have to clean them up - I might have missed them and will be handling them separately.

### Upgrade to newer Dear ImGui version

An upgrade with _major_ changes in the API should be on purpose and with coordination. Such a change requires a bump of the major version of this wrapper.

Otherwise, try to keep the API of this wrapper stable and keep compatible wrapper functions for changed/upgraded functions.
  
On an upgrade of **Dear ImGui**, apart from updating the actual files, be sure to do the following steps:
* In case `imconfig.h` is changed, be sure to keep the intentional changes: Obsolete functions should not be compiled, and the `iggAssert()` function must survive.
* Have a look at any extended enumerations. The Go variant will need extension/change as well, otherwise the constants will be wrong.
* Check for any documentation changes of exported functions. The Go documentation should reflect such changes as well.
* Run `go test ./...` . There is at least one test that is bound to the version and needs change as well.
* Update the screenshots of the examples, they show the version number.
* Update the `README.md` file, it indicates the version number.
* Check if the license of **Dear ImGui** has changed and update the `_licenses/imgui-LICENSE.txt` file. This may happen every year (copyright year).

#### Handling of removed functions

In order to avoid needing a major version bump of the wrapper just for one removed function, use the following pattern:

```
// NewFunction does new stuff.
func NewFunction() {
}

// OldFunction did something and is now delegating to NewFunction().
// Deprecated: Use NewFunction instead.  
func OldFunction() {
    NewFunction()
}
```

The `OldFunction` is implemented using the new API, and is marked as `Deprecated` in the comment.
IDEs tend to respect this and notify the user.

If, however, a whole set of functionality is replaced, this then probably warrants a major version bump.
