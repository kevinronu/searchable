<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";

import { Document, SearchResult } from "../models/document.model.ts";
import { useSearchResultStoreStore } from "../stores/SearchResultStore";
import ArrowLeftIcon from "../components/icons/ArrowLeftIcon.vue";
import MailToIcon from "../components/icons/MailToIcon.vue";
import MailFromIcon from "../components/icons/MailFromIcon.vue";
import CalendarIcon from "../components/icons/CalendarIcon.vue";

const indexName = import.meta.env.VITE_INDEX_NAME;
const BASE_URI = import.meta.env.VITE_BASE_URI;
const route = useRoute();
const searchResult = useSearchResultStoreStore();
const document = ref<Document>();

function parseDate(date: Date | string): string {
  if (typeof date === "string") {
    date = new Date(date);
  }

  if (date instanceof Date && !isNaN(date.getTime())) {
    const options: Intl.DateTimeFormatOptions = {
      weekday: "long",
      year: "numeric",
      month: "long",
      day: "numeric",
      timeZone: "UTC",
      hour: "2-digit",
      minute: "2-digit",
    };
    return date.toLocaleDateString("en-US", options);
  } else {
    return "Invalid Date";
  }
}

// async function fetchEmails(id: string) {
//   // console.log("fetch");
//   fetch(`${BASE_URI}/v1/${indexName}/${id}`)
//     .then((resp) => resp.json())
//     .then((data: SearchResult) => {
//       // console.log("data", data);
//       // searchResult.updateSearchResult(data);
//       return data;
//     })
//     .catch((error) => console.log(error));
// }

async function fetchEmails(id: string): Promise<SearchResult> {
  try {
    const response = await fetch(`${BASE_URI}/v1/${indexName}/${id}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch data for document with ID ${id}`);
    }

    const data: SearchResult = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error; // You can choose to handle the error differently here
  }
}

onMounted(async () => {
  if (route.params.documentId != "") {
    if (searchResult.checkIfEmpty) {
      fetchEmails(String(route.params.documentId))
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
      <p class="text-lg mb-4">
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
        {{ parseDate(document.date) }}
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
