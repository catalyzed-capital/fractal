import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRequestExchange } from "./types/fractal/fractal/tx";
import { MsgCancelExchange } from "./types/fractal/fractal/tx";
import { MsgSettleExchange } from "./types/fractal/fractal/tx";
import { MsgApproveExchange } from "./types/fractal/fractal/tx";
import { MsgSwapExchange } from "./types/fractal/fractal/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/fractal.fractal.MsgRequestExchange", MsgRequestExchange],
    ["/fractal.fractal.MsgCancelExchange", MsgCancelExchange],
    ["/fractal.fractal.MsgSettleExchange", MsgSettleExchange],
    ["/fractal.fractal.MsgApproveExchange", MsgApproveExchange],
    ["/fractal.fractal.MsgSwapExchange", MsgSwapExchange],
    
];

export { msgTypes }