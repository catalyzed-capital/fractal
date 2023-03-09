/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "fractal.balancecheck";

export interface MsgUsdcBalanceCheck {
  creator: string;
  chain: string;
  chainaddress: string;
  amount: string;
}

export interface MsgUsdcBalanceCheckResponse {
  id: number;
}

function createBaseMsgUsdcBalanceCheck(): MsgUsdcBalanceCheck {
  return { creator: "", chain: "", chainaddress: "", amount: "" };
}

export const MsgUsdcBalanceCheck = {
  encode(message: MsgUsdcBalanceCheck, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.chain !== "") {
      writer.uint32(18).string(message.chain);
    }
    if (message.chainaddress !== "") {
      writer.uint32(26).string(message.chainaddress);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUsdcBalanceCheck {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUsdcBalanceCheck();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.chain = reader.string();
          break;
        case 3:
          message.chainaddress = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUsdcBalanceCheck {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
      chainaddress: isSet(object.chainaddress) ? String(object.chainaddress) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: MsgUsdcBalanceCheck): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.chain !== undefined && (obj.chain = message.chain);
    message.chainaddress !== undefined && (obj.chainaddress = message.chainaddress);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUsdcBalanceCheck>, I>>(object: I): MsgUsdcBalanceCheck {
    const message = createBaseMsgUsdcBalanceCheck();
    message.creator = object.creator ?? "";
    message.chain = object.chain ?? "";
    message.chainaddress = object.chainaddress ?? "";
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseMsgUsdcBalanceCheckResponse(): MsgUsdcBalanceCheckResponse {
  return { id: 0 };
}

export const MsgUsdcBalanceCheckResponse = {
  encode(message: MsgUsdcBalanceCheckResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUsdcBalanceCheckResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUsdcBalanceCheckResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUsdcBalanceCheckResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgUsdcBalanceCheckResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUsdcBalanceCheckResponse>, I>>(object: I): MsgUsdcBalanceCheckResponse {
    const message = createBaseMsgUsdcBalanceCheckResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  UsdcBalanceCheck(request: MsgUsdcBalanceCheck): Promise<MsgUsdcBalanceCheckResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.UsdcBalanceCheck = this.UsdcBalanceCheck.bind(this);
  }
  UsdcBalanceCheck(request: MsgUsdcBalanceCheck): Promise<MsgUsdcBalanceCheckResponse> {
    const data = MsgUsdcBalanceCheck.encode(request).finish();
    const promise = this.rpc.request("fractal.balancecheck.Msg", "UsdcBalanceCheck", data);
    return promise.then((data) => MsgUsdcBalanceCheckResponse.decode(new _m0.Reader(data)));
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
