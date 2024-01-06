package utils

const (
	Success        = "Success"
	BadRequest     = "BadRequest"
	InternalServer = "InternalServer"
	NotFound       = "NotFound"
	Created        = "Created"
	Unauthorized   = "Unauthorized"
	Forbidden      = "Forbidden"
)

var StatusCode = map[string]int{
	"Success":        200,
	"Created":        201,
	"BadRequest":     400,
	"NotFound":       404,
	"InternalServer": 500,
	"Unauthorized":   401,
	"Forbidden":      403,
}
