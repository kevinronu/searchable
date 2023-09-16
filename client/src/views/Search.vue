<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute, onBeforeRouteUpdate } from "vue-router";

import { useDocumentsStore } from "../stores/DocumentsStore";
import { Documents, Hits, Body } from "../models/document.model.ts";
import MailToIcon from "../components/icons/MailToIcon.vue";
import MailFromIcon from "../components/icons/MailFromIcon.vue";
import CalendarIcon from "../components/icons/CalendarIcon.vue";
import Pagination from "../components/Pagination.vue";
const route = useRoute();
const searchResult = ref<Hits>();
const totalQuantity = ref<number>(0);

const documents = useDocumentsStore();

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
    .then((data: Documents) => {
      totalQuantity.value = data.hits.total.value;
      searchResult.value = data.hits;
      documents.updateHits(data.hits.hits);
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
        v-for="document in searchResult?.hits"
        :key="document._id"
        :to="`/document/${document._id}`"
        class="focus:outline-none focus:ring focus:ring-blue-400 focus:rounded"
      >
        <div
          class="bg-zinc-200 dark:bg-zinc-700 hover:bg-pink-400 dark:hover:bg-pink-600 border-2 border-pink-600 p-1"
        >
          <p class="text-xl font-semibold truncate">
            {{ document._source.subject }}
          </p>
          <p class="text-lg line-clamp-3">
            {{ document._source.body }}
          </p>
          <div class="flex flex-nowrap items-center gap-1">
            <MailFromIcon class="w-6 h-6" />
            <p class="h-5 truncate">From: {{ document._source.from }}</p>
          </div>
          <div class="flex flex-nowrap items-center gap-1">
            <MailToIcon class="w-6 h-6" />
            <p class="h-5 truncate">
              To: {{ document._source.to ? document._source.to[0] : "" }}
            </p>
          </div>
          <div class="flex flex-nowrap items-center gap-1">
            <CalendarIcon class="w-6 h-6" />
            <p class="h-5 truncate">
              {{ parseDate(document._source.date) }}
            </p>
          </div>
        </div>
      </router-link>
    </div>
    <Pagination
      :location="`/search/${String(route.params.query)}`"
      :current-page="Number(route.params.currentPage)"
      :total-quantity="totalQuantity"
      :quantity-per-page="20"
      :max-pages-generated="10"
    />
  </main>
</template>
