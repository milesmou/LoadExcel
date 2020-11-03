export interface Test2   {
    Test2_001: { [id: number]: Test2_001 };
    Test2_002: { [id: number]: Test2_002 };
}

export interface Test1   {
    Test1_001: { [id: number]: Test1_001 };
    Test1_002: { [id: number]: Test1_002 };
}

export interface Test2_001  {
    key: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: number;
    rate: number[];
}

export interface Test2_002  {
    key: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: number;
    rate: number[];
}

export interface Test1_001  {
    key: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: boolean;
    rate: boolean[];
}

export interface Test1_002  {
    key: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: boolean;
    rate: boolean[];
}

