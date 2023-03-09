import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSettleExchange } from "./types/fractal/fractal/tx";
import { MsgSwapExchange } from "./types/fractal/fractal/tx";
import { MsgRequestExchange } from "./types/fractal/fractal/tx";
import { MsgApproveExchange } from "./types/fractal/fractal/tx";
import { MsgCancelExchange } from "./types/fractal/fractal/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/fractal.fractal.MsgSettleExchange", MsgSettleExchange],
    ["/fractal.fractal.MsgSwapExchange", MsgSwapExchange],
    ["/fractal.fractal.MsgRequestExchange", MsgRequestExchange],
    ["/fractal.fractal.MsgApproveExchange", MsgApproveExchange],
    ["/fractal.fractal.MsgCancelExchange", MsgCancelExchange],
    
];

export { msgTypes }