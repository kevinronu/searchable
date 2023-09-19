<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";

import ArrowLeftIcon from "../components/icons/ArrowLeftIcon.vue";
import MailFromIcon from "../components/icons/MailFromIcon.vue";
import CalendarIcon from "../components/icons/CalendarIcon.vue";
import MailToIcon from "../components/icons/MailToIcon.vue";
import { Document, SearchResult } from "../models/document.model.ts";
import { useSearchResultStore } from "../stores/SearchResultStore";
import { getDocument } from "../services/documents-service";
import { parseDate } from "../utils/utils";

const route = useRoute();
const searchResult = useSearchResultStore();
const document = ref<Document>();

const options: Intl.DateTimeFormatOptions = {
  weekday: "long",
  year: "numeric",
  month: "long",
  day: "numeric",
  timeZone: "UTC",
  hour: "2-digit",
  minute: "2-digit",
};

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
    <div class="bg-stone-100 dark:bg-stone-800 shadow-lg rounded-lg p-4 mb-4">
      <p class="text-xl font-semibold mb-4 text-center">
        {{ document.subject }}
      </p>
      <p class="text-lg mb-4 break-all">
        {{ document.body }}
      </p>
      <div class="flex flex-nowrap items-center gap-2">
        <MailFromIcon class="w-6 h-6 fill-red-500" />
        <span class="font-medium text-cyan-700 dark:text-cyan-500">
          From:
        </span>
      </div>
      <p class="mb-4">
        {{ document.from }}
      </p>
      <div class="flex flex-nowrap items-center gap-2">
        <MailToIcon class="w-6 h-6" />
        <span class="font-medium text-cyan-700 dark:text-cyan-500">To:</span>
      </div>
      <p class="mb-4">
        {{ document.to ? document.to.join(", ") : "" }}
      </p>
      <div class="flex flex-nowrap items-center gap-2">
        <MailToIcon class="w-6 h-6" />
        <span class="font-medium text-cyan-700 dark:text-cyan-500">CC:</span>
      </div>
      <p class="mb-4">
        {{ document.cc ? document.cc.join(", ") : "" }}
      </p>
      <div class="flex flex-nowrap items-center gap-2">
        <CalendarIcon class="w-6 h-6" />
        <span class="font-medium text-cyan-700 dark:text-cyan-500">Date:</span>
      </div>
      <p class="mb-4">
        {{ parseDate(document.date, options) }}
      </p>
    </div>
    <a
      href="javascript:history.back()"
      class="grid place-items-center text-gray-200 bg-stone-100 dark:bg-stone-800 hover:bg-pink-700 font-semibold py-2 px-4 rounded-lg shadow-md transition duration-300 ease-in-out"
    >
      <ArrowLeftIcon class="h-6 w-6" />
    </a>
  </main>
</template>
