<script setup lang="ts">
import GithubIcon from "./icons/GithubIcon.vue";
</script>

<template>
  <footer class="p-4 bg-stone-400 dark:bg-stone-900">
    <div class="container m-auto flex flex-wrap justify-between items-center">
      <p>©2023 - Searchable</p>
      <a
        href="https://github.com/kevinronu"
        target="_blank"
        class="flex flex-wrap gap-1 cursor-pointer hover:bg-pink-600 focus:outline-none focus:ring focus:ring-blue-400 focus:rounded"
      >
        <GithubIcon class="w-6 h-6" />
        Kevin Robles
      </a>
    </div>
  </footer>
</template>
