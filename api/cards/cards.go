package cards

import (
	"github.com/catielanier/ringside/cards/global"
	"github.com/catielanier/ringside/cards/maneuver"
	"github.com/catielanier/ringside/cards/prematch"
	"github.com/catielanier/ringside/cards/reversal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Superstar = struct {
	Id                            primitive.ObjectID        `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                          string                    `json:"name,omitempty" bson:"name,omitempty"`
	IsRevolution                  bool                      `json:"isRevolution,omitempty" bson:"isRevolution,omitempty"`
	CanPackEnforcer               bool                      `json:"canPackEnforcer,omitempty" bson:"canPackEnforcer,omitempty"`
	IsFemale                      bool                      `json:"isFemale,omitempty" bson:"isFemale,omitempty"`
	IsDiva                        bool                      `json:"isDiva,omitempty" bson:"isDiva,omitempty"`
	IsLegend                      bool                      `json:"isLegend,omitempty" bson:"isLegend,omitempty"`
	IsNXT                         bool                      `json:"isNxt,omitempty" bson:"isNxt,omitempty"`
	IsNWO                         bool                      `json:"isNwo,omitempty" bson:"isNwo,omitempty"`
	IsTagTeam                     bool                      `json:"isTagTeam,omitempty" bson:"isTagTeam,omitempty"`
	HandSize                      int                       `json:"handSize,omitempty" bson:"handSize,omitempty"`
	SuperStarValue                int                       `json:"superStarValue,omitempty" bson:"superStarValue,omitempty"`
	ArsenalLimit                  int                       `json:"arsenalLimit,omitempty" bson:"arsenalLimit,omitempty"`
	BacklashLimit                 int                       `json:"backlashLimit,omitempty" bson:"backlashLimit,omitempty"`
	BackstageLimit                int                       `json:"backstageLimit,omitempty" bson:"backstageLimit,omitempty"`
	DrawSize                      int                       `json:"drawSize,omitempty" bson:"drawSize,omitempty"`
	Alignment                     global.Alignment          `json:"alignment,omitempty" bson:"alignment,omitempty"`
	SecondaryAlignment            global.SecondaryAlignment `json:"secondaryAlignment,omitempty" bson:"secondaryAlignment,omitempty"`
	EugeneMode                    bool                      `json:"eugeneMode,omitempty" bson:"eugeneMode,omitempty"`
	CanAlsoPackSuperstarSpecifics []primitive.ObjectID      `json:"canAlsoPackSuperstarSpecifics,omitempty" bson:"canAlsoPackSuperstarSpecifics,omitempty"`
	CanBeBackstage                bool                      `json:"canBeBackstage,omitempty" bson:"canBeBackstage,omitempty"`
	NonUniqueLimit                int                       `json:"nonUniqueLimit,omitempty" bson:"nonUniqueLimit,omitempty"`
	UniqueLimit                   int                       `json:"uniqueLimit,omitempty" bson:"uniqueLimit,omitempty"`
	Set                           []string                  `json:"set,omitempty" bson:"set,omitempty"`
	CardText                      string                    `json:"cardText,omitempty" bson:"cardText,omitempty"`
	ImageURL                      []string                  `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
}

