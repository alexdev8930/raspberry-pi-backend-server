# Backend architecture

The backend sits between the frontend and the systems it communicates with. Its main role is to receive requests, perform application operations, and return data. Raspberry Pi hardware is one possible integration, not a requirement for every operation.

## Components

The HTTP transport is the boundary used by the frontend. It converts HTTP and JSON into application requests and turns results into HTTP responses. Application services contain the operations and rules behind those requests. At startup, the application layer connects the transport, services, configuration, and any integrations the server needs.

Integrations sit behind interfaces owned by the backend. A device adapter communicates with Raspberry Pi hardware when an operation needs it. A repository can provide persistent data when a feature needs storage. The server can operate without either integration when its features do not require them.

## Request lifecycle

A request moves from the frontend to an HTTP handler, then to the service responsible for the operation. The service validates the request and performs the required work, using integrations when needed. The result returns through the handler as a JSON response.

Invalid input produces a client error. Requests that fail authentication or authorization checks should be rejected when those checks are implemented, and unexpected failures produce a server error. Responses do not include internal details or secrets.

## Trust and data

The backend treats frontend requests as untrusted. When accounts or shared access are supported, the backend verifies the caller's identity and permission. For operations targeting a device, a device ID identifies the target but does not grant access to it.

Credentials and other secrets should not be committed to source control. Applications that store sensitive data should use appropriate protections for stored credentials and data. Deployments exposed beyond a trusted local network should use HTTPS and a secure remote-access configuration.
