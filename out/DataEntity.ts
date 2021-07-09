export interface Test2_ts   {
    Test2_001: { [id: string]: Test2_001 };
    Test2_002: { [id: string]: Test2_002 };
    NPCList: { [id: string]: NPCList };
}

export interface Test2_001  {
    /** 第几天 */
    ID: number;
    /** 签到名称 */
    NameID: string;
    /** 任务类型 */
    QuestType: number;
    /** 道具ID */
    ItemID: number[];
    /** 数量 */
    ItemCount: boolean;
    /** 倍率(百分位) */
    rate: boolean[];
    /** 文本 */
    text: string;
}

export interface Test2_002  {
    /** 第几天 */
    ID: number;
    /** 签到名称 */
    NameID: string;
    /** 任务类型 */
    QuestType: number;
    /** 道具ID */
    ItemID: number;
    /** 数量 */
    ItemCount: boolean;
}

export interface NPCList  {
    /** NPCID */
    NPCID: number;
    /** 所属地点 */
    NPCofTiled: number;
    /** NPC类型 */
    NPCType: number;
    /** NPC名字 */
    NPCName: number;
    /** NPC职业 */
    NPCJob: number;
    /** 底框资源名称 */
    NPCFrameRes: string;
    /** NPC竖条立绘 */
    NPCIcon: string;
    /** NPC功能列表 */
    NPCFuncList: number[];
    /** 支线任务 */
    TiledIcon: string;
    /** 隐藏显示方案 */
    TiledNPCList: number;
    /** 默认对白方案 */
    InitNPCFunc: number;
    /** NPC固定对白 */
    TiledTipsWord: number;
}