type ArsenalCard struct {
	Id                       primitive.ObjectID          `json:"_id,omitempty" bson:"_id,omitempty"`
	Title                    string                      `json:"title,omitempty" bson:"title,omitempty"`
	IsRevolution             bool                        `json:"isRevolution,omitempty" bson:"isRevolution,omitempty"`
	IsThrowback              bool                        `json:"isThrowback,omitempty" bson:"isThrowback,omitempty"`
	IsSurvivorSeries         bool                        `json:"isSurvivorSeries,omitempty" bson:"isSurvivorSeries,omitempty"`
	IsForeignObject          bool                        `json:"isForeignObject,omitempty" bson:"isForeignObject,omitempty"`
	IsChain                  bool                        `json:"isChain,omitempty" bson:"isChain,omitempty"`
	IsHeat                   bool                        `json:"isHeat,omitempty" bson:"isHeat,omitempty"`
	IsRMS                    bool                        `json:"isRMS,omitempty" bson:"isRMS,omitempty"`
	IsUnique                 bool                        `json:"isUnique,omitempty" bson:"isUnique,omitempty"`
	IsActive                 bool                        `json:"isActive,omitempty" bson:"isActive,omitempty"`
	IsBash                   bool                        `json:"isBash,omitempty" bson:"isBash,omitempty"`
	IsSetUp                  bool                        `json:"isSetUp,omitempty" bson:"isSetUp,omitempty"`
	IsDivaOnly               bool                        `json:"isDivaOnly,omitempty" bson:"isDivaOnly,omitempty"`
	IsFemaleOnly             bool                        `json:"isFemaleOnly,omitempty" bson:"isFemaleOnly,omitempty"`
	IsNXTLogo                bool                        `json:"isNxtLogo,omitempty" bson:"isNxtLogo,omitempty"`
	IsNWOLogo                bool                        `json:"isNwoLogo,omitempty" bson:"isNwoLogo,omitempty"`
	IsLegendOnly             bool                        `json:"isLegendOnly,omitempty" bson:"isLegendOnly,omitempty"`
	IsTagTeamOnly            bool                        `json:"isTagTeamOnly,omitempty" bson:"isTagTeamOnly,omitempty"`
	Brand                    global.Brand                `json:"brand,omitempty" bson:"brand,omitempty"`
	HideCardsUnderneathValue int                         `json:"hideCardsUnderneathValue,omitempty" bson:"hideCardsUnderneathValue,omitempty"`
	Alignment                global.Alignment            `json:"alignment,omitempty" bson:"alignment,omitempty"`
	SecondaryAlignment       global.SecondaryAlignment   `json:"secondaryAlignment,omitempty" bson:"secondaryAlignment,omitempty"`
	SuperstarSpecific        string                      `json:"superstarSpecific,omitempty" bson:"superstarSpecific,omitempty"`
	SuperstarsSpecificTo     []primitive.ObjectID        `json:"superstarsSpecificTo,omitempty" bson:"superstarsSpecificTo,omitempty"`
	Volley                   int                         `json:"volley,omitempty" bson:"volley,omitempty"`
	StunValue                int                         `json:"stunValue,omitempty" bson:"stunValue,omitempty"`
	CardType                 []global.ArsenalCardType    `json:"cardType,omitempty" bson:"cardType,omitempty"`
	IsMultiAction            bool                        `json:"isMultiAction,omitempty" bson:"isMultiAction,omitempty"`
	IsMultiAntic             bool                        `json:"isMultiAntic,omitempty" bson:"isMultiAntic,omitempty"`
	IsHybridMultiActionAntic bool                        `json:"isHybridMultiAntic,omitempty" bson:"isHybridMultiAntic,omitempty"`
	ManeuverCardType         []maneuver.ManeuverCardType `json:"maneuverCardType,omitempty" bson:"maneuverCardType,omitempty"`
	ManeuverSubType          maneuver.ManeuverSubType    `json:"maneuverSubType,omitempty" bson:"maneuverSubType,omitempty"`
	ReversalType             []reversal.ReversalType     `json:"reversalType,omitempty" bson:"reversalType,omitempty"`
	Fortitude                int                         `json:"fortitude,omitempty" bson:"fortitude,omitempty"`
	Damage                   int                         `json:"damage,omitempty" bson:"damage,omitempty"`
	CardText                 string                      `json:"cardText,omitempty" bson:"cardText,omitempty"`
	Set                      []string                    `json:"set,omitempty" bson:"set,omitempty"`
	ImageURL                 []string                    `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
}

type Backlash struct {
	Id                       primitive.ObjectID          `json:"_id,omitempty" bson:"_id,omitempty"`
	Title                    string                      `json:"title,omitempty" bson:"title,omitempty"`
	IsPreMatch               bool                        `json:"isPreMatch,omitempty" bson:"isPreMatch,omitempty"`
	IsMidMatch               bool                        `json:"isMidMatch,omitempty" bson:"isMidMatch,omitempty"`
	IsThrowback              bool                        `json:"isThrowback,omitempty" bson:"isThrowback,omitempty"`
	IsSurvivorSeries         bool                        `json:"isSurvivorSeries,omitempty" bson:"isSurvivorSeries,omitempty"`
	IsForeignObject          bool                        `json:"isForeignObject,omitempty" bson:"isForeignObject,omitempty"`
	IsChain                  bool                        `json:"isChain,omitempty" bson:"isChain,omitempty"`
	IsHeat                   bool                        `json:"isHeat,omitempty" bson:"isHeat,omitempty"`
	IsRMS                    bool                        `json:"isRMS,omitempty" bson:"isRMS,omitempty"`
	IsUnique                 bool                        `json:"isUnique,omitempty" bson:"isUnique,omitempty"`
	IsActive                 bool                        `json:"isActive,omitempty" bson:"isActive,omitempty"`
	IsBash                   bool                        `json:"isBash,omitempty" bson:"isBash,omitempty"`
	IsSetUp                  bool                        `json:"isSetUp,omitempty" bson:"isSetUp,omitempty"`
	IsDivaOnly               bool                        `json:"isDivaOnly,omitempty" bson:"isDivaOnly,omitempty"`
	IsFemaleOnly             bool                        `json:"isFemaleOnly,omitempty" bson:"isFemaleOnly,omitempty"`
	IsNXTLogo                bool                        `json:"isNxtLogo,omitempty" bson:"isNxtLogo,omitempty"`
	IsNWOLogo                bool                        `json:"isNwoLogo,omitempty" bson:"isNwoLogo,omitempty"`
	IsLegendOnly             bool                        `json:"isLegendOnly,omitempty" bson:"isLegendOnly,omitempty"`
	IsTagTeamOnly            bool                        `json:"isTagTeamOnly,omitempty" bson:"isTagTeamOnly,omitempty"`
	Brand                    global.Brand                `json:"brand,omitempty" bson:"brand,omitempty"`
	Alignment                global.Alignment            `json:"alignment,omitempty" bson:"alignment,omitempty"`
	SecondaryAlignment       global.SecondaryAlignment   `json:"secondaryAlignment,omitempty" bson:"secondaryAlignment,omitempty"`
	SuperstarSpecific        string                      `json:"superstarSpecific,omitempty" bson:"superstarSpecific,omitempty"`
	SuperstarsSpecificTo     []primitive.ObjectID        `json:"superstarsSpecificTo,omitempty" bson:"superstarsSpecificTo,omitempty"`
	Volley                   int                         `json:"volley,omitempty" bson:"volley,omitempty"`
	StunValue                int                         `json:"stunValue,omitempty" bson:"stunValue,omitempty"`
	PreMatchType             prematch.PreMatchType       `json:"preMatchType,omitempty" bson:"preMatchType,omitempty"`
	MidMatchType             []global.ArsenalCardType    `json:"midMatchType,omitempty" bson:"midMatchType,omitempty"`
	IsMultiAction            bool                        `json:"isMultiAction,omitempty" bson:"isMultiAction,omitempty"`
	HideCardsUnderneathValue int                         `json:"hideCardsUnderneathValue,omitempty" bson:"hideCardsUnderneathValue,omitempty"`
	IsMultiAntic             bool                        `json:"isMultiAntic,omitempty" bson:"isMultiAntic,omitempty"`
	IsHybridMultiActionAntic bool                        `json:"isHybridMultiAntic,omitempty" bson:"isHybridMultiAntic,omitempty"`
	ManeuverCardType         []maneuver.ManeuverCardType `json:"maneuverCardType,omitempty" bson:"maneuverCardType,omitempty"`
	ManeuverSubType          maneuver.ManeuverSubType    `json:"maneuverSubType,omitempty" bson:"maneuverSubType,omitempty"`
	ReversalType             []reversal.ReversalType     `json:"reversalType,omitempty" bson:"reversalType,omitempty"`
	Fortitude                int                         `json:"fortitude,omitempty" bson:"fortitude,omitempty"`
	Damage                   int                         `json:"damage,omitempty" bson:"damage,omitempty"`
	Set                      []string                    `json:"set,omitempty" bson:"set,omitempty"`
	CardText                 string                      `json:"cardText,omitempty" bson:"cardText,omitempty"`
	ImageURL                 []string                    `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
}

