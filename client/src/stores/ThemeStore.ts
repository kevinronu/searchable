import { defineStore } from "pinia";

export const useThemeStore = defineStore("themeStore", {
  state: () => ({
    theme: localStorage.theme === "dark" || false,
  }),
  getters: {
    isDarkMode(state) {
      return state.theme;
    },
  },
  actions: {
    toggleTheme() {
      this.theme = !this.theme;
      localStorage.theme = this.theme ? "dark" : "light";
    },
  },
});
