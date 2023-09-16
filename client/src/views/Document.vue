<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";

import { Hit } from "../models/document.model.ts";
import { useDocumentsStore } from "../stores/DocumentsStore";
import MailToIcon from "../components/icons/MailToIcon.vue";
import MailFromIcon from "../components/icons/MailFromIcon.vue";
import CalendarIcon from "../components/icons/CalendarIcon.vue";

const route = useRoute();
const documents = useDocumentsStore();
const document = ref<Hit>();

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

onMounted(() => {
  if (route.params.documentId != "") {
    document.value = documents.getHitById(String(route.params.documentId));
  }
});
</script>

<template>
  <main class="container m-auto" v-if="document">
    <div class="bg-stone-100 dark:bg-stone-800 shadow-lg rounded-lg p-4">
      <p class="text-xl font-semibold mb-4 text-center">
        {{ document._source.subject }}
      </p>
      <p class="text-lg mb-4">
        {{ document._source.body }}
      </p>
      <div class="flex flex-nowrap items-center gap-2">
        <MailFromIcon class="w-6 h-6 fill-red-500" />
        <span class="font-medium text-cyan-700 dark:text-cyan-500">
          From:
        </span>
      </div>
      <p class="mb-4">
        {{ document._source.from }}
      </p>
      <div class="flex flex-nowrap items-center gap-2">
        <MailToIcon class="w-6 h-6" />
        <span class="font-medium text-cyan-700 dark:text-cyan-500">To:</span>
      </div>
      <p class="mb-4">
        {{ document._source.to ? document._source.to.join(", ") : "" }}
      </p>
      <div class="flex flex-nowrap items-center gap-2">
        <MailToIcon class="w-6 h-6" />
        <span class="font-medium text-cyan-700 dark:text-cyan-500">CC:</span>
      </div>
      <p class="mb-4">
        {{ document._source.cc ? document._source.cc.join(", ") : "" }}
      </p>
      <div class="flex flex-nowrap items-center gap-2">
        <CalendarIcon class="w-6 h-6" />
        <span class="font-medium text-cyan-700 dark:text-cyan-500">Date:</span>
      </div>
      <p class="mb-4">
        {{ parseDate(document._source.date) }}
      </p>
    </div>
  </main>
</template>