type Backstage struct {
	Id                       primitive.ObjectID        `json:"_id,omitempty" bson:"_id,omitempty"`
	Title                    string                    `json:"title,omitempty" bson:"title,omitempty"`
	IsEnforcer               bool                      `json:"isEnforcer,omitempty" bson:"isEnforcer,omitempty"`
	IsObject                 bool                      `json:"isObject,omitempty" bson:"isObject,omitempty"`
	IsHideable               bool                      `json:"isHideable,omitempty" bson:"isHideable,omitempty"`
	Brand                    global.Brand              `json:"brand,omitempty" bson:"brand,omitempty"`
	Enforcer                 []Enforcer                `json:"enforcer,omitempty" bson:"enforcer,omitempty"`
	Alignment                global.Alignment          `json:"alignment,omitempty" bson:"alignment,omitempty"`
	SecondaryAlignment       global.SecondaryAlignment `json:"secondaryAlignment,omitempty" bson:"secondaryAlignment,omitempty"`
	SuperstarSpecific        string                    `json:"superstarSpecific,omitempty" bson:"superstarSpecific,omitempty"`
	SuperstarsSpecificTo     []primitive.ObjectID      `json:"superstarsSpecificTo,omitempty" bson:"superstarsSpecificTo,omitempty"`
	HideCardsUnderneathValue int                       `json:"hideCardsUnderneathValue,omitempty" bson:"hideCardsUnderneathValue,omitempty"`
	Set                      []string                  `json:"set,omitempty" bson:"set,omitempty"`
	CardText                 string                    `json:"cardText,omitempty" bson:"cardText,omitempty"`
	ImageURL                 []string                  `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
}

type Enforcer struct {
	Name           string `json:"name,omitempty" bson:"name,omitempty"`
	HandSize       int    `json:"handSize,omitempty" bson:"handSize,omitempty"`
	SuperStarValue int    `json:"superStarValue,omitempty" bson:"superStarValue,omitempty"`
	DrawSize       int    `json:"drawSize,omitempty" bson:"drawSize,omitempty"`
	CardText       string `json:"cardText,omitempty" bson:"cardText,omitempty"`
}
