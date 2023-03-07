/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "fractal.fractal";

export interface Exchange {
  id: number;
  finalunit: string;
  amount: string;
  startunit: string;
  unitratio: string;
  settledate: string;
  state: string;
  entity: string;
  provider: string;
}

function createBaseExchange(): Exchange {
  return {
    id: 0,
    finalunit: "",
    amount: "",
    startunit: "",
    unitratio: "",
    settledate: "",
    state: "",
    entity: "",
    provider: "",
  };
}

export const Exchange = {
  encode(message: Exchange, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
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
    if (message.state !== "") {
      writer.uint32(58).string(message.state);
    }
    if (message.entity !== "") {
      writer.uint32(66).string(message.entity);
    }
    if (message.provider !== "") {
      writer.uint32(74).string(message.provider);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Exchange {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExchange();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
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
        case 7:
          message.state = reader.string();
          break;
        case 8:
          message.entity = reader.string();
          break;
        case 9:
          message.provider = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Exchange {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      finalunit: isSet(object.finalunit) ? String(object.finalunit) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      startunit: isSet(object.startunit) ? String(object.startunit) : "",
      unitratio: isSet(object.unitratio) ? String(object.unitratio) : "",
      settledate: isSet(object.settledate) ? String(object.settledate) : "",
      state: isSet(object.state) ? String(object.state) : "",
      entity: isSet(object.entity) ? String(object.entity) : "",
      provider: isSet(object.provider) ? String(object.provider) : "",
    };
  },

  toJSON(message: Exchange): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.finalunit !== undefined && (obj.finalunit = message.finalunit);
    message.amount !== undefined && (obj.amount = message.amount);
    message.startunit !== undefined && (obj.startunit = message.startunit);
    message.unitratio !== undefined && (obj.unitratio = message.unitratio);
    message.settledate !== undefined && (obj.settledate = message.settledate);
    message.state !== undefined && (obj.state = message.state);
    message.entity !== undefined && (obj.entity = message.entity);
    message.provider !== undefined && (obj.provider = message.provider);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Exchange>, I>>(object: I): Exchange {
    const message = createBaseExchange();
    message.id = object.id ?? 0;
    message.finalunit = object.finalunit ?? "";
    message.amount = object.amount ?? "";
    message.startunit = object.startunit ?? "";
    message.unitratio = object.unitratio ?? "";
    message.settledate = object.settledate ?? "";
    message.state = object.state ?? "";
    message.entity = object.entity ?? "";
    message.provider = object.provider ?? "";
    return message;
  },
};

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
