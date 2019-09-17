github.com/teamlint/iris
  ├ bytes
  ├ context
  ├ encoding/json
  ├ fmt
  ├ io
  ├ io/ioutil
  ├ log
  ├ net
  ├ net/http
  ├ os
  ├ os/exec
  ├ os/user
  ├ path/filepath
  ├ runtime
  ├ strings
  ├ sync
  ├ time
  ├ github.com/teamlint/iris/cache
  │ ├ strconv
  │ ├ time
  │ ├ github.com/teamlint/iris/cache/client
  │ │ ├ bytes
  │ │ ├ io/ioutil
  │ │ ├ net/http
  │ │ ├ sync
  │ │ ├ time
  │ │ ├ github.com/teamlint/iris/cache/cfg
  │ │ │ └ time
  │ │ ├ github.com/teamlint/iris/cache/client/rule
  │ │ │ ├ github.com/teamlint/iris/cache/ruleset
  │ │ │ └ github.com/teamlint/iris/context
  │ │ │   ├ bufio
  │ │ │   ├ bytes
  │ │ │   ├ encoding/json
  │ │ │   ├ encoding/xml
  │ │ │   ├ fmt
  │ │ │   ├ io
  │ │ │   ├ io/ioutil
  │ │ │   ├ math
  │ │ │   ├ mime
  │ │ │   ├ mime/multipart
  │ │ │   ├ net
  │ │ │   ├ net/http
  │ │ │   ├ net/url
  │ │ │   ├ os
  │ │ │   ├ path
  │ │ │   ├ path/filepath
  │ │ │   ├ reflect
  │ │ │   ├ regexp
  │ │ │   ├ runtime
  │ │ │   ├ strconv
  │ │ │   ├ strings
  │ │ │   ├ sync
  │ │ │   ├ sync/atomic
  │ │ │   ├ time
  │ │ │   ├ github.com/teamlint/iris/core/errors
  │ │ │   │ ├ fmt
  │ │ │   │ ├ runtime
  │ │ │   │ ├ strings
  │ │ │   │ ├ sync
  │ │ │   │ └ github.com/teamlint/iris/vendor/github.com/iris-contrib/go.uuid
  │ │ │   │   ├ bytes
  │ │ │   │   ├ crypto/md5
  │ │ │   │   ├ crypto/rand
  │ │ │   │   ├ crypto/sha1
  │ │ │   │   ├ database/sql/driver
  │ │ │   │   ├ encoding/binary
  │ │ │   │   ├ encoding/hex
  │ │ │   │   ├ fmt
  │ │ │   │   ├ hash
  │ │ │   │   ├ io
  │ │ │   │   ├ net
  │ │ │   │   ├ os
  │ │ │   │   ├ sync
  │ │ │   │   └ time
  │ │ │   ├ github.com/teamlint/iris/core/memstore
  │ │ │   │ ├ bytes
  │ │ │   │ ├ encoding/gob
  │ │ │   │ ├ fmt
  │ │ │   │ ├ io
  │ │ │   │ ├ math
  │ │ │   │ ├ reflect
  │ │ │   │ ├ strconv
  │ │ │   │ ├ strings
  │ │ │   │ ├ time
  │ │ │   │ └ github.com/teamlint/iris/core/errors
  │ │ │   ├ github.com/teamlint/iris/macro
  │ │ │   │ ├ fmt
  │ │ │   │ ├ reflect
  │ │ │   │ ├ regexp
  │ │ │   │ ├ strconv
  │ │ │   │ ├ strings
  │ │ │   │ ├ unicode
  │ │ │   │ ├ github.com/teamlint/iris/core/memstore
  │ │ │   │ ├ github.com/teamlint/iris/macro/interpreter/ast
  │ │ │   │ └ github.com/teamlint/iris/macro/interpreter/parser
  │ │ │   │   ├ fmt
  │ │ │   │   ├ strconv
  │ │ │   │   ├ strings
  │ │ │   │   ├ github.com/teamlint/iris/macro/interpreter/ast
  │ │ │   │   ├ github.com/teamlint/iris/macro/interpreter/lexer
  │ │ │   │   │ └ github.com/teamlint/iris/macro/interpreter/token
  │ │ │   │   └ github.com/teamlint/iris/macro/interpreter/token
  │ │ │   ├ github.com/teamlint/iris/vendor/github.com/Shopify/goreferrer
  │ │ │   │ ├ encoding/json
  │ │ │   │ ├ io
  │ │ │   │ ├ net/url
  │ │ │   │ ├ path
  │ │ │   │ ├ strings
  │ │ │   │ └ github.com/teamlint/iris/vendor/golang.org/x/net/publicsuffix
  │ │ │   │   ├ fmt
  │ │ │   │   ├ net/http/cookiejar
  │ │ │   │   └ strings
  │ │ │   ├ github.com/teamlint/iris/vendor/github.com/fatih/structs
  │ │ │   │ ├ errors
  │ │ │   │ ├ fmt
  │ │ │   │ ├ reflect
  │ │ │   │ └ strings
  │ │ │   ├ github.com/teamlint/iris/vendor/github.com/iris-contrib/blackfriday
  │ │ │   │ ├ bytes
  │ │ │   │ ├ fmt
  │ │ │   │ ├ html
  │ │ │   │ ├ io
  │ │ │   │ ├ regexp
  │ │ │   │ ├ strconv
  │ │ │   │ ├ strings
  │ │ │   │ ├ unicode/utf8
  │ │ │   │ └ github.com/teamlint/iris/vendor/github.com/shurcooL/sanitized_anchor_name
  │ │ │   │   └ unicode
  │ │ │   ├ github.com/teamlint/iris/vendor/github.com/iris-contrib/schema
  │ │ │   │ ├ encoding
  │ │ │   │ ├ errors
  │ │ │   │ ├ fmt
  │ │ │   │ ├ reflect
  │ │ │   │ ├ strconv
  │ │ │   │ ├ strings
  │ │ │   │ └ sync
  │ │ │   ├ github.com/teamlint/iris/vendor/github.com/json-iterator/go
  │ │ │   │ ├ bytes
  │ │ │   │ ├ encoding
  │ │ │   │ ├ encoding/base64
  │ │ │   │ ├ encoding/json
  │ │ │   │ ├ errors
  │ │ │   │ ├ fmt
  │ │ │   │ ├ io
  │ │ │   │ ├ math
  │ │ │   │ ├ math/big
  │ │ │   │ ├ reflect
  │ │ │   │ ├ sort
  │ │ │   │ ├ strconv
  │ │ │   │ ├ strings
  │ │ │   │ ├ sync
  │ │ │   │ ├ unicode
  │ │ │   │ ├ unicode/utf16
  │ │ │   │ ├ unicode/utf8
  │ │ │   │ ├ unsafe
  │ │ │   │ ├ github.com/teamlint/iris/vendor/github.com/modern-go/concurrent
  │ │ │   │ │ ├ context
  │ │ │   │ │ ├ fmt
  │ │ │   │ │ ├ io/ioutil
  │ │ │   │ │ ├ log
  │ │ │   │ │ ├ os
  │ │ │   │ │ ├ reflect
  │ │ │   │ │ ├ runtime
  │ │ │   │ │ ├ runtime/debug
  │ │ │   │ │ ├ sync
  │ │ │   │ │ └ time
  │ │ │   │ └ github.com/teamlint/iris/vendor/github.com/modern-go/reflect2
  │ │ │   │   ├ reflect
  │ │ │   │   ├ runtime
  │ │ │   │   ├ strings
  │ │ │   │   ├ sync
  │ │ │   │   ├ unsafe
  │ │ │   │   └ github.com/teamlint/iris/vendor/github.com/modern-go/concurrent
  │ │ │   ├ github.com/teamlint/iris/vendor/github.com/klauspost/compress/gzip
  │ │ │   │ ├ bufio
  │ │ │   │ ├ encoding/binary
  │ │ │   │ ├ errors
  │ │ │   │ ├ fmt
  │ │ │   │ ├ hash/crc32
  │ │ │   │ ├ io
  │ │ │   │ ├ time
  │ │ │   │ └ github.com/teamlint/iris/vendor/github.com/klauspost/compress/flate
  │ │ │   │   ├ bufio
  │ │ │   │   ├ fmt
  │ │ │   │   ├ io
  │ │ │   │   ├ math
  │ │ │   │   ├ math/bits
  │ │ │   │   ├ sort
  │ │ │   │   ├ strconv
  │ │ │   │   ├ sync
  │ │ │   │   └ github.com/teamlint/iris/vendor/github.com/klauspost/cpuid
  │ │ │   │     └ strings
  │ │ │   ├ github.com/teamlint/iris/vendor/github.com/microcosm-cc/bluemonday
  │ │ │   │ ├ bytes
  │ │ │   │ ├ encoding/base64
  │ │ │   │ ├ io
  │ │ │   │ ├ net/url
  │ │ │   │ ├ regexp
  │ │ │   │ ├ strings
  │ │ │   │ └ github.com/teamlint/iris/vendor/golang.org/x/net/html
  │ │ │   │   ├ bufio
  │ │ │   │   ├ bytes
  │ │ │   │   ├ errors
  │ │ │   │   ├ fmt
  │ │ │   │   ├ io
  │ │ │   │   ├ strconv
  │ │ │   │   ├ strings
  │ │ │   │   ├ unicode/utf8
  │ │ │   │   └ github.com/teamlint/iris/vendor/golang.org/x/net/html/atom
  │ │ │   ├ github.com/teamlint/iris/vendor/github.com/teamlint/golog
  │ │ │   └ github.com/teamlint/iris/vendor/gopkg.in/yaml.v2
  │ │ │     ├ bytes
  │ │ │     ├ encoding
  │ │ │     ├ encoding/base64
  │ │ │     ├ errors
  │ │ │     ├ fmt
  │ │ │     ├ io
  │ │ │     ├ math
  │ │ │     ├ reflect
  │ │ │     ├ regexp
  │ │ │     ├ sort
  │ │ │     ├ strconv
  │ │ │     ├ strings
  │ │ │     ├ sync
  │ │ │     ├ time
  │ │ │     ├ unicode
  │ │ │     └ unicode/utf8
  │ │ ├ github.com/teamlint/iris/cache/entry
  │ │ │ ├ net/http
  │ │ │ ├ regexp
  │ │ │ ├ strconv
  │ │ │ ├ time
  │ │ │ └ github.com/teamlint/iris/cache/cfg
  │ │ ├ github.com/teamlint/iris/cache/ruleset
  │ │ ├ github.com/teamlint/iris/cache/uri
  │ │ │ ├ net/url
  │ │ │ ├ strconv
  │ │ │ ├ strings
  │ │ │ ├ time
  │ │ │ └ github.com/teamlint/iris/cache/cfg
  │ │ └ github.com/teamlint/iris/context
  │ └ github.com/teamlint/iris/context
  ├ github.com/teamlint/iris/context
  ├ github.com/teamlint/iris/core/errors
  ├ github.com/teamlint/iris/core/handlerconv
  │ ├ net/http
  │ ├ github.com/teamlint/iris/context
  │ └ github.com/teamlint/iris/core/errors
  ├ github.com/teamlint/iris/core/host
  │ ├ context
  │ ├ crypto/tls
  │ ├ fmt
  │ ├ io
  │ ├ net
  │ ├ net/http
  │ ├ net/http/httputil
  │ ├ net/url
  │ ├ os
  │ ├ os/signal
  │ ├ runtime
  │ ├ strings
  │ ├ sync
  │ ├ sync/atomic
  │ ├ syscall
  │ ├ time
  │ ├ github.com/teamlint/iris/core/errors
  │ ├ github.com/teamlint/iris/core/netutil
  │ │ ├ crypto/tls
  │ │ ├ net
  │ │ ├ net/http
  │ │ ├ os
  │ │ ├ regexp
  │ │ ├ strconv
  │ │ ├ strings
  │ │ ├ time
  │ │ ├ github.com/teamlint/iris/core/errors
  │ │ ├ github.com/teamlint/iris/vendor/github.com/teamlint/golog
  │ │ └ github.com/teamlint/iris/vendor/golang.org/x/crypto/acme/autocert
  │ │   ├ bytes
  │ │   ├ context
  │ │   ├ crypto
  │ │   ├ crypto/ecdsa
  │ │   ├ crypto/elliptic
  │ │   ├ crypto/rand
  │ │   ├ crypto/rsa
  │ │   ├ crypto/tls
  │ │   ├ crypto/x509
  │ │   ├ crypto/x509/pkix
  │ │   ├ encoding/pem
  │ │   ├ errors
  │ │   ├ fmt
  │ │   ├ io
  │ │   ├ io/ioutil
  │ │   ├ log
  │ │   ├ math/rand
  │ │   ├ net
  │ │   ├ net/http
  │ │   ├ os
  │ │   ├ path
  │ │   ├ path/filepath
  │ │   ├ runtime
  │ │   ├ strings
  │ │   ├ sync
  │ │   ├ time
  │ │   ├ github.com/teamlint/iris/vendor/golang.org/x/crypto/acme
  │ │   │ ├ bytes
  │ │   │ ├ context
  │ │   │ ├ crypto
  │ │   │ ├ crypto/ecdsa
  │ │   │ ├ crypto/elliptic
  │ │   │ ├ crypto/rand
  │ │   │ ├ crypto/rsa
  │ │   │ ├ crypto/sha256
  │ │   │ ├ crypto/sha512
  │ │   │ ├ crypto/tls
  │ │   │ ├ crypto/x509
  │ │   │ ├ crypto/x509/pkix
  │ │   │ ├ encoding/asn1
  │ │   │ ├ encoding/base64
  │ │   │ ├ encoding/hex
  │ │   │ ├ encoding/json
  │ │   │ ├ encoding/pem
  │ │   │ ├ errors
  │ │   │ ├ fmt
  │ │   │ ├ io
  │ │   │ ├ io/ioutil
  │ │   │ ├ math/big
  │ │   │ ├ net/http
  │ │   │ ├ runtime/debug
  │ │   │ ├ strconv
  │ │   │ ├ strings
  │ │   │ ├ sync
  │ │   │ └ time
  │ │   └ github.com/teamlint/iris/vendor/golang.org/x/net/idna
  │ │     ├ fmt
  │ │     ├ math
  │ │     ├ strings
  │ │     ├ unicode/utf8
  │ │     ├ github.com/teamlint/iris/vendor/golang.org/x/text/secure/bidirule
  │ │     │ ├ errors
  │ │     │ ├ unicode/utf8
  │ │     │ ├ github.com/teamlint/iris/vendor/golang.org/x/text/transform
  │ │     │ │ ├ bytes
  │ │     │ │ ├ errors
  │ │     │ │ ├ io
  │ │     │ │ └ unicode/utf8
  │ │     │ └ github.com/teamlint/iris/vendor/golang.org/x/text/unicode/bidi
  │ │     │   ├ container/list
  │ │     │   ├ fmt
  │ │     │   ├ log
  │ │     │   ├ sort
  │ │     │   └ unicode/utf8
  │ │     ├ github.com/teamlint/iris/vendor/golang.org/x/text/unicode/bidi
  │ │     └ github.com/teamlint/iris/vendor/golang.org/x/text/unicode/norm
  │ │       ├ encoding/binary
  │ │       ├ fmt
  │ │       ├ io
  │ │       ├ sync
  │ │       ├ unicode/utf8
  │ │       └ github.com/teamlint/iris/vendor/golang.org/x/text/transform
  │ └ github.com/teamlint/iris/vendor/golang.org/x/crypto/acme/autocert
  ├ github.com/teamlint/iris/core/netutil
  ├ github.com/teamlint/iris/core/router
  │ ├ bytes
  │ ├ fmt
  │ ├ io
  │ ├ io/ioutil
  │ ├ mime
  │ ├ net/http
  │ ├ net/url
  │ ├ os
  │ ├ path
  │ ├ path/filepath
  │ ├ sort
  │ ├ strconv
  │ ├ strings
  │ ├ sync
  │ ├ time
  │ ├ github.com/teamlint/iris/context
  │ ├ github.com/teamlint/iris/core/errors
  │ ├ github.com/teamlint/iris/core/netutil
  │ ├ github.com/teamlint/iris/macro
  │ ├ github.com/teamlint/iris/macro/handler
  │ │ ├ github.com/teamlint/iris/context
  │ │ └ github.com/teamlint/iris/macro
  │ ├ github.com/teamlint/iris/macro/interpreter/ast
  │ ├ github.com/teamlint/iris/macro/interpreter/lexer
  │ └ github.com/teamlint/iris/vendor/github.com/teamlint/golog
  ├ github.com/teamlint/iris/middleware/logger
  │ ├ fmt
  │ ├ strconv
  │ ├ time
  │ ├ github.com/teamlint/iris/context
  │ └ github.com/teamlint/iris/vendor/github.com/ryanuber/columnize
  │   ├ bytes
  │   ├ fmt
  │   └ strings
  ├ github.com/teamlint/iris/middleware/recover
  │ ├ fmt
  │ ├ runtime
  │ ├ strconv
  │ └ github.com/teamlint/iris/context
  ├ github.com/teamlint/iris/vendor/github.com/BurntSushi/toml
  │ ├ bufio
  │ ├ encoding
  │ ├ errors
  │ ├ fmt
  │ ├ io
  │ ├ io/ioutil
  │ ├ math
  │ ├ reflect
  │ ├ sort
  │ ├ strconv
  │ ├ strings
  │ ├ sync
  │ ├ time
  │ ├ unicode
  │ └ unicode/utf8
  ├ github.com/teamlint/iris/vendor/github.com/teamlint/golog
  │ ├ fmt
  │ ├ io
  │ ├ os
  │ ├ strings
  │ ├ sync
  │ ├ time
  │ └ github.com/teamlint/iris/vendor/github.com/teamlint/pio
  │   ├ bufio
  │   ├ bytes
  │   ├ encoding/json
  │   ├ encoding/xml
  │   ├ errors
  │   ├ fmt
  │   ├ io
  │   ├ io/ioutil
  │   ├ os/exec
  │   ├ runtime
  │   ├ sort
  │   ├ strconv
  │   ├ strings
  │   ├ sync
  │   ├ sync/atomic
  │   └ github.com/teamlint/iris/vendor/github.com/teamlint/pio/terminal
  │     ├ io
  │     ├ os
  │     ├ syscall
  │     └ unsafe
  ├ github.com/teamlint/iris/vendor/gopkg.in/yaml.v2
  └ github.com/teamlint/iris/view
    ├ bytes
    ├ fmt
    ├ html/template
    ├ io
    ├ io/ioutil
    ├ os
    ├ path
    ├ path/filepath
    ├ reflect
    ├ strings
    ├ sync
    ├ time
    ├ github.com/teamlint/iris/context
    ├ github.com/teamlint/iris/core/errors
    └ github.com/teamlint/iris/vendor/github.com/CloudyKit/jet
      ├ bytes
      ├ encoding/json
      ├ errors
      ├ fmt
      ├ html
      ├ io
      ├ io/ioutil
      ├ net/url
      ├ os
      ├ path
      ├ path/filepath
      ├ reflect
      ├ runtime
      ├ strconv
      ├ strings
      ├ sync
      ├ text/template
      ├ unicode
      ├ unicode/utf8
      └ github.com/teamlint/iris/vendor/github.com/CloudyKit/fastprinter
        ├ fmt
        ├ io
        ├ math
        ├ reflect
        └ sync
120 dependencies (68 internal, 52 external, 0 testing).
