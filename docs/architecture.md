# architecture

## Bootstrap
Document the service entrypoint, router registration, and config loading flow used during early API wiring.
## Auth service
Note why authentication reads remain thin and mostly defer persistence logic to the model layer.
## Token handling
Describe the current JWT ownership model and the short list of claims used by the backend.
## Login token flow
Summarize how login resolves a user, checks the password, and issues a token for downstream calls.
## Service boundaries
Describe the split between controller level parameter handling and service level account rules.
## Model mapping
Summarize how user and verification models map to the database layer and stay thin on business rules.
## Session handling
Describe the current session model as stateless JWT validation backed by user lookups when required.
## Token renewal discussion
Capture the open question around refresh support and why the initial backend keeps token behavior minimal.
## Ownership
Record the backend ownership assumptions for auth, account verification, and deployment coordination.
