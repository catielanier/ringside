package prematch

type PreMatchType int

const (
	Venue PreMatchType = iota
	Feud
	Stipulation
	Event
	Object
	PPVEvent
	Manager
	NonPreMatch
)
