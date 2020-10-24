export interface IJsonData   {
    Test1: { [id: number]: Test1 };
    Test2: { [id: number]: Test2 };
}

export interface Test1  {
    key: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: boolean;
    rate: boolean[];
}

export interface Test2  {
    key: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: number;
    rate: number[];
}

