package domain

type LedColor int

const (
	Red LedColor = iota + 1
	Green
	Orange
)

type LedState struct {
	Color LedColor
}
