# go-dep-test

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
# reason: only when built on Windows (via Go build tags)
github.com/inconshreveable/mousetrap v1.1.0

# reason: in package 'doc' (part of module 'cobra' but not used)
github.com/cpuguy83/go-md2man/v2
|-- github.com/russross/blackfriday/v2
gopkg.in/yaml.v3
|-- gopkg.in/check.v1

# reason: test dependencies of strftime
gubraun/app
|-- github.com/lestrrat-go/strftime
    |-- github.com/lestrrat-go/strftime.test
        |-- github.com/stretchr/testify/assert
        |-- github.com/davecgh/go-spew/spew
        |-- github.com/lestrrat-go/envload
        |-- github.com/stretchr/testify/assert
            |-- github.com/pmezard/go-difflib/difflib

# reason: in package 'mock' (part of module 'testify' but not used)
github.com/stretchr/objx
```
