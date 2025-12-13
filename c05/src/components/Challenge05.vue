<script setup lang="ts">
import WelcomeItem from './WelcomeItem.vue'
import SupportIcon from './icons/IconSupport.vue'

import metadata from '@/metadata.ts';
import { ref } from 'vue'

import Web3 from 'web3';

let web3 = new Web3('https://sepolia.infura.io/v3/44b179a14e644182a0cfa54aa4666f05');

type DataResult = {
  address: string;
  tx_challenge01: string;
  tx_challenge02: string;
  contract_challenge03: string;
  contract_challenge04: string;
}

let contract: any = new web3.eth.Contract(metadata.abi, '0x3819C7071f2bc39C83187Bf5B5aeA79Fa3e37C42');

const getRegisteredAddressesResult = ref<string[]>([]);
const getDataResult = ref<DataResult[]>();

function sleep(ms: number) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

(async () => {
  try {
    let result: string[] = await contract.methods["getRegisteredAddresses()"]().call();
    console.log("getRegisteredAddresses", result);
    getRegisteredAddressesResult.value = result;

    (async () => {
      try {
        let results: DataResult[] = [];
        for (let address of getRegisteredAddressesResult.value) {
          let result: DataResult = await contract.methods["getData"](address).call();
          console.log("getData", result);

          let dataResult: DataResult = {
            address: address,
            tx_challenge01: result.tx_challenge01,
            tx_challenge02: result.tx_challenge02,
            contract_challenge03: result.contract_challenge03,
            contract_challenge04: result.contract_challenge04
          };

          results.push(dataResult);

          let newResults = results.slice(); // create a shallow copy

          getDataResult.value = newResults;

          await sleep(1000); // to avoid rate limiting
        }

      } catch (error) {
        console.error("Error calling getData:", error);
      }
    })();
  } catch (error) {
    console.error("Error calling getRegisteredAddresses:", error);
  }

})();



</script>

<template>
  <WelcomeItem>
    <template #icon>
      <SupportIcon />
    </template>
    <template #heading>getRegisteredAddresses</template>

    <li v-for="address in getRegisteredAddressesResult" :key="address">
      {{ address }}
    </li>
  </WelcomeItem>

  <div v-for="data in getDataResult">
    <WelcomeItem>
      <template #icon>
        <SupportIcon />
      </template>
      <template #heading>getData {{ data.address }}</template>

      <li v-for="(value, key) in data" :key="key">
        {{ key }}: {{ value }}
      </li>
    </WelcomeItem>
  </div>
</template>


<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: relative;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {

  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}
</style>