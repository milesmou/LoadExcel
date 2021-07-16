export interface TestBattleKin   {
    BattleRoles: { [id: string]: BattleRoles };
}

export interface BattleRoles  {
    /** 侠客id */
    ID: number;
    /** 名字 */
    Name: string;
    /** 形象ID */
    RolePic: string;
    /** 技能名字 */
    SkillName: string;
    /** 等级 */
    Level: number;
    /** 调用元素 */
    Element: number;
    /** 普攻 */
    EommonAtk: string;
    /** 技能 */
    SkillAtk: string;
    /** 生命 */
    Hp: number;
    /** 物理攻击 */
    attack_damage: number;
    /** 火焰攻击 */
    fire_damage: number;
    /** 冰霜攻击 */
    ice_damage: number;
    /** 毒素攻击 */
    poison_damage: number;
    /** 神圣攻击 */
    holy_damage: number;
    /** 暗影攻击 */
    dark_damage: number;
    /** 雷电攻击 */
    Lightning_damage: number;
    /** 物理抗性 */
    attack_defense: number;
    /** 火焰抗性 */
    fire_defense: number;
    /** 冰霜抗性 */
    ice_defense: number;
    /** 毒素抗性 */
    poison_defense: number;
    /** 神圣抗性 */
    holy_defense: number;
    /** 暗影抗性 */
    dark_defense: number;
    /** 雷电抗性 */
    Lightning_defense: number;
    /** 速度 */
    speed: number;
    /** 精准% */
    accurate: number;
    /** 闪避% */
    dodge: number;
    /** 暴击% */
    crit: number;
    /** 韧性% */
    tenacity: number;
    /** 暴击伤害 */
    critDamage: number;
    /** 真实伤害 */
    realDamage: number;
    /** 击中生命恢复 */
    hitAddHp: number;
    /** 击杀生命恢复 */
    killAddHp: number;
    /** 战斗生命恢复 */
    battleAddHp: number;
    /** 行动生命恢复 */
    actionAddHp: number;
    /** 免伤 */
    finalReduced: number;
    /** 治疗加成 */
    increasedHealEff: number;
    /** 满怒 */
    fullAnger: number;
    /** 被击加怒 */
    beHitAddAnger: number;
    /** 攻击加怒 */
    hitAddAnger: number;
    /** 击杀加怒 */
    killAddAnger: number;
    /** 战斗加怒 */
    battleAddAnger: number;
}

