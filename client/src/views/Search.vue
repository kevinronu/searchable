<script setup lang="ts">
import { useRoute } from "vue-router";
import { onMounted, ref } from "vue";
import { Documents, Hit, Body } from "../models/document.model.ts";
import { BASE_URI } from "../config.ts";
import MailToIcon from "../components/icons/MailToIcon.vue";
import MailFromIcon from "../components/icons/MailFromIcon.vue";
import CalendarIcon from "../components/icons/CalendarIcon.vue";
const query = String(useRoute().params.query);
const documents = ref<Hit[]>([]);

const body: Body = {
  query,
  sort: "-date",
  pagination: {
    from: 0,
    size: 20,
  },
};

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

async function fetchEmails() {
  fetch(`${BASE_URI}/v1/${indexName}/search`, {
    method: "POST",
    body: JSON.stringify(body),
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((resp) => resp.json())
    .then((data: Documents) => (documents.value = data.hits.hits))
    .catch((error) => console.log(error));
}

function truncateTextWithSize(input: string | string[], size: number): string {
  if (typeof input === "string") {
    input = input.replace(/\n/g, " ");
    if (input.length <= size) {
      return input;
    } else {
      return input.substr(0, size) + "...";
    }
  } else if (Array.isArray(input)) {
    const concatenatedText = input.join(", ");
    if (concatenatedText.length <= size) {
      return concatenatedText;
    } else {
      return concatenatedText.substr(0, size) + "...";
    }
  } else {
    throw new Error("Invalid input type. Expected string or array of strings.");
  }
}

onMounted(() => {
  if (query != "") {
    fetchEmails();
  }
});
</script>
<template>
  <main>
    <div
      class="container m-auto grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5"
    >
      <div
        v-for="document in documents"
        :key="document._id"
        class="border-2 border-pink-600"
      >
        <p class="text-xl font-semibold">
          {{ truncateTextWithSize(document._source.subject, 35) }}
        </p>
        <p class="text-lg">
          {{ truncateTextWithSize(document._source.body, 100) }}
        </p>
        <div class="flex flex-wrap items-center">
          <MailFromIcon class="w-6 h-6" />
          <p>From: {{ truncateTextWithSize(document._source.from, 37) }}</p>
        </div>
        <div class="flex flex-wrap items-center">
          <MailToIcon class="w-6 h-6" />
          <p>To: {{ truncateTextWithSize(document._source.to, 39) }}</p>
        </div>
        <div class="flex flex-wrap items-center">
          <CalendarIcon class="w-6 h-6" />
          <p>
            {{ parseDate(document._source.date) }}
          </p>
        </div>
      </div>
    </div>
  </main>
</template>
