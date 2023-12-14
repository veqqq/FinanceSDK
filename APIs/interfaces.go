package APIs

// used in e.check to determine if json responses are empty (and thus unwriteable)
type IsEmptyer interface {
	IsEmpty() bool
}
