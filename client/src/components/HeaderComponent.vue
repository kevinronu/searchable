<script setup lang="ts">
import { useThemeStore } from "../stores/ThemeStore"; // Adjust the import path accordingly

const themeStore = useThemeStore();
</script>

<template>
  <header class="text-center p-4 bg-stone-300 dark:bg-stone-800 mb-6">
    <div
      class="container m-auto flex flex-wrap gap-4 items-center justify-between"
    >
      <router-link
        to="/"
        class="cursor-pointer focus:outline-none focus:ring focus:ring-blue-400 focus:rounded"
      >
        <h1
          class="text-gray-700 dark:text-gray-200 text-2xl font-medium hover:bg-pink-400 dark:hover:bg-pink-600"
        >
          { Search<span class="text-2xl text-pink-600">able</span> }
        </h1>
      </router-link>
      <label class="relative text-[1em] inline-block w-[3.5em] h-[2em]">
        <input
          type="checkbox"
          class="peer w-0 h-0 opacity-0"
          :checked="themeStore.isDarkMode"
          @change="themeStore.toggleTheme"
        />
        <span
          class="absolute cursor-pointer inset-0 bg-sky-500 transition duration-400 rounded-[1.875rem] before:absolute before:content-[''] before:h-[1.4em] before:w-[1.4rem] before:rounded-[1.25rem] before:left-[0.3rem] before:bottom-[0.3em] before:bg-yellow-300 before:transition before:duration-400 peer-checked:before:bg-black peer-checked:before:rounded-[3.125rem] peer-checked:before:shadow-moon peer-checked:bg-black peer-checked:before:translate-x-[1.5em]"
        />
      </label>
    </div>
  </header>
</template>
