<script setup lang="ts">
import { useRoute, onBeforeRouteUpdate } from "vue-router";
import { onMounted, ref } from "vue";
import { Documents, Hits, Body } from "../models/document.model.ts";
import { BASE_URI } from "../config.ts";
import MailToIcon from "../components/icons/MailToIcon.vue";
import MailFromIcon from "../components/icons/MailFromIcon.vue";
import CalendarIcon from "../components/icons/CalendarIcon.vue";
import Pagination from "../components/Pagination.vue";
const query = String(useRoute().params.query);
const currentPage = Number(useRoute().params.currentPage);
const searchResult = ref<Hits>();
const totalQuantity = ref<number>(0);

const indexName = import.meta.env.VITE_INDEX_NAME;

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

// async function fetchEmails() {
//   const body: Body = {
//     query: query,
//     sort: "-date",
//     pagination: {
//       from: 20 * (currentPage - 1),
//       size: 20,
//     },
//   };

//   fetch(`${BASE_URI}/v1/${indexName}/search`, {
//     method: "POST",
//     body: JSON.stringify(body),
//     headers: {
//       "Content-Type": "application/json",
//     },
//   })
//     .then((resp) => resp.json())
//     .then((data: Documents) => {
//       totalQuantity.value = data.hits.total.value;
//       searchResult.value = data.hits;
//       return data;
//     })
//     .catch((error) => console.log(error));
// }

async function fetchEmails(): Promise<Hits> {
  const body: Body = {
    query: query,
    sort: "-date",
    pagination: {
      from: 20 * (currentPage - 1),
      size: 20,
    },
  };

  try {
    const resp = await fetch(`${BASE_URI}/v1/${indexName}/search`, {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (resp.ok) {
      const data: Documents = await resp.json();
      totalQuantity.value = data.hits.total.value;
      searchResult.value = data.hits;
      return data.hits;
    } else {
      throw new Error("Failed to fetch data");
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}

onMounted(() => {
  if (query != "") {
    fetchEmails();
  }
});

onBeforeRouteUpdate(async (to, from) => {
  if (to.params.currentPage !== from.params.currentPage) {
    console.log("Hi");
    searchResult.value = await fetchEmails();
  }
});
</script>
<template>
  <main>
    <div
      class="container m-auto grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5"
    >
      <div
        v-for="document in searchResult?.hits"
        :key="document._id"
        class="border-2 border-pink-600"
      >
        <p class="text-xl font-semibold truncate">
          {{ document._source.subject }}
        </p>
        <p class="text-lg line-clamp-4">
          {{ document._source.body }}
        </p>
        <div class="flex flex-nowrap items-center">
          <MailFromIcon class="w-6 h-6" />
          <p class="h-5 truncate">From: {{ document._source.from }}</p>
        </div>
        <div class="flex flex-nowrap items-center">
          <MailToIcon class="w-6 h-6" />
          <p class="h-5 truncate">To: {{ document._source.to }}</p>
        </div>
        <div class="flex flex-nowrap items-center">
          <CalendarIcon class="w-6 h-6" />
          <p class="h-5 truncate">
            {{ parseDate(document._source.date) }}
          </p>
        </div>
      </div>
    </div>
    <Pagination
      :location="`/search/${query}`"
      :current-page="currentPage"
      :total-quantity="totalQuantity"
      :quantity-per-page="20"
      :max-pages-generated="10"
    />
  </main>
</template>
