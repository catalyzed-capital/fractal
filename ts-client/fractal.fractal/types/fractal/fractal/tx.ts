/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "fractal.fractal";

export interface MsgRequestExchange {
  creator: string;
  finalunit: string;
  amount: string;
  startunit: string;
  unitratio: string;
  settledate: string;
}

export interface MsgRequestExchangeResponse {
}

export interface MsgApproveExchange {
  creator: string;
  id: number;
}

export interface MsgApproveExchangeResponse {
}

export interface MsgSettleExchange {
  creator: string;
  id: number;
}

export interface MsgSettleExchangeResponse {
}

export interface MsgSwapExchange {
  creator: string;
  id: number;
}

export interface MsgSwapExchangeResponse {
}

export interface MsgCancelExchange {
  creator: string;
  id: number;
}

export interface MsgCancelExchangeResponse {
}

function createBaseMsgRequestExchange(): MsgRequestExchange {
  return { creator: "", finalunit: "", amount: "", startunit: "", unitratio: "", settledate: "" };
}

export const MsgRequestExchange = {
  encode(message: MsgRequestExchange, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.finalunit !== "") {
      writer.uint32(18).string(message.finalunit);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    if (message.startunit !== "") {
      writer.uint32(34).string(message.startunit);
    }
    if (message.unitratio !== "") {
      writer.uint32(42).string(message.unitratio);
    }
    if (message.settledate !== "") {
      writer.uint32(50).string(message.settledate);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRequestExchange {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRequestExchange();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.finalunit = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        case 4:
          message.startunit = reader.string();
          break;
        case 5:
          message.unitratio = reader.string();
          break;
        case 6:
          message.settledate = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRequestExchange {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      finalunit: isSet(object.finalunit) ? String(object.finalunit) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      startunit: isSet(object.startunit) ? String(object.startunit) : "",
      unitratio: isSet(object.unitratio) ? String(object.unitratio) : "",
      settledate: isSet(object.settledate) ? String(object.settledate) : "",
    };
  },

  toJSON(message: MsgRequestExchange): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.finalunit !== undefined && (obj.finalunit = message.finalunit);
    message.amount !== undefined && (obj.amount = message.amount);
    message.startunit !== undefined && (obj.startunit = message.startunit);
    message.unitratio !== undefined && (obj.unitratio = message.unitratio);
    message.settledate !== undefined && (obj.settledate = message.settledate);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRequestExchange>, I>>(object: I): MsgRequestExchange {
    const message = createBaseMsgRequestExchange();
    message.creator = object.creator ?? "";
    message.finalunit = object.finalunit ?? "";
    message.amount = object.amount ?? "";
    message.startunit = object.startunit ?? "";
    message.unitratio = object.unitratio ?? "";
    message.settledate = object.settledate ?? "";
    return message;
  },
};

function createBaseMsgRequestExchangeResponse(): MsgRequestExchangeResponse {
  return {};
}

export const MsgRequestExchangeResponse = {
  encode(_: MsgRequestExchangeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRequestExchangeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRequestExchangeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRequestExchangeResponse {
    return {};
  },

  toJSON(_: MsgRequestExchangeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRequestExchangeResponse>, I>>(_: I): MsgRequestExchangeResponse {
    const message = createBaseMsgRequestExchangeResponse();
    return message;
  },
};

function createBaseMsgApproveExchange(): MsgApproveExchange {
  return { creator: "", id: 0 };
}

export const MsgApproveExchange = {
  encode(message: MsgApproveExchange, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApproveExchange {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApproveExchange();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgApproveExchange {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgApproveExchange): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApproveExchange>, I>>(object: I): MsgApproveExchange {
    const message = createBaseMsgApproveExchange();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgApproveExchangeResponse(): MsgApproveExchangeResponse {
  return {};
}

export const MsgApproveExchangeResponse = {
  encode(_: MsgApproveExchangeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApproveExchangeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApproveExchangeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgApproveExchangeResponse {
    return {};
  },

  toJSON(_: MsgApproveExchangeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApproveExchangeResponse>, I>>(_: I): MsgApproveExchangeResponse {
    const message = createBaseMsgApproveExchangeResponse();
    return message;
  },
};

function createBaseMsgSettleExchange(): MsgSettleExchange {
  return { creator: "", id: 0 };
}

export const MsgSettleExchange = {
  encode(message: MsgSettleExchange, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSettleExchange {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSettleExchange();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSettleExchange {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgSettleExchange): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSettleExchange>, I>>(object: I): MsgSettleExchange {
    const message = createBaseMsgSettleExchange();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgSettleExchangeResponse(): MsgSettleExchangeResponse {
  return {};
}

export const MsgSettleExchangeResponse = {
  encode(_: MsgSettleExchangeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSettleExchangeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSettleExchangeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSettleExchangeResponse {
    return {};
  },

  toJSON(_: MsgSettleExchangeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSettleExchangeResponse>, I>>(_: I): MsgSettleExchangeResponse {
    const message = createBaseMsgSettleExchangeResponse();
    return message;
  },
};

function createBaseMsgSwapExchange(): MsgSwapExchange {
  return { creator: "", id: 0 };
}

export const MsgSwapExchange = {
  encode(message: MsgSwapExchange, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSwapExchange {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSwapExchange();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSwapExchange {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgSwapExchange): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSwapExchange>, I>>(object: I): MsgSwapExchange {
    const message = createBaseMsgSwapExchange();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgSwapExchangeResponse(): MsgSwapExchangeResponse {
  return {};
}

export const MsgSwapExchangeResponse = {
  encode(_: MsgSwapExchangeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSwapExchangeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSwapExchangeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSwapExchangeResponse {
    return {};
  },

  toJSON(_: MsgSwapExchangeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSwapExchangeResponse>, I>>(_: I): MsgSwapExchangeResponse {
    const message = createBaseMsgSwapExchangeResponse();
    return message;
  },
};

function createBaseMsgCancelExchange(): MsgCancelExchange {
  return { creator: "", id: 0 };
}

export const MsgCancelExchange = {
  encode(message: MsgCancelExchange, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelExchange {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelExchange();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancelExchange {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgCancelExchange): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancelExchange>, I>>(object: I): MsgCancelExchange {
    const message = createBaseMsgCancelExchange();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgCancelExchangeResponse(): MsgCancelExchangeResponse {
  return {};
}

export const MsgCancelExchangeResponse = {
  encode(_: MsgCancelExchangeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelExchangeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelExchangeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCancelExchangeResponse {
    return {};
  },

  toJSON(_: MsgCancelExchangeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancelExchangeResponse>, I>>(_: I): MsgCancelExchangeResponse {
    const message = createBaseMsgCancelExchangeResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  RequestExchange(request: MsgRequestExchange): Promise<MsgRequestExchangeResponse>;
  ApproveExchange(request: MsgApproveExchange): Promise<MsgApproveExchangeResponse>;
  SettleExchange(request: MsgSettleExchange): Promise<MsgSettleExchangeResponse>;
  SwapExchange(request: MsgSwapExchange): Promise<MsgSwapExchangeResponse>;
  CancelExchange(request: MsgCancelExchange): Promise<MsgCancelExchangeResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RequestExchange = this.RequestExchange.bind(this);
    this.ApproveExchange = this.ApproveExchange.bind(this);
    this.SettleExchange = this.SettleExchange.bind(this);
    this.SwapExchange = this.SwapExchange.bind(this);
    this.CancelExchange = this.CancelExchange.bind(this);
  }
  RequestExchange(request: MsgRequestExchange): Promise<MsgRequestExchangeResponse> {
    const data = MsgRequestExchange.encode(request).finish();
    const promise = this.rpc.request("fractal.fractal.Msg", "RequestExchange", data);
    return promise.then((data) => MsgRequestExchangeResponse.decode(new _m0.Reader(data)));
  }

  ApproveExchange(request: MsgApproveExchange): Promise<MsgApproveExchangeResponse> {
    const data = MsgApproveExchange.encode(request).finish();
    const promise = this.rpc.request("fractal.fractal.Msg", "ApproveExchange", data);
    return promise.then((data) => MsgApproveExchangeResponse.decode(new _m0.Reader(data)));
  }

  SettleExchange(request: MsgSettleExchange): Promise<MsgSettleExchangeResponse> {
    const data = MsgSettleExchange.encode(request).finish();
    const promise = this.rpc.request("fractal.fractal.Msg", "SettleExchange", data);
    return promise.then((data) => MsgSettleExchangeResponse.decode(new _m0.Reader(data)));
  }

  SwapExchange(request: MsgSwapExchange): Promise<MsgSwapExchangeResponse> {
    const data = MsgSwapExchange.encode(request).finish();
    const promise = this.rpc.request("fractal.fractal.Msg", "SwapExchange", data);
    return promise.then((data) => MsgSwapExchangeResponse.decode(new _m0.Reader(data)));
  }

  CancelExchange(request: MsgCancelExchange): Promise<MsgCancelExchangeResponse> {
    const data = MsgCancelExchange.encode(request).finish();
    const promise = this.rpc.request("fractal.fractal.Msg", "CancelExchange", data);
    return promise.then((data) => MsgCancelExchangeResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
