package constants

var (
	ErrUnauthorized           = "You are not authorized to access this resource"
	ErrForbidden              = "You are forbidden to access this resource"
	ErrUnauthenticated        = "You are not authenticated"
	ErrInternalServer         = "Internal server error"
	ErrFailParseID            = "Failed to parse ID"
	ErrNotFound               = "Resource not found"
	ErrAddressNotFound        = "Address not found"
	ErrInvalidEmailorPassword = "Invalid email or password"
	ErrInvalidID              = "Invalid ID"
	ErrAuthenticationFailed   = "Authentication failed"
	ErrEmailAlreadyExist      = "Email already exist"
	ErrBadRequest             = "Bad request"
	ErrFieldRequired          = "Field is required"
	ErrUserAlreadyActive      = "User already active"
	ErrUserAlreadyDeleted     = "User not found or already deleted"
	ErrInvalidUserIDToken     = "Invalid user ID in token"
)
