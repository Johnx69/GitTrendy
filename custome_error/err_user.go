package custom_error

import "errors"

var (
    UserConflict   = errors.New("User already exists")
    UserNotFound   = errors.New("User not found")
    UserNotUpdated = errors.New("Failed to update user information")
    SignUpFail     = errors.New("Sign up failed")
)
