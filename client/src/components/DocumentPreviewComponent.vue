<script setup lang="ts">
import MailFromIcon from "./icons/MailFromIcon.vue";
import MailToIcon from "./icons/MailToIcon.vue";
import CalendarIcon from "./icons/CalendarIcon.vue";
import { parseDate } from "../utils/utils";

const options: Intl.DateTimeFormatOptions = {
  year: "numeric",
  month: "numeric",
  day: "numeric",
  timeZone: "UTC",
};

const props = defineProps({
  subject: {
    type: String,
    required: true,
  },
  body: {
    type: String,
    required: true,
  },
  from: {
    type: String,
    required: true,
  },
  to: {
    type: [Array],
    default: [],
    required: false,
  },
  cc: {
    type: Array,
    default: [],
    required: false,
  },
  date: {
    type: [Date, String],
    required: true,
  },
});
</script>

<template>
  <div
    class="bg-zinc-200 dark:bg-zinc-700 hover:bg-pink-400 dark:hover:bg-pink-600 border-2 border-pink-600 p-1"
  >
    <p class="text-xl font-semibold truncate">
      {{ props.subject }}
    </p>
    <p class="text-lg line-clamp-3">
      {{ props.body }}
    </p>
    <div class="flex flex-nowrap items-center gap-1">
      <MailFromIcon class="w-6 h-6" />
      <p class="h-5 truncate">From: {{ props.from }}</p>
    </div>
    <div class="flex flex-nowrap items-center gap-1">
      <MailToIcon class="w-6 h-6" />
      <p class="h-5 truncate">To: {{ props.to ? props.to[0] : "" }}</p>
    </div>
    <div class="flex flex-nowrap items-center gap-1">
      <CalendarIcon class="w-6 h-6" />
      <p class="h-5 truncate">
        {{ parseDate(props.date, options) }}
      </p>
    </div>
  </div>
</template>
