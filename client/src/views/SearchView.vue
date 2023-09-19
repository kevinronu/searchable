<script setup lang="ts">
import { onMounted } from "vue";
import { useRoute, onBeforeRouteUpdate } from "vue-router";

import DocumentPreviewComponent from "../components/DocumentPreviewComponent.vue";
import PaginationComponent from "../components/PaginationComponent.vue";
import { useSearchResultStore } from "../stores/SearchResultStore";
import { searchDocument } from "../services/documents-service";
import { SearchResult } from "../models/document.model.ts";
const route = useRoute();

const searchResult = useSearchResultStore();

onMounted(() => {
  if (route.params.query != "") {
    searchDocument(String(route.params.query), Number(route.params.currentPage))
      .then((data: SearchResult) => {
        searchResult.updateSearchResult(data);
      })
      .catch((error) => console.log(error));
  }
});

onBeforeRouteUpdate(async (to, from) => {
  if (
    to.params.query !== from.params.query ||
    to.params.currentPage !== from.params.currentPage
  ) {
    searchDocument(String(to.params.query), Number(to.params.currentPage))
      .then((data: SearchResult) => {
        searchResult.updateSearchResult(data);
      })
      .catch((error) => console.log(error));
  }
});
</script>
<template>
  <main>
    <div
      class="container m-auto grid grid-cols-1 gap-4 md:grid-cols-2 xl:grid-cols-4 p-4"
    >
      <router-link
        v-for="document in searchResult?.getDocuments"
        :key="document.id"
        :to="`/document/${document.id}`"
        class="focus:outline-none focus:ring focus:ring-blue-400 focus:rounded"
      >
        <DocumentPreviewComponent
          :subject="document.subject"
          :body="document.body"
          :from="document.from"
          :to="document.to"
          :cc="document.cc"
          :date="document.date"
        />
      </router-link>
    </div>
    <PaginationComponent
      v-if="searchResult"
      :location="`/search/${String(route.params.query)}`"
      :current-page="Number(route.params.currentPage)"
      :total-quantity="searchResult.getTotal"
      :quantity-per-page="20"
      :max-pages-generated="10"
    />
  </main>
</template>
