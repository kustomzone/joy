package window

// js:"EventTarget,omit"
type EventTarget interface {

	// AddEventListener
	AddEventListener(kind string, listener *func(evt Event) handleEvent, useCapture *bool)

	// DispatchEvent
	DispatchEvent(evt Event) (b bool)

	// RemoveEventListener
	RemoveEventListener(kind string, listener *func(evt Event) handleEvent, useCapture *bool)
}
