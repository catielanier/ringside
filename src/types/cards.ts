import {ObjectID} from "mongodb";
import {Brand, Alignment, SecondaryAlignment} from "../utils/constants";

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
    brand: Brand;
    isGM: boolean;
    handSize: number;
    superStarValue: number;
    arsenalLimit: number;
    backlashLimit: number;
    backstageLimit: number;
    drawSize: number;
    alignment: Alignment;
    secondaryAlignment: SecondaryAlignment;
    eugeneMode: boolean;
    canAlsoPackSuperstarSpecifics: ObjectID[];
    canBeBackstage: boolean;
    nonUniqueLimit: number;
    uniqueLimit: number;
    set: string[];
    cardText: string;
    imageURL: string[];
}