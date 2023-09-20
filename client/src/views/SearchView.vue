<script setup lang="ts">
import { watch, onMounted } from "vue";
import { useRoute, onBeforeRouteUpdate } from "vue-router";

import FiltersComponent from "../components/FiltersComponent.vue";
import NoResultsComponent from "../components/NoResultsComponent.vue";
import DocumentPreviewComponent from "../components/DocumentPreviewComponent.vue";
import PaginationComponent from "../components/PaginationComponent.vue";
import { useSearchResultStore } from "../stores/SearchResultStore";
import { searchDocument } from "../services/documents-service";
import { SearchResult } from "../models/document.model.ts";
import { useFiltersStore } from "../stores/FiltersStore";

const route = useRoute();

const searchResult = useSearchResultStore();
const filters = useFiltersStore();

watch(
  [() => filters.getSort, () => filters.getFrom, () => filters.getTo],
  ([newSort, newFrom, newTo]) => {
    searchDocument(
      String(route.params.query),
      Number(route.params.currentPage),
      {
        sort: newSort,
        from: newFrom,
        to: newTo,
      }
    )
      .then((data: SearchResult) => {
        searchResult.updateSearchResult(data);
      })
      .catch((error) => {
        searchResult.clearSearchResult();
        console.log(error);
      });
  }
);

onMounted(() => {
  if (route.params.query != "") {
    searchDocument(
      String(route.params.query),
      Number(route.params.currentPage),
      {
        sort: filters.getSort,
        from: filters.getFrom,
        to: filters.getTo,
      }
    )
      .then((data: SearchResult) => {
        searchResult.updateSearchResult(data);
      })
      .catch((error) => {
        searchResult.clearSearchResult();
        console.log(error);
      });
  }
});

onBeforeRouteUpdate(async (to, from) => {
  if (
    to.params.query !== from.params.query ||
    to.params.currentPage !== from.params.currentPage
  ) {
    searchDocument(String(to.params.query), Number(to.params.currentPage), {
      sort: filters.getSort,
      from: filters.getFrom,
      to: filters.getTo,
    })
      .then((data: SearchResult) => {
        searchResult.updateSearchResult(data);
      })
      .catch((error) => {
        searchResult.clearSearchResult();
        console.log(error);
      });
  }
});
</script>
<template>
  <main class="flex-auto flex flex-col justify-between">
    <FiltersComponent />
    <div
      class="container m-auto grid grid-cols-1 gap-4 md:grid-cols-2 xl:grid-cols-4 p-4"
      v-if="searchResult.checkIfHaveDocuments"
    >
      <router-link
        v-for="document in searchResult.getDocuments"
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
    <NoResultsComponent v-if="!searchResult.checkIfHaveDocuments" />
    <PaginationComponent
      v-if="searchResult.checkIfHaveDocuments"
      :location="`/search/${String(route.params.query)}`"
      :current-page="Number(route.params.currentPage)"
      :total-quantity="searchResult.getTotal"
      :quantity-per-page="20"
      :max-pages-generated="10"
    />
    <div v-if="!searchResult.checkIfHaveDocuments" class="text-center">
      <router-link
        :to="`/search/${String(route.params.query)}/page/1`"
        class="cursor-pointer py-1 px-1 rounded"
        active-class="border-2 border-pink-600 dark:border-pink-600 bg-pink-200 dark:bg-pink-900"
      >
        1
      </router-link>
    </div>
  </main>
</template>
