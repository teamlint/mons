<!-- # History/Changelog <a href="HISTORY_ZH.md"> <img width="20px" src="https://iris-go.com/images/flag-china.svg?v=10" /></a><a href="HISTORY_ID.md"> <img width="20px" src="https://iris-go.com/images/flag-indonesia.svg?v=10" /></a><a href="HISTORY_GR.md"> <img width="20px" src="https://iris-go.com/images/flag-greece.svg?v=10" /></a> -->

# Changelog

### Looking for free and real-time support?

    https://github.com/teamlint/iris/issues
    https://chat.iris-go.com

### Looking for previous versions?

    https://github.com/teamlint/iris/releases

<<<<<<< HEAD
### Should I upgrade my Iris?

Developers are not forced to upgrade if they don't really need it. Upgrade whenever you feel ready.

> Iris uses the [vendor directory](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo) feature, so you get truly reproducible builds, as this method guards against upstream renames and deletes.

**How to upgrade**: Open your command-line and execute this command: `go get -u github.com/teamlint/iris` or let the automatic updater do that for you.

# Fr, 11 January 2019 | v11.1.1

Happy new year! This is a minor release, contains mostly bug fixes.

Strange that we don't have major features in this release, right? Don't worry, I am not out of ideas (at least not yet!).
I have some features in-mind but lately I do not have the time to humanize those ideas for you due to my new position in [Netdata Inc.](https://github.com/netdata/netdata), so be patient and [stay-tuned](https://github.com/kataras/iris/stargazers). Read the current changelog below:

- session/redis: fix unused service config var. IdleTimeout witch was replaced by default values. [#1140](https://github.com/kataras/iris/pull/1140) ([@d7561985](https://github.com/d7561985))

- fix [#1141](https://github.com/kataras/iris/issues/1141) and [#1142](https://github.com/kataras/iris/issues/1142). [2bd7a8e88777766d1f4cac7562feec304112d2b1](https://github.com/kataras/iris/commit/2bd7a8e88777766d1f4cac7562feec304112d2b1) (@kataras)

- fix cache corruption due to recorder reuse. [#1146](https://github.com/kataras/iris/pull/1146) ([@Slamper](https://github.com/Slamper))

- add `StatusTooEarly`, compatible with: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/425#Browser_compatibility. [31b2913447aa9e41e16a3eb33eb0019427e15cea](https://github.com/kataras/iris/commit/31b2913447aa9e41e16a3eb33eb0019427e15cea) (@kataras)

- fix [#1164](https://github.com/kataras/iris/issues/1164). [701e8e46c20395f87fa34bf9fabd145074c7b78c](https://github.com/kataras/iris/commit/701e8e46c20395f87fa34bf9fabd145074c7b78c) (@kataras)

- `context#ReadForm` can skip unkown fields by `IsErrPath(err)`, fixes: [#1157](https://github.com/kataras/iris/issues/1157). [1607bb5113568af6a34142f23bfa44903205b314](https://github.com/kataras/iris/commit/1607bb5113568af6a34142f23bfa44903205b314) (@kataras)


Doc updates:

- fix grammar and misspell. [5069e9afd8700d20dfd04cdc008efd671b5d0b40](https://github.com/kataras/iris/commit/5069e9afd8700d20dfd04cdc008efd671b5d0b40) (@kataras)

- fix link for httpexpect in README. [#1148](https://github.com/kataras/iris/pull/1148) ([@drenel18](https://github.com/drenel18))

- translate _examples/README.md into Chinese. [#1156](https://github.com/kataras/iris/pull/1156) ([@fduxiao](https://github.com/fduxiao))

- add https://github.com/snowlyg/IrisApiProject to starter kits (Chinese). [ea12533871253afc34e40e36ba658b51955ea82d](https://github.com/kataras/iris/commit/ea12533871253afc34e40e36ba658b51955ea82d)

- add https://github.com/yz124/superstar to starter kits (Chinese). [0e734ff8445f07482c28881347c1e564dc5aab9c](https://github.com/kataras/iris/commit/0e734ff8445f07482c28881347c1e564dc5aab9c)

# Su, 18 November 2018 | v11.1.0

PR: https://github.com/kataras/iris/pull/1130

This release contains a new feature for versioning your Iris APIs. The initial motivation and feature request came by https://github.com/kataras/iris/issues/1129.

The [versioning](https://github.com/kataras/iris/tree/master/versioning) package provides [semver](https://semver.org/) versioning for your APIs. It implements all the suggestions written at [api-guidelines](https://github.com/byrondover/api-guidelines/blob/master/Guidelines.md#versioning) and more.


The version comparison is done by the [go-version](https://github.com/hashicorp/go-version) package. It supports matching over patterns like `">= 1.0, < 3"` and etc.

## Features

- per route version matching, a normal iris handler with "switch" cases via Map for version => handler
- per group versioned routes and deprecation API
- version matching like ">= 1.0, < 2.0" or just "2.0.1" and etc.
- version not found handler (can be customized by simply adding the versioning.NotFound: customNotMatchVersionHandler on the Map)
- version is retrieved from the "Accept" and "Accept-Version" headers (can be customized via middleware)
- respond with "X-API-Version" header, if version found.
- deprecation options with customizable "X-API-Warn", "X-API-Deprecation-Date", "X-API-Deprecation-Info" headers via `Deprecated` wrapper.

## Get version

Current request version is retrieved by `versioning.GetVersion(ctx)`.

By default the `GetVersion` will try to read from:
- `Accept` header, i.e `Accept: "application/json; version=1.0"`
- `Accept-Version` header, i.e `Accept-Version: "1.0"`

You can also set a custom version for a handler via a middleware by using the context's store values.
For example:
```go
func(ctx iris.Context) {
    ctx.Values().Set(versioning.Key, ctx.URLParamDefault("version", "1.0"))
    ctx.Next()
}
```

## Match version to handler

The `versioning.NewMatcher(versioning.Map) iris.Handler` creates a single handler which decides what handler need to be executed based on the requested version.

```go
app := iris.New()

// middleware for all versions.
myMiddleware := func(ctx iris.Context) {
    // [...]
    ctx.Next()
}

myCustomNotVersionFound := func(ctx iris.Context) {
    ctx.StatusCode(404)
    ctx.Writef("%s version not found", versioning.GetVersion(ctx))
}

userAPI := app.Party("/api/user")
userAPI.Get("/", myMiddleware, versioning.NewMatcher(versioning.Map{
    "1.0":               sendHandler(v10Response),
    ">= 2, < 3":         sendHandler(v2Response),
    versioning.NotFound: myCustomNotVersionFound,
}))
```

### Deprecation

Using the `versioning.Deprecated(handler iris.Handler, options versioning.DeprecationOptions) iris.Handler` function you can mark a specific handler version as deprecated.
=======
### Want to be hired?
>>>>>>> upstream/master

    https://facebook.com/iris.framework

### Should I upgrade my Iris?

Developers are not forced to upgrade if they don't really need it. Upgrade whenever you feel ready.

**How to upgrade**: Open your command-line and execute this command: `go get github.com/kataras/iris@master`.

# Fr, 16 August 2019 | v11.2.8

- Set `Cookie.SameSite` to `Lax` when subdomains sessions share is enabled[*](https://github.com/kataras/iris/commit/6bbdd3db9139f9038641ce6f00f7b4bab6e62550)
- Add and update all [experimental handlers](https://github.com/kataras/iris/tree/master/_examples/experimental-handlers) 
- New `XMLMap` function which wraps a `map[string]interface{}` and converts it to a valid xml content to render through `Context.XML` method
- Add new `ProblemOptions.XML` and `RenderXML` fields to render the `Problem` as XML(application/problem+xml) instead of JSON("application/problem+json) and enrich the `Negotiate` to easily accept the `application/problem+xml` mime.

Commit log: https://github.com/kataras/iris/compare/v11.2.7...v11.2.8

# Th, 15 August 2019 | v11.2.7

This minor version contains improvements on the Problem Details for HTTP APIs implemented on [v11.2.5](#mo-12-august-2019--v1125).

- Fix https://github.com/kataras/iris/issues/1335#issuecomment-521319721
- Add `ProblemOptions` with `RetryAfter` as requested at: https://github.com/kataras/iris/issues/1335#issuecomment-521330994.
- Add `iris.JSON` alias for `context#JSON` options type.

[Example](https://github.com/kataras/iris/blob/45d7c6fedb5adaef22b9730592255f7bb375e809/_examples/routing/http-errors/main.go#L85) and [wikis](https://github.com/kataras/iris/wiki/Routing-error-handlers#the-problem-type) updated. 

References:

- https://tools.ietf.org/html/rfc7231#section-7.1.3
- https://tools.ietf.org/html/rfc7807

Commit log: https://github.com/kataras/iris/compare/v11.2.6...v11.2.7

# We, 14 August 2019 | v11.2.6

Allow [handle more than one route with the same paths and parameter types but different macro validation functions](https://github.com/kataras/iris/issues/1058#issuecomment-521110639).

```go
app.Get("/{alias:string regexp(^[a-z0-9]{1,10}\\.xml$)}", PanoXML)
app.Get("/{alias:string regexp(^[a-z0-9]{1,10}$)}", Tour)
```

Commit log: https://github.com/kataras/iris/compare/v11.2.5...v11.2.6

# Mo, 12 August 2019 | v11.2.5

- [New Feature: Problem Details for HTTP APIs](https://github.com/kataras/iris/pull/1336)
- [Add Context.AbsoluteURI](https://github.com/kataras/iris/pull/1336/files#diff-15cce7299aae8810bcab9b0bf9a2fdb1R2368)

Commit log: https://github.com/kataras/iris/compare/v11.2.4...v11.2.5

# Fr, 09 August 2019 | v11.2.4

- Fixes [iris.Jet: no view engine found for '.jet' or '.html'](https://github.com/kataras/iris/issues/1327)
- Fixes [ctx.ViewData not work with JetEngine](https://github.com/kataras/iris/issues/1330)
- **New Feature**: [HTTP Method Override](https://github.com/kataras/iris/issues/1325)
- Fixes [Poor performance of session.UpdateExpiration on 200 thousands+ keys with new radix lib](https://github.com/kataras/iris/issues/1328) by introducing the `sessions.Config.Driver` configuration field which defaults to `Redigo()` but can be set to `Radix()` too, future additions are welcomed.

Commit log: https://github.com/kataras/iris/compare/v11.2.3...v11.2.4

# Tu, 30 July 2019 | v11.2.3

- [New Feature: Handle different parameter types in the same path](https://github.com/kataras/iris/issues/1315)
- [New Feature: Content Negotiation](https://github.com/kataras/iris/issues/1319)
- [Context.ReadYAML](https://github.com/kataras/iris/tree/master/_examples/http_request/read-yaml)
- Fixes https://github.com/kataras/neffos/issues/1#issuecomment-515698536

# We, 24 July 2019 | v11.2.2

Sessions as middleware:

```go
import "github.com/kataras/iris/sessions"
// [...]

app := iris.New()
sess := sessions.New(sessions.Config{...})

app.Get("/path", func(ctx iris.Context){
    session := sessions.Get(ctx)
    // [work with session...]
})
```

- Add `Session.Len() int` to return the total number of stored values/entries.
- Make `Context.HTML` and `Context.Text` to accept an optional, variadic, `args ...interface{}` input arg(s) too.

## v11.1.1

<<<<<<< HEAD
# Mo, 15 January 2018 | v10.0.1

Not any serious problems were found to be resolved here but one, the first one which is important for devs that used the [cache](cache) package.

- fix a single one cache handler didn't work across multiple route handlers at the same time https://github.com/kataras/iris/pull/852, as reported at https://github.com/kataras/iris/issues/850
- merge PR https://github.com/kataras/iris/pull/862
- do not allow concurrent access to the `ExecuteWriter -> Load` when `view#Engine##Reload` was true, as requested at https://github.com/kataras/iris/issues/872
- badge for open-source projects powered by Iris, learn how to add that badge to your open-source project at [FAQ.md](FAQ.md) file
- upstream update for `golang/crypto` to apply the fix about the [tls-sni challenge disabled](https://letsencrypt.status.io/pages/incident/55957a99e800baa4470002da/5a55777ed9a9c1024c00b241) https://github.com/golang/crypto/commit/13931e22f9e72ea58bb73048bc752b48c6d4d4ac (**relative to iris.AutoTLS**)

## New Backers

1. https://opencollective.com/cetin-basoz

## New Translations

1. The Chinese README_ZH.md and HISTORY_ZH.md was translated by @Zeno-Code via https://github.com/kataras/iris/pull/858
2. New Russian README_RU.md translations by @merrydii via https://github.com/kataras/iris/pull/857
3. New Greek README_GR.md and HISTORY_GR.md translations via https://github.com/kataras/iris/commit/8c4e17c2a5433c36c148a51a945c4dc35fbe502a#diff-74b06c740d860f847e7b577ad58ddde0 and https://github.com/kataras/iris/commit/bb5a81c540b34eaf5c6c8e993f644a0e66a78fb8

## New Examples

1. [MVC - Register Middleware](_examples/mvc/middleware)

## New Articles

1. [A Todo MVC Application using Iris and Vue.js](https://hackernoon.com/a-todo-mvc-application-using-iris-and-vue-js-5019ff870064)
2. [A Hasura starter project with a ready to deploy Golang hello-world web app with IRIS](bit.ly/2lmKaAZ)

# Mo, 01 January 2018 | v10.0.0

We must thanks [Mrs. Diana](https://www.instagram.com/merry.dii/) for our awesome new [logo](https://iris-go.com/images/icon.svg)!

You can [contact](mailto:Kovalenkodiana8@gmail.com) her for any  design-related enquiries or explore and send a direct message via [instagram](https://www.instagram.com/merry.dii/).

<p align="center">
<img width="145px" src="https://iris-go.com/images/icon.svg?v=a" />
</p>

At this version we have many internal improvements but just two major changes and one big feature, called **hero**.

> The new version adds 75 plus new commits, the PR is located [here](https://github.com/teamlint/iris/pull/849) read the internal changes if you are developing a web framework based on Iris. Why 9 was skipped? Because.

## Hero

The new package [hero](hero) contains features for binding any object or function that `handlers` may use, these are called dependencies. Hero funcs can also return any type of values, these values will be dispatched to the client.

> You may saw binding before but you didn't have code editor's support, with Iris you get truly safe binding thanks to the new `hero` package. It's also fast, near to raw handlers performance because Iris calculates everything before server ran!

Below you will see some screenshots we prepared for you in order to be easier to understand:

### 1. Path Parameters - Built'n Dependencies

![](https://github.com/teamlint/explore/raw/master/iris/hero/hero-1-monokai.png)

### 2. Services - Static Dependencies

![](https://github.com/teamlint/explore/raw/master/iris/hero/hero-2-monokai.png)

### 3. Per-Request - Dynamic Dependencies

![](https://github.com/teamlint/explore/raw/master/iris/hero/hero-3-monokai.png)

`hero funcs` are very easy to understand and when you start using them **you never go back**.

Examples:

- [Basic](_examples/hero/basic/main.go)
- [Overview](_examples/hero/overview)

## MVC

You have to understand the `hero` package in order to use the `mvc`, because `mvc` uses the `hero` internally for the controller's methods you use as routes, the same rules applied to those controller's methods of yours as well.

With this version you can register **any controller's methods as routes manually**, you can **get a route based on a method name and change its `Name` (useful for reverse routing inside templates)**, you can use any **dependencies** registered from `hero.Register` or `mvc.New(iris.Party).Register` per mvc application or per-controller, **you can still use `BeginRequest` and `EndRequest`**, you can catch **`BeforeActivation(b mvc.BeforeActivation)` to add dependencies per controller and `AfterActivation(a mvc.AfterActivation)` to make any post-validations**, **singleton controllers when no dynamic dependencies are used**, **Websocket controller, as simple as a `websocket.Connection` dependency** and more...

Examples:

**If you used MVC before then read very carefully: MVC CONTAINS SOME BREAKING CHANGES BUT YOU CAN DO A LOT MORE AND EVEN FASTER THAN BEFORE**

**PLEASE READ THE EXAMPLES CAREFULLY, WE'VE MADE THEM FOR YOU**

Old examples are here as well. Compare the two different versions of each example to understand what you win if you upgrade now.

| NEW | OLD |
| -----------|-------------|
| [Hello world](_examples/mvc/hello-world/main.go) | [OLD Hello world](https://github.com/teamlint/iris/blob/v8/_examples/mvc/hello-world/main.go) |
| [Session Controller](_examples/mvc/session-controller/main.go) | [OLD Session Controller](https://github.com/teamlint/iris/blob/v8/_examples/mvc/session-controller/main.go) |
| [Overview - Plus Repository and Service layers](_examples/mvc/overview) | [OLD Overview - Plus Repository and Service layers](https://github.com/teamlint/iris/tree/v8/_examples/mvc/overview) |
| [Login showcase - Plus Repository and Service layers](_examples/mvc/login) | [OLD Login showcase - Plus Repository and Service layers](https://github.com/teamlint/iris/tree/v8/_examples/mvc/login) |
| [Singleton](_examples/mvc/singleton) |  **NEW** |
| [Websocket Controller](_examples/mvc/websocket) |  **NEW** |
| [Vue.js Todo MVC](_examples/tutorial/vuejs-todo-mvc) |  **NEW** |

## context#PostMaxMemory

Remove the old static variable `context.DefaultMaxMemory` and replace it with the configuration `WithPostMaxMemory`.

```go
// WithPostMaxMemory sets the maximum post data size
// that a client can send to the server, this differs
// from the overral request body size which can be modified
// by the `context#SetMaxRequestBodySize` or `iris#LimitRequestBodySize`.
//
// Defaults to 32MB or 32 << 20 if you prefer.
func WithPostMaxMemory(limit int64) Configurator
```

If you used that old static field you will have to change that single line.

Usage:

```go
import "github.com/teamlint/iris"

func main() {
    app := iris.New()
    // [...]

    app.Run(iris.Addr(":8080"), iris.WithPostMaxMemory(10 << 20))
}
```

## context#UploadFormFiles

New method to upload multiple files, should be used for common upload actions, it's just a helper function.

```go
// UploadFormFiles uploads any received file(s) from the client
// to the system physical location "destDirectory".
//
// The second optional argument "before" gives caller the chance to
// modify the *miltipart.FileHeader before saving to the disk,
// it can be used to change a file's name based on the current request,
// all FileHeader's options can be changed. You can ignore it if
// you don't need to use this capability before saving a file to the disk.
//
// Note that it doesn't check if request body streamed.
//
// Returns the copied length as int64 and
// a not nil error if at least one new file
// can't be created due to the operating system's permissions or
// http.ErrMissingFile if no file received.
//
// If you want to receive & accept files and manage them manually you can use the `context#FormFile`
// instead and create a copy function that suits your needs, the below is for generic usage.
//
// The default form's memory maximum size is 32MB, it can be changed by the
//  `iris#WithPostMaxMemory` configurator at main configuration passed on `app.Run`'s second argument.
//
// See `FormFile` to a more controlled to receive a file.
func (ctx *context) UploadFormFiles(
        destDirectory string,
        before ...func(string, string),
    ) (int64, error)
```

Example can be found [here](_examples/http_request/upload-files/main.go).

## context#View

Just a minor addition, add a second optional variadic argument to the `context#View` method to accept a single value for template binding.
When you just want one value and not key-value pairs, you used to use an empty string on the `ViewData`, which is fine, especially if you preload these from a previous handler/middleware in the request handlers chain.

```go
func(ctx iris.Context) {
    ctx.ViewData("", myItem{Name: "iris" })
    ctx.View("item.html")
}
```

Same as:

```go
func(ctx iris.Context) {
    ctx.View("item.html", myItem{Name: "iris" })
}
```

```html
Item's name: {{.Name}}
```

## context#YAML

Add a new `context#YAML` function, it renders a yaml from a structured value.

```go
// YAML marshals the "v" using the yaml marshaler and renders its result to the client.
func YAML(v interface{}) (int, error)
```
=======
- https://github.com/kataras/iris/issues/1298
- https://github.com/kataras/iris/issues/1207
>>>>>>> upstream/master

# Tu, 23 July 2019 | v11.2.0

Read about the new release at: https://www.facebook.com/iris.framework/posts/3276606095684693
