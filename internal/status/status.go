package status

// Status represents the status of an anime in the watchlist.
type Status int

// Enum values for Status using iota for auto-increment.
const (
	NotWatched  Status = iota // NotWatched indicates the anime has not been watched.
	Watching                  // Watching indicates the anime is currently being watched.
	Completed                 // Completed indicates the anime has been finished.
	PlanToWatch               // PlanToWatch indicates the anime is planned to be watched.
	OnHold                    // OnHold indicates the anime is temporarily paused.
	Dropped                   // Dropped indicates the anime has been abandoned.
	Unknown                   // Unknown is used for any undefined status.
)

// Value returns the integer value of the Status.
func (s Status) Value() int {
	return int(s)
}

// String returns the string representation of the Status.
func (s Status) String() string {
	switch s {
	case NotWatched:
		return "NotWatched"
	case Watching:
		return "Watching"
	case Completed:
		return "Completed"
	case PlanToWatch:
		return "PlanToWatch"
	case OnHold:
		return "OnHold"
	case Dropped:
		return "Dropped"
	default:
		return "Unknown"
	}
}

// Valid checks if a given integer value is a valid Status.
func Valid(value int) bool {
	if value < NotWatched.Value() || value >= Unknown.Value() {
		return false
	}
	return true
}
