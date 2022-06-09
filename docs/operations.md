# operations

## Mail integration
Track the provider assumptions, domain ownership, and the minimal operational data required to send verification mail.
## Shutdown behavior
Record the expectation that in flight requests receive a short drain window during process termination.
## Migration command
Add a short note on the standalone migration entrypoint and the expectation that schema work remains explicit.
## Database expectations
Document the current database bootstrap path and the assumption that connection settings arrive through ini files.
## Mail templates
Capture the assumption that outbound verification mail uses a small shared template set owned by the backend team.
## Operational logging
Note the limited logging footprint in the service and the need to keep startup and shutdown messages readable.
## Production mail caveats
Record the need to protect provider credentials and verify domain health before enabling mail dependent signup flows.
## Migration reminders
Capture the preference for explicit database changes and the need to test migration entrypoints against safe environments first.
## Verification cleanup
Record the expectation that expired verification data can be cleaned by later maintenance work without changing request flows.
## Timeout choices
Capture why the server keeps read and write timeout values explicit in config rather than hidden in code.
## Migration safety
Add reminders to validate database backups and environment targeting before running migration commands.
