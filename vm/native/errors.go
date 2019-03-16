package native

import "errors"

var (
	errB64Decode = errors.New("illegal base64 data")
	errJSONData  = errors.New("illegal json data")
)
