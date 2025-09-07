package execptions

import (
	"errors"
)

const (
	notFoundErrorMsg string = "Not Found"
	noLampConnected  string = "Lamp not connected"
	badDomainEvent   string = "Bad domaine event"
	badRequest       string = "Bad request"
	colorNotExist    string = "This color doesn't exist"
)

var (
	NotFoundError        error = errors.New(notFoundErrorMsg)
	NoLampConnectedError       = errors.New(noLampConnected)
	BadDomainEventError        = errors.New(badDomainEvent)
	BadRequestError            = errors.New(badRequest)
	ColorNotExistError         = errors.New(colorNotExist)
)
