import { ObjectID } from "mongodb";
import {
    BRAND,
    ALIGNMENT,
    SECONDARY_ALIGNMENT,
    ARSENAL_CARD_TYPE,
    REVERSAL_TYPE,
    MANEUVER_TYPE,
    MANEUVER_SUB_TYPE,
    PRE_MATCH_TYPE
} from "../utils/constants";

export interface Superstar {
    _id: ObjectID;
    name: string;
    isRevolution: boolean;
    canPackEnforcer: boolean;
    isFemale: boolean;
    isDiva: boolean;
    isLegend: boolean;
    isNXT: boolean;
    isNWO: boolean;
    isTagTeam: boolean;
    brand: typeof BRAND[number];
    isGM: boolean;
    handSize: number;
    superStarValue: number;
    arsenalLimit: number;
    backlashLimit: number;
    backstageLimit: number;
    drawSize: number;
    alignment: typeof ALIGNMENT[number];
    secondaryAlignment: typeof SECONDARY_ALIGNMENT[number];
    eugeneMode: boolean;
    canAlsoPackSuperstarSpecifics: ObjectID[];
    canBeBackstage: boolean;
    nonUniqueLimit: number;
    uniqueLimit: number;
    set: string[];
    cardText: string;
    imageURL: string[];
}

export interface ArsenalCard {
    id: ObjectID;
    title: string;
    isRevolution: boolean;
    isThrowback: boolean;
    isSurvivorSeries: boolean;
    isForeignObject: boolean;
    isChain: boolean;
    isHeat: boolean;
    isRMS: boolean;
    isUnique: boolean;
    isActive: boolean;
    isBash: boolean;
    isSetUp: boolean;
    isRunIn: boolean;
    isDivaOnly: boolean;
    isFemaleOnly: boolean;
    isNXTLogo: boolean;
    isNWOLogo: boolean;
    isLegendOnly: boolean;
    isTagTeamOnly: boolean;
    brand: typeof BRAND[number];
    hideCardsUnderneathValue: number;
    alignment: typeof ALIGNMENT[number];
    secondaryAlignment: typeof SECONDARY_ALIGNMENT[number];
    superstarsSpecificTo: ObjectID[];
    volley: number;
    stunValue: number;
    cardType: typeof ARSENAL_CARD_TYPE[number][];
    isMultiAction: boolean;
    isMultiAntic: boolean;
    isHybridMultiActionAntic: boolean;
    maneuverCardType: typeof MANEUVER_TYPE[number][];
    maneuverSubType: typeof MANEUVER_SUB_TYPE[number];
    reversalType: typeof REVERSAL_TYPE[number][];
    fortitude: number;
    damage: number;
    cardText: string;
    quote: string;
    set: string[];
    imageURL: string[];
}

export interface Backlash {
    id: ObjectID;
    title: string;
    isPreMatch: boolean;
    isMidMatch: boolean;
    isThrowback: boolean;
    isSurvivorSeries: boolean;
    isForeignObject: boolean;
    isChain: boolean;
    isHeat: boolean;
    isRMS: boolean;
    isUnique: boolean;
    isActive: boolean;
    isBash: boolean;
    isSetUp: boolean;
    isRunIn: boolean;
    isDivaOnly: boolean;
    isFemaleOnly: boolean;
    isNxtLogo: boolean;
    isNwoLogo: boolean;
    isLegendOnly: boolean;
    isTagTeamOnly: boolean;
    brand: typeof BRAND[number];
    alignment: typeof ALIGNMENT[number];
    secondaryAlignment: typeof SECONDARY_ALIGNMENT[number];
    superstarsSpecificTo: ObjectID[];
    volley: number;
    stunValue: number;
    preMatchType: typeof PRE_MATCH_TYPE[number];
    midMatchType: typeof ARSENAL_CARD_TYPE[number][];
    isMultiAction: boolean;
    hideCardsUnderneathValue: number;
    isMultiAntic: boolean;
    isHybridMultiActionAntic: boolean;
    maneuverCardType: typeof MANEUVER_TYPE[number][];
    maneuverSubType: typeof MANEUVER_SUB_TYPE[number];
    reversalType: typeof REVERSAL_TYPE[number][];
    fortitude: number;
    damage: number;
    set: string[];
    cardText: string;
    quote: string;
    imageUrl: string[];
}

interface Enforcer {
    name: string;
    handSize: number;
    superStarValue: number;
    drawSize: number;
    cardText: string;
}

export interface Backstage {
    id: ObjectID;
    title: string;
    isEnforcer: boolean;
    isObject: boolean;
    isHideable: boolean;
    brand: typeof BRAND[number];
    enforcer: Enforcer[];
    alignment: typeof ALIGNMENT[number];
    secondaryAlignment: typeof SECONDARY_ALIGNMENT[number];
    superstarsSpecificTo: ObjectID[];
    hideCardsUnderneathValue: number;
    set: string[];
    cardText: string;
    imageUrl: string[];
}