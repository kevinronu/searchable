<script setup lang="ts">
import { watch } from "vue";
import Footer from "./components/Footer.vue";
import Header from "./components/Header.vue";
import { useThemeStore } from "./stores/ThemeStore"; // Adjust the import path accordingly

const themeStore = useThemeStore();

watch(
  () => themeStore.isDarkMode,
  (newValue, _) => {
    if (newValue) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  }
);
</script>

<template>
  <Header />
  <router-view />
  <Footer />
</template>
