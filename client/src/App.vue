<script setup lang="ts">
import { onMounted, watch } from "vue";

import Footer from "./components/Footer.vue";
import Header from "./components/Header.vue";
import { useThemeStore } from "./stores/ThemeStore";

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

onMounted(() => {
  if (themeStore.isDarkMode) {
    document.documentElement.classList.add("dark");
  }
});
</script>

<template>
  <Header />
  <router-view />
  <Footer />
</template>
