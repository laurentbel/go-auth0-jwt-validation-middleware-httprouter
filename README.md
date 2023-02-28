# go-auth0-jwt-validation-middleware-httprouter
A basic middleware to implement jwt validation for Auth0 while using julienschmidt/httprouter

# Usage
Replace the variables issuer and scope in main.go with your Auth0 settings.

Then run:
```
go mod download
go run main.go
```

Browse the following url to test:
- http://localhost:8080/helloUnsecure/John
- http://localhost:8080/helloSecure/John

The second url above require est JWT token passed as a bearer token.

The following Url are illustrating how the handlers can be used with parameters. In this example, we just append a string " !!!":
- http://localhost:8080/helloUnsecureWrapped/John
- http://localhost:8080/helloSecureWrapped/John