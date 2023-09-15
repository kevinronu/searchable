<script setup lang="ts">
import SearchFolderIconVue from "./icons/SearchFolderIcon.vue";
</script>

<template>
  <router-link to="/">
    <header class="text-center p-4 bg-stone-400 dark:bg-stone-900">
      <div
        class="container m-auto flex flex-wrap gap-4 items-center justify-between"
      >
        <h1 class="text-gray-700 dark:text-gray-200 text-2xl font-medium">
          { Search<span class="text-2xl text-pink-600">able</span> }
        </h1>
        <SearchFolderIconVue class="w-12 h-12 dark:bg-pink-600 rounded" />
      </div>
    </header>
  </router-link>
</template>
