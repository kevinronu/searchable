<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";

import LoadingComponent from "../components/LoadingComponent.vue";
import DocumentComponent from "../components/DocumentComponent.vue";
import { Document, SearchResult } from "../models/document.model.ts";
import { useSearchResultStore } from "../stores/SearchResultStore";
import ArrowLeftIcon from "../components/icons/ArrowLeftIcon.vue";
import { getDocument } from "../services/documents-service";
import NoResultsComponent from "../components/NoResultsComponent.vue";

const route = useRoute();
const searchResult = useSearchResultStore();
const document = ref<Document>();
const fetched = ref(false);
const loading = ref(false);

onMounted(() => {
  if (route.params.documentId != "") {
    if (!searchResult.checkIfHaveDocuments && !fetched.value) {
      loading.value = true;
      getDocument(String(route.params.documentId))
        .then((data: SearchResult) => {
          searchResult.updateSearchResult(data);
          document.value = searchResult.getDocumentById(
            String(route.params.documentId)
          );
        })
        .catch((error) => {
          console.error(error);
        })
        .finally(() => {
          loading.value = true;
          fetched.value = true;
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
  <main class="flex-auto container m-auto p-4 flex flex-col justify-between">
    <DocumentComponent
      v-if="document"
      :subject="document.subject"
      :body="document.body"
      :from="document.from"
      :to="document.to"
      :cc="document.cc"
      :date="document.date"
    />
    <LoadingComponent v-if="loading" />
    <NoResultsComponent v-if="!document && !loading" />
    <div class="text-center flex items-center justify-center">
      <a
        href="javascript:history.back()"
        class="py-2 px-4 rounded-lg hover:bg-pink-700 hover:dark:bg-pink-700 bg-stone-100 dark:bg-stone-800 cursor-pointer shadow-md transition duration-300 ease-in-out"
      >
        <ArrowLeftIcon class="h-6 w-6" />
      </a>
    </div>
  </main>
</template>
