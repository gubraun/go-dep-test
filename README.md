# go-dep-test

## Reasons
1. Dependencies of unused package of dependent module
   * Example: dependent module `cobra` has packages `cobra` and `doc`, but only package `cobra` is imported. Thus, all dependencies of package `doc` are not required.
2. Transitive test dependencies
   * Example: module `strftime` has a package `strftime.test`, which has dependencies. Similar to reason 1 above.
3. Direct test dependency
   * Example: `main_test.go` has dependencies.
4. Dependencies specific to build tags
   * Example: Window-specific dependencies.

## Actual Dependencies
```
github.com/gubraun/go-dep-test
|-- github.com/lestrrat-go/strftime   v1.0.5
    |-- github.com/pkg/errors         v0.9.1
|-- github.com/spf13/cobra            v1.7.0
    |-- github.com/spf13/pflag        v1.0.5
```
 
## Unused Dependencies
```
# reason (4): only when built on Windows (via Go build tags)
github.com/inconshreveable/mousetrap v1.1.0

# reason (1): in package 'doc' (part of module 'cobra' but not used)
github.com/cpuguy83/go-md2man/v2
|-- github.com/russross/blackfriday/v2
gopkg.in/yaml.v3
|-- gopkg.in/check.v1

# reason (2): test dependencies of strftime
gubraun/app
|-- github.com/lestrrat-go/strftime
    |-- github.com/lestrrat-go/strftime.test
        |-- github.com/stretchr/testify/assert
        |-- github.com/davecgh/go-spew/spew
        |-- github.com/lestrrat-go/envload
        |-- github.com/stretchr/testify/assert
            |-- github.com/pmezard/go-difflib/difflib

# reason (1,2): in package 'mock' (part of module 'testify' but not used)
github.com/stretchr/objx

# reason (3): used by 'main_test.go' only
gitlab.com/nyarla/go-crypt
```
