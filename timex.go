package timex

// Keep on imports of "time" only.
import . "time"

const (
	// The ZoneName is "Timex".
	ZoneName = "Timex"
	// A Day is 24 hours.
	Day Duration = 24 * Hour
	// A Week is 7 days.
	Week Duration = 7 * Day
)

var initialLocation = Local

// TimeFunc arguments time.Time return error.
type TimeFunc func(tm Time) error

// SetLocation sets local location.
func SetLocation(offset int) {
	Local = FixedZone(ZoneName, offset)
}

// SetUTCLocation sets UTC location.
func SetUTCLocation() {
	Local = UTC
}

// ResetInitialLocation sets initial location.
func ResetInitialLocation() {
	Local = initialLocation
}

// NowTruncate returns the truncated current local time.
func NowTruncate(d Duration) Time {
	return Now().Truncate(d)
}

// Between checks target time between before time and after time.
func Between(target, before, after Time) bool {
	return before.Before(target) && after.After(target)
}
