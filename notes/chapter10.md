# 10.3 User signup and password encryption

## Interpreted vs raw string literal
    
```go
    "Hello\nWorld!" // Interpreted string literal
    // -> Hello 
    // -> World!

    `Hello\nWorld!` // Raw string literal
    // -> Hello\nWorld!
```

## Password Encryption
Golang has a package called bcypt that implements the bcrypt algorithm with
helper functions specifically for creating and checking passwords. 

We don't want to store the password in plain text in the database. Instead, we
store a hash.

Using bcrypt, we generate a hash with a "12" in cost. 12 is a good minimum, but 
higher is more difficult to crack. Higher also means it is more expensive to generate
on user signup, so load testing may be needed to best understand the cost.

It's usually best to do password encryption in application as opposed to a database
encryption function. Encryption is not a constant time function, so it can be used
to determine the length of the password. Also, database logs may capture a password
in plain text before encryption (GitHub, Twitter).

## Storing User Details
Use `errors.As` to check if an error is a `MySQLError`. If it is, we can check the
number to see if it is violating the unique email constraint.

If there is an issue with duplicate emails, then display this to the user in the 
signup form. If it was a success, flash it to the user and redirect to login.

# 10.4 User login
## Checking credentials
First check that the given email exists, then check that the given password's hash 
matches the stored hash. If both are true, then the user is authenticated and an ID 
is returned. If either is false then display a generic error to the user for invalid 
credentials. 

# 10.5 User logout
In order to logout, all we need to do is remove the "authenticatedUserID" from the session.

# 10.6 User authorization
Use whether there is an authenticated user ID in the session to determine if the user is 
authenticated. If they are, then change UI elements and block pages. Easiest way to block
pages/routes is via middleware.

# 10.7 CSRF Protection
Cross Site Request Forgery is when a malicious site sends state-changing HTTP requests to
your website (explanation [here](https://www.gnucitizen.org/blog/csrf-demystified/)). General
idea is that a logged in user navigates on their browser to a malicious site, and that site
will then send a cross-site request to a POST endpoint to change state with the user's session
cookie. Our site processes the request with the authenticated cookie, and the user has no idea.

Setting the `SameSite` attribute on the cooke will mitigate sending the cookie cross-site, but
it's not yet supported by all browsers so isn't guaranteed.

The best way to protect against CSRF for all users is a token check. Using a common lib lik
`gorilla/csrf` is a good way to implement this. This will send a random CSRF token in a CSRF
cookie. That token is then added to a hidden field in each HTML form that is vulnerable to CSRF.
When the form is submitted, the token is checked against the cookie. If they match, then the request
