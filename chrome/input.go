package chrome

// MouseEventType is a JavaScript mouse event type.
type MouseEventType string

const (
	MousePressed  MouseEventType = "mousePressed"
	MouseReleased                = "mouseReleased"
	MouseMoved                   = "mouseMoved"
)

// MouseButton is a button identifier for a mouse event.
type MouseButton string

const (
	LeftButton   MouseButton = "left"
	MiddleButton MouseButton = "middle"
	RightButton  MouseButton = "right"
)

// MouseEvent stores information about a JavaScript mouse
// event.
type MouseEvent struct {
	Type MouseEventType
	X    int
	Y    int

	// May be "" for movement events.
	Button MouseButton

	ClickCount int
}

// DispatchMouseEvent triggers a mouse event.
func (c *Conn) DispatchMouseEvent(m *MouseEvent) error {
	params := map[string]interface{}{
		"type":       m.Type,
		"x":          m.X,
		"y":          m.Y,
		"clickCount": m.ClickCount,
	}
	if m.Button != "" {
		params["button"] = m.Button
	}
	return c.call("Input.dispatchMouseEvent", params, nil)
}
