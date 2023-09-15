import { defineStore } from "pinia";

import { Hit } from "../models/document.model";

export const useDocumentsStore = defineStore("documentsStore", {
  state: () => ({
    documents: { fetched: false, hits: [] as Hit[] },
  }),
  getters: {
    getHitById(state) {
      const hits = state.documents.hits;
      return (id: string) => hits.find((hit) => hit._id === id);
    },
    getFetched(state) {
      return state.documents.fetched;
    },
    getHits(state) {
      return state.documents.hits;
    },
  },
  actions: {
    updateHits(hits: Hit[]) {
      this.documents.hits = hits;
    },
    toggleFetched() {
      const bool = this.documents.fetched;
      this.documents.fetched = !bool;
    },
  },
});
