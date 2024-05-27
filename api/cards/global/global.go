package global

type ArsenalCardType int
type Alignment int
type SecondaryAlignment int
type Brand int

const (
	Heel Alignment = iota
	Face
	Neutral
	Any
)

const (
	Cheater SecondaryAlignment = iota
	FanFavorite
	SecondaryNeutral
	SecondaryAny
)

const (
	Maneuver ArsenalCardType = iota
	Action
	Antic
	Reversal
)

const (
	Raw Brand = iota
	Smackdown
	NoBrand
)
