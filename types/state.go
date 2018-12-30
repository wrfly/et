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
		return "StateNormal"
	case StateStopped:
		return "StateStopped"
	case StateResumed:
		return "StateResumed"
	case StateTerminated:
		return "StateTerminated"
	default:
		return "unknown"
	}
}
