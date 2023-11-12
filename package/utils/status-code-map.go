package utils

const (
	Success        = "Success"
	BadRequest     = "BadRequest"
	InternalServer = "InternalServer"
	NotFound       = "NotFound"
)

var StatusCode = map[string]int{
	"Success":        200,
	"BadRequest":     400,
	"NotFound":       404,
	"InternalServer": 500,
}
