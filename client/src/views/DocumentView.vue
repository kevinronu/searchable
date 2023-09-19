<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";

import DocumentComponent from "../components/DocumentComponent.vue";
import { Document, SearchResult } from "../models/document.model.ts";
import { useSearchResultStore } from "../stores/SearchResultStore";
import ArrowLeftIcon from "../components/icons/ArrowLeftIcon.vue";
import { getDocument } from "../services/documents-service";

const route = useRoute();
const searchResult = useSearchResultStore();
const document = ref<Document>();

onMounted(() => {
  if (route.params.documentId != "") {
    if (searchResult.checkIfEmpty) {
      getDocument(String(route.params.documentId))
        .then((data: SearchResult) => {
          searchResult.updateSearchResult(data);
          document.value = searchResult.getDocumentById(
            String(route.params.documentId)
          );
        })
        .catch((error) => {
          console.error(error);
        });
    } else {
      document.value = searchResult.getDocumentById(
        String(route.params.documentId)
      );
    }
  }
});
</script>

<template>
  <main class="container m-auto p-4 grid place-items-center" v-if="document">
    <DocumentComponent
      :subject="document.subject"
      :body="document.body"
      :from="document.from"
      :to="document.to"
      :cc="document.cc"
      :date="document.date"
    />
    <a
      href="javascript:history.back()"
      class="grid place-items-center text-gray-200 bg-stone-100 dark:bg-stone-800 hover:bg-pink-700 font-semibold py-2 px-4 rounded-lg shadow-md transition duration-300 ease-in-out"
    >
      <ArrowLeftIcon class="h-6 w-6" />
    </a>
  </main>
</template>
