# Auth0 Setup

This service is an API. Sign users in through an Auth0 application, request an access token for your Auth0 API, and send that token to this service:

```http
Authorization: Bearer <access-token>
```

Create an API in the Auth0 Dashboard and set its signing algorithm to RS256. Configure these environment variables before starting the service:

```sh
export AUTH0_ISSUER="https://YOUR_TENANT.auth0.com/"
export AUTH0_AUDIENCE="YOUR_AUTH0_API_IDENTIFIER"
```

`AUTH0_ISSUER` must exactly match the token's `iss` claim, including the trailing slash. `AUTH0_AUDIENCE` must match the API identifier and the token's `aud` claim. The service downloads the tenant's signing keys from its OpenID Connect metadata and caches them for five minutes. Tokens must use RS256 and have a valid issuer, audience, and expiration.

The `/login` and `/register` API endpoints no longer authenticate users. Auth0 handles those flows; use the access token issued to your client when calling `/employees`.