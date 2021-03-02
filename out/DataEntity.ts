export interface Test1   {
    Test1_002: { [id: string]: Test1_002 };
    Test1_001: { [id: string]: Test1_001 };
}

export interface Test2   {
    Test2_002: { [id: string]: Test2_002 };
    Test2_001: { [id: string]: Test2_001 };
}

export interface Test1_001  {
    ID: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: boolean;
    rate: boolean[];
}

export interface Test1_002  {
    ID: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: boolean;
    rate: boolean[];
}

export interface Test2_001  {
    ID: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: boolean;
    rate: boolean[];
}

export interface Test2_002  {
    ID: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: boolean;
    rate: boolean[];
}

