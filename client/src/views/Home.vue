<script setup lang="ts">
import { ref } from "vue";

import router from "../router/index.ts";
import SearchIcon from "../components/icons/SearchIcon.vue";
import SearchFileIcon from "../components/icons/SearchFileIcon.vue";

const query = ref("");
const submit = () => {
  if (query.value == "") return;
  router.push({ path: `/search/${encodeURIComponent(query.value)}/page/1` });
};
</script>

<template>
  <main class="container m-auto flex-auto grid place-items-center">
    <div class="w-full max-w-xl px-4">
      <div class="flex flex-nowrap items-center justify-center gap-2 mb-2">
        <SearchFileIcon class="w-10 h-10" />
        <h1 class="text-3xl text-pink-600">Search by Words</h1>
      </div>
      <form @submit.prevent="submit">
        <div
          class="flex items-center bg-stone-50 dark:bg-stone-800 dark:placeholder:text-gray-400 outline-none ring ring-pink-600 rounded"
        >
          <input
            v-model="query"
            type="text"
            name="term"
            placeholder="+marketing +campaign -action"
            class="appearance-none px-4 py-2 bg-stone-50 dark:bg-stone-800 pr-2 w-full rounded focus:outline-none focus:ring focus:ring-blue-400 focus:rounded placeholder:text-gray-500 placeholder:dark:text-gray-400"
          />
          <button
            type="submit"
            class="ml-[3px] rounded bg-[#2CBCB2] focus:outline-none focus:ring focus:ring-blue-400 focus:rounded p-2"
          >
            <SearchIcon class="w-6 h-6" />
          </button>
        </div>
      </form>
      <p class="mt-2 text-center">
        Instructions:
        <span class="text-green-500">+word</span> (include),
        <span class="text-red-500">-word</span> (exclude)
      </p>
    </div>
  </main>
</template>
