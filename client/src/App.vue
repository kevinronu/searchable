<script setup lang="ts">
import { onMounted, watch } from "vue";

import FooterComponent from "./components/FooterComponent.vue";
import HeaderComponent from "./components/HeaderComponent.vue";
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
  <HeaderComponent />
  <router-view />
  <FooterComponent />
</template>
