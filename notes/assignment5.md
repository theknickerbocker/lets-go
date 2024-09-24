# 4.1 - 4.9 ("Setting up MySQL" -> "Transactions and other details")
Not writing notes for 4.1-4.2 as they relate to setting up MySQL

## 4.3. Modules and reproducible builds
The go.mod file is the module manifest. It contains the module identifier for the required dependencies 
of the module. The go.sum file contains the cryptographic checksums of the module's dependencies. 
This ensures that the same version of a dependency is used across all builds.

`go mod tidy` will remove any dependencies that are no longer required by the module.
`go mod download` will download the module's dependencies.
`go mod verify` will verify the checksums of the module's dependencies.

## 4.4. Creating a database connection pool
A connection pool is a cache of database connections that can be reused. It is safe to use for
concurrent access and is more effiecient than opening a new connection for each request. Instead
it pools active connections that are ready to be used.

To initialize a module that is not referenced in code, use the `_` identifier to avoid the "unused
import" error.

## 4.5. Designing a database model
Encapsulating the database model in a separate package allows for separation of
concerns. The pattern of implementing functions on the model struct is known as
the "active record" pattern. It's an error-prone pattern for anything other than
simple cases.
