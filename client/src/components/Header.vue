<script setup lang="ts">
import SearchFolderIconVue from "./icons/SearchFolderIcon.vue";
</script>

<template>
  <header class="text-center p-4 bg-stone-400 dark:bg-stone-900 mb-6">
    <div
      class="container m-auto flex flex-wrap gap-4 items-center justify-between"
    >
      <router-link to="/" class="cursor-pointer">
        <h1
          class="text-gray-700 dark:text-gray-200 text-2xl font-medium hover:bg-pink-400 dark:hover:bg-pink-600"
        >
          { Search<span class="text-2xl text-pink-600">able</span> }
        </h1>
      </router-link>
      <SearchFolderIconVue class="w-12 h-12 dark:bg-pink-600 rounded" />
    </div>
  </header>
</template>
