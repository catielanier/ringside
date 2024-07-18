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
	TrademarkFinisher
	Trademark
	NonManeuver
)

type ManeuverSubType int

const (
	Core ManeuverCardType = iota
	Mean
	Colossal
	NonRevolution
)
