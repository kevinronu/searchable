<script setup lang="ts">
import { ref } from "vue";
import router from "../router/index.ts";
import SearchIcon from "../components/icons/SearchIcon.vue";
import SearchFileIcon from "../components/icons/SearchFileIcon.vue";

const query = ref("");
const submit = () => {
  if (query.value == "") return;
  router.push({ path: `/search/${encodeURIComponent(query.value)}` });
};
</script>

<template>
  <main class="min-h-screen grid place-items-center">
    <div class="flex flex-wrap items-center justify-center gap-2">
      <SearchFileIcon class="w-16 pt-2" />
      <div>
        <h1 class="text-3xl text-pink-600 mb-2">Fast Search</h1>
        <form
          class="flex items-center justify-center gap-1 bg-[#2CBCB2] outline-none ring ring-pink-600 rounded"
          @submit.prevent="submit"
        >
          <input
            v-model="query"
            type="text"
            name="term"
            placeholder="Search by word..."
            class="rounded focus:outline-none focus:ring focus:ring-blue-400 focus:rounded placeholder:text-gray-500 bg-stone-50 dark:bg-stone-500 dark:placeholder:text-gray-400"
          />
          <button type="submit">
            <SearchIcon class="w-6 h-6" />
          </button>
        </form>
      </div>
    </div>
  </main>
</template>
