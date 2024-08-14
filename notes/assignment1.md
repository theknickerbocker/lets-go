# 1. Introduction

We'll be making a web app called "Snippetbox", where users can store and display text snippets.
It will start simple, and each chapter will build upon it.

## 1.1. Prerequisites

The book is designed for people new to go, but in order to get the most out of it a general
understanding of the Go syntax is recommended. The [Go Tour](https://go.dev/tour/welcome/1) is
a good starting place, and the [Little Book of Go](http://openmymind.net/The-Little-Go-Book/) is
a good follow-up.

This book uses
- Go 1.22
- Curl
- A web browser

# 2. Foundations
## 2.1. Project setup and creating a module
Create a project directory called `snippetbox`

Modules are Go's solution to dependency management. They are a collection of related Go packages
that are versioned together as a single unit. Modules record dependency requirements and create
reproducible builds.

A module path is how that module is identified, and must be unique. If you’re creating a module
for use in other projects, then it’s good practice for the module path to be the URL where the 
module is hosted.
```bash
go mod init github.com/theknickerbocker/snippetbox
```

This will create a go.mod file in the root of the project directory. This file is the module
manifest, and contains module identifiers for the required dependencies of the module.

Create a `main.go` file and a simple hello world program.

Running `go run .` will compile the module and run the `main()` function in `./main.go`

## 2.2. Web application basics
We will create:
1. A handler - These are essentially the controllers (from MVC pattern) of the web app.
2. A router - This will match incoming requests to the correct handler function.
3. A server - This will listen for incoming requests and pass them to the router.

Update your `main.go` file to include a simple web server that listens on port 4000.

## 2.3. Routing requests
Add 2 more routes to get a feel for how routing works.

Go's servemux has different route matching rules depending on if the given route ends with a
trailing slash. If the route pattern doesn't end with a slash, then it will only match when 
the URL is exact. For example, `/snippet/views !-> /snippet/view`. If the route pattern does
end with a slash, then it will match any URL that has the pattern as a prefix. For example,
`/snippet/views -> /snippet/view`. This matching pattern is known as a subtree path pattern.

To see this in action, try navigating to `/snippet/views`. You will be redirected to `/`, because
the `/` route pattern essentially means match a single slash, followed by anything.

To explicitly prevent this behavior, add a `{$}` to the end of the route  pattern. Notice that now
`/snippet/views` will return a 404.

**Worth reading the "Additional Information: The default servemux" section**
Using http.Handle() and http.HandleFunc() registers routes with the default servemux, http.
DefaultServeMux, which is automatically initialized by Go. This approach can make code shorter but
is less explicit and can lead to security and maintainability issues due to its global nature. For
clarity and safety, it's better to use a locally-scoped servemux.

## 2.4. Wildcard route patterns
Wildcard route patterns can be used to define flexible routing rules, commonly for path parameters.

For example `/user/{id}` uses the `{id}` wildcard to match any value in the URL path. This value can
be retrieved in code with `r.PathValue("id")`.

Update the `/snippet/view` path to accept an `{id}`.

**Worth reading the "Additional Information: Precedence and conflicts" section**
Depending on the router you're using, it will define an order of precedence for route patterns when
multiple patterns overlap. For example, in the standard Go servemux if you have `/user/{id}` and
`/user/new`, the precedence of these is determined by specificy of the pattern. The more specific
pattern will be matched first, so `/user/new` will be matched before `/user/{id}`.
