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
        <label class="text-3xl text-pink-600" for="search-input">
          Search by Words
        </label>
      </div>
      <search>
        <form @submit.prevent="submit">
          <div
            class="flex items-center bg-stone-50 dark:bg-stone-800 dark:placeholder:text-gray-400 outline-none ring ring-pink-600 rounded"
          >
            <input
              v-model="query"
              type="search"
              data-testid="search-input"
              id="search-input"
              placeholder="+marketing +campaign -action"
              class="appearance-none px-4 py-2 pr-2 w-full rounded bg-stone-50 dark:bg-stone-800 hover:outline-none hover:ring hover:ring-blue-300 hover:rounded focus:outline-none focus:ring focus:ring-blue-400 focus:rounded placeholder:text-gray-500 placeholder:dark:text-gray-400"
            />
            <button
              type="submit"
              data-testid="search-button"
              class="ml-[3px] rounded bg-[#2CBCB2] hover:outline-none hover:ring hover:ring-blue-300 hover:rounded focus:outline-none focus:ring focus:ring-blue-400 focus:rounded p-2"
            >
              <SearchIcon class="w-6 h-6" />
            </button>
          </div>
        </form>
      </search>
      <p class="mt-2 text-center">
        Instructions:
        <span class="text-green-500">+word</span> (include),
        <span class="text-red-500">-word</span> (exclude)
      </p>
    </div>
  </main>
</template>
