export interface GameConfig   {
    Config: { [id: string]: Config };
}

export interface Config  {
    /** ID */
    ID: number;
    /** 结算插屏起始关卡 */
    LvEnd_InsAD: number[];
    /** 结算插屏弹出概率 */
    LvEnd_InsAD_Random: number[];
    /** 关卡插屏起始关卡 */
    LvOpint_InsAD: number[];
    /** 关卡插屏弹出概率 */
    LvOpint_InsAD_Random: number[];
}

