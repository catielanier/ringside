package cards

import (
	"github.com/catielanier/ringside/cards/global"
	"github.com/catielanier/ringside/cards/maneuver"
	"github.com/catielanier/ringside/cards/prematch"
	"github.com/catielanier/ringside/cards/reversal"
)

type Superstar = struct {
	Name               string
	IsRevolution       bool
	CanPackEnforcer    bool
	HandSize           int
	SuperStarValue     int
	ArsenalLimit       int
	BacklashLimit      int
	BackstageLimit     int
	DrawSize           int
	Alignment          global.Alignment
	SecondaryAlignment global.SecondaryAlignment
	EugeneMode         bool
	NonUniqueLimit     int
	UniqueLimit        int
	CardText           string
}

type ArsenalCard struct {
	Title                    string
	IsRevolution             bool
	IsThrowback              bool
	IsSurvivorSeries         bool
	IsForeignObject          bool
	IsChain                  bool
	IsHeat                   bool
	IsRMS                    bool
	IsUnique                 bool
	IsActive                 bool
	IsBash                   bool
	IsSetUp                  bool
	Brand                    global.Brand
	Alignment                global.Alignment
	SecondaryAlignment       global.SecondaryAlignment
	SuperstarSpecific        string
	Volley                   int
	StunValue                int
	CardType                 []global.ArsenalCardType
	IsMultiAction            bool
	IsMultiAntic             bool
	IsHybridMultiActionAntic bool
	ManeuverCardType         []maneuver.ManeuverCardType
	ManeuverSubType          maneuver.ManeuverSubType
	ReversalType             []reversal.ReversalType
	Fortitude                int
	Damage                   int
	CardText                 string
}

type Backlash struct {
	Title                    string
	IsPreMatch               bool
	IsMidMatch               bool
	IsThrowback              bool
	IsSurvivorSeries         bool
	IsForeignObject          bool
	IsChain                  bool
	IsHeat                   bool
	IsRMS                    bool
	IsUnique                 bool
	IsActive                 bool
	IsBash                   bool
	IsSetUp                  bool
	Brand                    global.Brand
	Alignment                global.Alignment
	SecondaryAlignment       global.SecondaryAlignment
	SuperstarSpecific        string
	Volley                   int
	StunValue                int
	PreMatchType             prematch.PreMatchType
	MidMatchType             []global.ArsenalCardType
	IsMultiAction            bool
	IsMultiAntic             bool
	IsHybridMultiActionAntic bool
	ManeuverCardType         []maneuver.ManeuverCardType
	ManeuverSubType          maneuver.ManeuverSubType
	ReversalType             []reversal.ReversalType
	Fortitude                int
	Damage                   int
	CardText                 string
}

type Backstage struct {
	Title              string
	IsEnforcer         bool
	IsObject           bool
	IsHideable         bool
	Brand              global.Brand
	Enforcer           []Enforcer
	Alignment          global.Alignment
	SecondaryAlignment global.SecondaryAlignment
	CardText           string
}

type Enforcer struct {
	Name           string
	HandSize       int
	SuperStarValue int
	CardText       string
}
