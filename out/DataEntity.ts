export interface Test2_ts   {
    Test2_001: { [id: string]: Test2_001 };
    Test2_002: { [id: string]: Test2_002 };
}

export interface Test2_001  {
    ID: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: boolean;
    rate: boolean[];
    text: string;
}

export interface Test2_002  {
    ID: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: boolean;
    rate: boolean[];
}

