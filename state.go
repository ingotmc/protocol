package protocol

// State represents the state of the connection
type State int32

const (
	Handshaking State = iota
	Status
	Login
	Play
)


