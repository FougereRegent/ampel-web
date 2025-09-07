package domain

type Event interface {
	GetEventName() string
}

type EnableLedEvent struct {
	Color LedColor
}

type DisbaleLedEvent struct {
	Color LedColor
}

func (*EnableLedEvent) GetEventName() string {
	return "led.enabled"
}

func (*DisbaleLedEvent) GetEventName() string {
	return "led.disabled"
}
