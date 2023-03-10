import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCancelExchange } from "./types/fractal/fractal/tx";
import { MsgRequestExchange } from "./types/fractal/fractal/tx";
import { MsgSwapExchange } from "./types/fractal/fractal/tx";
import { MsgApproveExchange } from "./types/fractal/fractal/tx";
import { MsgSettleExchange } from "./types/fractal/fractal/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/fractal.fractal.MsgCancelExchange", MsgCancelExchange],
    ["/fractal.fractal.MsgRequestExchange", MsgRequestExchange],
    ["/fractal.fractal.MsgSwapExchange", MsgSwapExchange],
    ["/fractal.fractal.MsgApproveExchange", MsgApproveExchange],
    ["/fractal.fractal.MsgSettleExchange", MsgSettleExchange],
    
];

export { msgTypes }