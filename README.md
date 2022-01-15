# bookstore_oauth-api

authentication Service

## Overview
In order to securitize our applications, sometimes we require to authenticate users before allowing them to enter and interact with our system.
`Access tokens` are the thing that applications use to make API requests on behalf of a user. The access token represents the authorization of a specific application to access specific parts of a user’s data.

`Access tokens` do not have to be of any particular format, although there are different considerations for different options which will be discussed later in this chapter. As far as the client application is concerned, the access token is an opaque string, and it will take whatever the string is and use it in an HTTP request. The resource server will need to understand what the access token means and how to validate it, but applications will never be concerned with understanding what an access token means.

`Access tokens` must be kept confidential in transit and in storage. The only parties that should ever see the access token are the application itself, the authorization server, and resource server. The application should ensure the storage of the access token is not accessible to other applications on the same device. The access token can only be used over an HTTPS connection, since passing it over a non-encrypted channel would make it trivial for third parties to intercept.
## Login and authentication flow
First, a user request an access token to the OAuth api and this returns the access token (AT).
Then, there is other API that require validate this AT, and in the case that the user was validate, this return the validation.

![alt](src/assets/oAuth_diagram.png)

## Project Structure
This app use Domain Driven Development design, a hexagonal architecture and Casandra database to persiste the users access token.

 