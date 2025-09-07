package usecases

import (
	"ampel-web/internal/domain"
	"ampel-web/internal/execptions"
	"ampel-web/pkg/ampel"
)

var colorMapping map[domain.LedColor]ampel.Color = map[domain.LedColor]ampel.Color{
	domain.Red:    ampel.Red,
	domain.Orange: ampel.Orange,
	domain.Green:  ampel.Green,
}

type LedUsesCases struct {
}

func New() *LedUsesCases {
	return &LedUsesCases{}
}

func (*LedUsesCases) TreatLedEvent(event domain.Event) error {
	if evt, ok := event.(*domain.EnableLedEvent); ok {
		return ampel.LedOn(colorMapping[evt.Color])
	} else if evt, ok := event.(*domain.DisbaleLedEvent); ok {
		return ampel.LedOff(colorMapping[evt.Color])
	} else {
		return execptions.BadDomainEventError
	}
}
