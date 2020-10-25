export interface Testdir1   {
    Test1: { [id: number]: Test1 };
    Test2: { [id: number]: Test2 };
}

export interface Testdir2   {
    Test3: { [id: number]: Test3 };
    Test4: { [id: number]: Test4 };
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

export interface Test3  {
    key: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: boolean;
    rate: boolean[];
}

export interface Test4  {
    key: number;
    NameID: string;
    QuestType: number;
    ItemID: number;
    ItemCount: number;
    rate: number[];
}

