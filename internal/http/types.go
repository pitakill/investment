package http

var (
	// Invalid json
	errInvalidJSON = "not valid JSON"
)

type BadRequest struct {
	Error string `json:"error"`
}
