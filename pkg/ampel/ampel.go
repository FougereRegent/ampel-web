package ampel

// #cgo LDFLAGS: -L./lib/ -lled -lusb-1.0
// #include "ampel-lib.h"
import "C"
import (
	"errors"
	"unsafe"
)

type Color int

var ampelLed *C.libampel_ampel_led

var (
	Red    Color = 0x10
	Orange Color = 0x11
	Green  Color = 0x12
)

func init() {
	result := C.init((**C.libampel_ampel_led)(unsafe.Pointer(&ampelLed)))
	if result < 0 {
		err := treatErrorCode(result)
		panic(err)
	}

	LedOff(Green)
	LedOff(Red)
	LedOff(Orange)
}

func LedOn(color Color) error {
	cColor := colorMapping(color)
	return applyLed(cColor, true)
}

func LedOff(color Color) error {
	cColor := colorMapping(color)
	return applyLed(cColor, false)
}

func GetLastState() (color Color, isOn bool) {
	var resultLastState C.struct_libampel_state
	resultLastState = C.libampel_get_last_led(ampelLed)
	color = reverseColorMapping(resultLastState.color)

	if resultLastState.state == C.ON {
		isOn = true
	} else {
		isOn = false
	}
	return
}

func ReleaseAmpelLed() {
	C.release_ampel(ampelLed)
}

func applyLed(color C.LED_COLOR, isOn bool) error {
	var state C.struct_libampel_state
	state.color = color
	if isOn {
		state.state = C.ON
	} else {
		state.state = C.OFF
	}

	result := C.libampel_apply_value(ampelLed, state)
	if result < 0 {
		return treatErrorCode(result)
	}
	return nil
}

func colorMapping(color Color) (colorLed C.LED_COLOR) {
	switch color {
	case Red:
		colorLed = C.RED
	case Green:
		colorLed = C.GREEN
	case Orange:
		colorLed = C.ORANGE
	}
	return
}

func reverseColorMapping(colorLed C.LED_COLOR) (color Color) {
	switch colorLed {
	case C.RED:
		color = Red
	case C.GREEN:
		color = Green
	case C.ORANGE:
		color = Orange
	}
	return
}

func treatErrorCode(errorCode C.int) error {
	errorMessage := C.libampel_strerror(errorCode)
	return errors.New(C.GoString(errorMessage))
}
