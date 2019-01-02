package types

type State int

// stats
// normal -> stopped -> resumed -> terminated
const (
	StateNormal State = iota
	StateStopped
	StateResumed
	StateTerminated
)

func (s State) String() string {
	switch s {
	case StateNormal:
		return "normal"
	case StateStopped:
		return "stopped"
	case StateResumed:
		return "resumed"
	case StateTerminated:
		return "terminated"
	default:
		return "unknown"
	}
}
