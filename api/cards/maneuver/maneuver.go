package maneuver

type ManeuverCardType int

const (
	Strike ManeuverCardType = iota
	Grapple
	Submission
	HighRisk
	Assault
	Hold
	Throw
	Extreme
	NonManeuver
)

type ManeuverSubType int

const (
	Core ManeuverCardType = iota
	Mean
	Colossal
	NonRevolution
)
