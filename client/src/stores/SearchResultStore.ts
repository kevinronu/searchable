import { defineStore } from "pinia";

import { SearchResult } from "../models/document.model";

function initializeSearchResult(): SearchResult {
  const storedSearchResult = JSON.parse(
    localStorage.getItem("searchResult") || "{}"
  );
  if (Object.keys(storedSearchResult).length === 0) {
    return { total_found: 0, documents: [] };
  }
  return storedSearchResult;
}

export const useSearchResultStore = defineStore("searchResultStore", {
  state: () => ({
    searchResult: initializeSearchResult(),
  }),
  getters: {
    getDocumentById(state) {
      const documents = state.searchResult.documents;
      return (id: string) => documents.find((document) => document.id === id);
    },
    getDocuments(state) {
      return state.searchResult.documents;
    },
    getTotal(state) {
      return state.searchResult.total_found;
    },
    checkIfHaveDocuments(state) {
      return state.searchResult.documents.length != 0;
    },
  },
  actions: {
    updateSearchResult(searchResult: SearchResult) {
      this.searchResult = searchResult;
      localStorage.setItem("searchResult", JSON.stringify(searchResult));
    },
    clearSearchResult() {
      this.searchResult = {
        total_found: 0,
        documents: [],
      };
      localStorage.setItem("searchResult", JSON.stringify(this.searchResult));
    },
  },
});
