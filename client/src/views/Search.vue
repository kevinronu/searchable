<script setup lang="ts">
import { onMounted } from "vue";
import { useRoute, onBeforeRouteUpdate } from "vue-router";

import { useSearchResultStoreStore } from "../stores/SearchResultStore";
import { SearchResult, Body } from "../models/document.model.ts";
import MailToIcon from "../components/icons/MailToIcon.vue";
import MailFromIcon from "../components/icons/MailFromIcon.vue";
import CalendarIcon from "../components/icons/CalendarIcon.vue";
import Pagination from "../components/Pagination.vue";
const route = useRoute();

const searchResult = useSearchResultStoreStore();

const indexName = import.meta.env.VITE_INDEX_NAME;
const BASE_URI = import.meta.env.VITE_BASE_URI;

function parseDate(date: Date | string): string {
  if (typeof date === "string") {
    date = new Date(date);
  }

  if (date instanceof Date && !isNaN(date.getTime())) {
    const options: Intl.DateTimeFormatOptions = {
      year: "numeric",
      month: "numeric",
      day: "numeric",
      timeZone: "UTC",
    };
    return date.toLocaleDateString("en-US", options);
  } else {
    return "Invalid Date";
  }
}

async function fetchEmails(query: string, currentPage: number) {
  const body: Body = {
    query,
    sort: "-date",
    pagination: {
      from: 20 * (currentPage - 1),
      size: 20,
    },
  };

  fetch(`${BASE_URI}/v1/${indexName}/search`, {
    method: "POST",
    body: JSON.stringify(body),
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((resp) => resp.json())
    .then((data: SearchResult) => {
      searchResult.updateSearchResult(data);
    })
    .catch((error) => console.log(error));
}

onMounted(() => {
  if (route.params.query != "") {
    fetchEmails(String(route.params.query), Number(route.params.currentPage));
  }
});

onBeforeRouteUpdate(async (to, from) => {
  if (
    to.params.query !== from.params.query ||
    to.params.currentPage !== from.params.currentPage
  ) {
    fetchEmails(String(to.params.query), Number(to.params.currentPage));
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
        <div
          class="bg-zinc-200 dark:bg-zinc-700 hover:bg-pink-400 dark:hover:bg-pink-600 border-2 border-pink-600 p-1"
        >
          <p class="text-xl font-semibold truncate">
            {{ document.subject }}
          </p>
          <p class="text-lg line-clamp-3">
            {{ document.body }}
          </p>
          <div class="flex flex-nowrap items-center gap-1">
            <MailFromIcon class="w-6 h-6" />
            <p class="h-5 truncate">From: {{ document.from }}</p>
          </div>
          <div class="flex flex-nowrap items-center gap-1">
            <MailToIcon class="w-6 h-6" />
            <p class="h-5 truncate">
              To: {{ document.to ? document.to[0] : "" }}
            </p>
          </div>
          <div class="flex flex-nowrap items-center gap-1">
            <CalendarIcon class="w-6 h-6" />
            <p class="h-5 truncate">
              {{ parseDate(document.date) }}
            </p>
          </div>
        </div>
      </router-link>
    </div>
    <Pagination
      v-if="searchResult"
      :location="`/search/${String(route.params.query)}`"
      :current-page="Number(route.params.currentPage)"
      :total-quantity="searchResult.getTotal"
      :quantity-per-page="20"
      :max-pages-generated="10"
    />
  </main>
</template>
