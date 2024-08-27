# Chapter 2 (cont.)
## 2.7. Project structure and organization
There aren't hard and fast rules to Go project structure, but there are some 
common conventions and patterns.

- `internal` directory - contains packages that are not meant to be imported by 
  other projects. This is a way to enforce encapsulation and prevent accidental 
  misuse of packages. A good place to store application-agnostic logic for reuse
- `cmd` directory - contains the main entry points for the application. This is 
  where the application-specific logic should live.

#### Useful Resources:
- [Go Server Project Layout](https://go.dev/doc/modules/layout#server-project)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

## 2.8. HTML templating and inheritance
Go has a built-in templating package that allows you to create and render HTML 
templates `html/template`.

Use `http.Error()` to send an error response to the client with a message and status.
