<script setup lang="ts">
import { ref, watch, onMounted } from "vue";
import { useRoute, onBeforeRouteUpdate } from "vue-router";

import LoadingComponent from "../components/LoadingComponent.vue";
import FiltersComponent from "../components/FiltersComponent.vue";
import NoResultsComponent from "../components/NoResultsComponent.vue";
import DocumentPreviewComponent from "../components/DocumentPreviewComponent.vue";
import PaginationComponent from "../components/PaginationComponent.vue";
import { useSearchResultStore } from "../stores/SearchResultStore";
import { searchDocument } from "../services/documents-service";
import { SearchResult } from "../models/document.model.ts";
import { useFiltersStore } from "../stores/FiltersStore";

const resultsPerPage = Number(import.meta.env.VITE_RESULTS_PER_PAGE);
const maximumPagesGenerated = Number(
  import.meta.env.VITE_MAXIMUM_PAGES_GENERATED
);
const route = useRoute();

const searchResult = useSearchResultStore();
const filters = useFiltersStore();

const loading = ref(false);

watch(
  [() => filters.getSort, () => filters.getFrom, () => filters.getTo],
  ([newSort, newFrom, newTo]) => {
    loading.value = true;
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
      })
      .finally(() => (loading.value = false));
  }
);

onMounted(() => {
  if (route.params.query != "") {
    loading.value = true;
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
      })
      .finally(() => (loading.value = false));
  }
});

onBeforeRouteUpdate(async (to, from) => {
  if (
    to.params.query !== from.params.query ||
    to.params.currentPage !== from.params.currentPage
  ) {
    loading.value = true;
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
      })
      .finally(() => (loading.value = false));
  }
});
</script>
<template>
  <main class="flex-auto flex flex-col justify-between">
    <FiltersComponent />
    <div
      class="container m-auto flex-auto grid grid-cols-1 gap-4 md:grid-cols-2 xl:grid-cols-4 p-4"
      v-if="searchResult.checkIfHaveDocuments && !loading"
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
    <LoadingComponent v-if="loading" />
    <NoResultsComponent v-if="!searchResult.checkIfHaveDocuments && !loading" />
    <PaginationComponent
      v-if="searchResult.checkIfHaveDocuments && !loading"
      :location="`/search/${String(route.params.query)}`"
      :current-page="Number(route.params.currentPage)"
      :total-quantity="searchResult.getTotal"
      :quantity-per-page="resultsPerPage"
      :max-pages-generated="maximumPagesGenerated"
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
