package execptions

import "errors"

const (
	notFoundErrorMsg string = "Not Found"
	noLampConnected  string = "Lamp not connected"
)

var (
	NotFoundError        error = errors.New(notFoundErrorMsg)
	NoLampConnectedError       = errors.New(noLampConnected)
)
