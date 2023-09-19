import { defineStore } from "pinia";

import { SearchResult } from "../models/document.model";

export const useSearchResultStore = defineStore("searchResultStore", {
  state: () => ({
    searchResult: JSON.parse(
      localStorage.getItem("searchResult") || "{}"
    ) as SearchResult,
  }),
  getters: {
    getDocumentById(state) {
      const documents = state.searchResult.documents;
      // if (!documents) return;
      // console.log("documents", documents);
      return (id: string) => documents.find((document) => document.id === id);
    },
    getDocuments(state) {
      return state.searchResult.documents;
    },
    getTotal(state) {
      return state.searchResult.total_found;
    },
    checkIfEmpty(state) {
      return Object.keys(state.searchResult).length === 0;
    },
  },
  actions: {
    updateSearchResult(searchResult: SearchResult) {
      this.searchResult = searchResult;
      localStorage.setItem("searchResult", JSON.stringify(searchResult));
    },
  },
});