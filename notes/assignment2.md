# Chapter 2 (cont.)
## 2.5. Method-based routing
Specifying an HTTP method on the mux handler route will only match requests with the same method.
They must be uppercase and have a single trailing space. Example: `POST /...`

Registering a GET method will match both GET and HEAD requests, but other methods require an 
exact match.

Paths will also automatically return a 405 for any method that is not registered, and return an
Allow header with the registered methods.

Note that Justworks uses Datadog's version of the gorilla/mux router, which has some additional
DD integrations on top of the original.

## 2.6. Customizing Responses
By default, a handler will respond to a request with a 200 status code, and a Date, 
Content-Length, and Content-Type header.

Call `w.WriteHeader()` to set the status code. This function can only be called once per request.

The `http` package provides constants for status codes (e.g. `http.StatusCreated`).

Headers can be added using `w.Header().Add(key, value)`. They must be set before 
`w.WriteHeader()` or `w.Write()` is called, or they will not be written to the response.

Go leverages interfaces (specifically the `io.Writer` interface) to allow many 
different implementations of writing response bodies.
