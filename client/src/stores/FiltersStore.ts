import { defineStore } from "pinia";
import { Filters } from "../models/document.model";

function initializeFilters(): Filters {
  const storedFilters = JSON.parse(localStorage.getItem("filters") || "{}");
  if (Object.keys(storedFilters).length === 0) {
    return { sort: "-date", from: null, to: null };
  }
  return storedFilters;
}

export const useFiltersStore = defineStore("filtersStore", {
  state: () => ({
    filters: initializeFilters(),
  }),
  getters: {
    getSort(state) {
      return state.filters.sort;
    },
    getFrom(state) {
      return state.filters.from;
    },
    getTo(state) {
      return state.filters.to;
    },
  },
  actions: {
    setSort(sort: string) {
      this.filters.sort = sort;
      localStorage.setItem("filters", JSON.stringify(this.filters));
    },
    setFrom(from: string) {
      this.filters.from = from;
      localStorage.setItem("filters", JSON.stringify(this.filters));
    },
    setTo(to: string) {
      this.filters.to = to;
      localStorage.setItem("filters", JSON.stringify(this.filters));
    },
    clearFilters() {
      this.filters = { sort: "-date", from: null, to: null };
      localStorage.setItem("filters", JSON.stringify(this.filters));
    },
  },
});
