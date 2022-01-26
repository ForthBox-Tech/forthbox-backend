# account lifecycle

## Account lifecycle
Describe the intended email and mobile sign up journey, including verification and password setup touch points.
## Sign up validation
Summarize the username, invite code, verification code, and signup method requirements expected by the service layer.
## Password reset
Capture the guard rails around verification checks, matching passwords, and accepted lookup inputs.
## Existence checks
Explain the query parameters accepted by the account availability endpoint and why mobile values include a region prefix.
## Verification rules
Capture the lifetime, input matching, and retry assumptions around verification token checks.
## Email signup
Write down the assumptions around verified email addresses, invite codes, and password initialization during email based signup.
## Mobile signup
Clarify the mobile number format, region code handling, and verification expectations used by the signup endpoint.
## Reset password checks
Capture why password confirmation and lookup validation stay in the service boundary instead of the controller layer.
## Password storage
Document the current hashed password path and the assumption that raw passwords never leave request scope.
## User service notes
Add notes on where user specific rules live today and which validations should remain close to persistence.
