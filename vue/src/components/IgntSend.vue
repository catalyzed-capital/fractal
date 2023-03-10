<template>
  <div>
    <div class="pt-8">
      <div class="text-xs text-gray-600">Exchange Request Details</div>

      <div>
        <input
          v-model="state.tx.finalunit"
          class="mt-1 py-2 px-4 h-12 bg-gray-100 border-xs text-base leading-tight w-full rounded-xl outline-0"
          placeholder="Requested Currency (USDC, USD, or BTC)"
        />
      </div>
    </div>

    <div>
      <input
        v-model="state.tx.amount"
        class="mt-1 py-2 px-4 h-12 bg-gray-100 border-xs text-base leading-tight w-full rounded-xl outline-0"
        placeholder="Amount"
      />
    </div>

    <div>
      <input
        v-model="state.tx.startunit"
        class="mt-1 py-2 px-4 h-12 bg-gray-100 border-xs text-base leading-tight w-full rounded-xl outline-0"
        placeholder="Beginning Currency"
      />
    </div>

    <div>
      <input
        v-model="state.tx.unitratio"
        class="mt-1 py-2 px-4 h-12 bg-gray-100 border-xs text-base leading-tight w-full rounded-xl outline-0"
        placeholder="Conversion Rate for Testing"
      />
    </div>

    <div>
      <input
        v-model="state.tx.settledate"
        class="mt-1 py-2 px-4 h-12 bg-gray-100 border-xs text-base leading-tight w-full rounded-xl outline-0"
        placeholder="Select Settlement Date"
      />
    </div>

    <hr class="my-4 border-gray-300">

    <div>
      <IgntButton
        style="width: 100%"
        @click="sendTx"
        >Request Exchange</IgntButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";
import { useClient } from "@/composables/useClient";
import { IgntButton } from "@ignt/vue-library";
import type { MsgRequestExchange } from "/Users/ccteam/fractal/ts-client/fractal.fractal/types/fractal/fractal/tx";

interface Exchange {
  finalunit: string;
  amount: string;
  startunit: string;
  unitratio: string;
  settledate: string;
}

enum UI_STATE {
  "FRESH" = 1,
  "TX_SIGNING" = 300,
  "TX_SUCCESS" = 301,
  "TX_ERROR" = 302,
}

interface State {
  tx: Exchange;
  currentUIState: UI_STATE;
}

const initialState: State = {
  tx: {
    finalunit: "",
    amount: "",
    startunit: "",
    unitratio: "",
    settledate: "",
  },
  currentUIState: UI_STATE.FRESH,
};

const state = reactive(initialState);
const client = useClient();
const sendMsgRequestExchange = client.FractalFractal.tx.sendMsgRequestExchange;

const sendTx = async () => {
  const msg: MsgRequestExchange = {
    creator: "my_creator",
    finalunit: state.tx.finalunit,
    amount: state.tx.amount,
    startunit: state.tx.startunit,
    unitratio: state.tx.unitratio,
    settledate: state.tx.settledate,
  };

try {
    state.currentUIState = UI_STATE.TX_SIGNING;
    const result = await sendMsgRequestExchange({ value: msg });
    state.currentUIState = UI_STATE.TX_SUCCESS;
    console.log("Exchange request transaction sent:", result);
  } catch (error) {
    state.currentUIState = UI_STATE.TX_ERROR;
    console.error("Error submitting exchange request transaction:", error);
  }
};
</script>