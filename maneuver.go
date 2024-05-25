package main

type maneuverCardType int

const (
	Strike maneuverCardType = iota
	Grapple
	Submission
	HighRisk
	Assault
	Hold
	Throw
	Extreme
	NonManeuver
)

type maneuverSubType int

const (
	Core maneuverCardType = iota
	Mean
	Colossal
	NonRevolution
)
