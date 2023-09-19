<script setup lang="ts">
import MailFromIcon from "./icons/MailFromIcon.vue";
import MailToIcon from "./icons/MailToIcon.vue";
import CalendarIcon from "./icons/CalendarIcon.vue";
import { parseDate } from "../utils/utils";

const options: Intl.DateTimeFormatOptions = {
  weekday: "long",
  year: "numeric",
  month: "long",
  day: "numeric",
  timeZone: "UTC",
  hour: "2-digit",
  minute: "2-digit",
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
  <div class="bg-stone-100 dark:bg-stone-800 shadow-lg rounded-lg p-4 mb-4">
    <p class="text-xl font-semibold mb-4 text-center">
      {{ props.subject }}
    </p>
    <p class="text-lg mb-4 break-all">
      {{ props.body }}
    </p>
    <div class="flex flex-nowrap items-center gap-2">
      <MailFromIcon class="w-6 h-6 fill-red-500" />
      <span class="font-medium text-cyan-700 dark:text-cyan-500"> From: </span>
    </div>
    <p class="mb-4">
      {{ props.from }}
    </p>
    <div class="flex flex-nowrap items-center gap-2">
      <MailToIcon class="w-6 h-6" />
      <span class="font-medium text-cyan-700 dark:text-cyan-500">To:</span>
    </div>
    <p class="mb-4">
      {{ props.to ? props.to.join(", ") : "" }}
    </p>
    <div class="flex flex-nowrap items-center gap-2">
      <MailToIcon class="w-6 h-6" />
      <span class="font-medium text-cyan-700 dark:text-cyan-500">CC:</span>
    </div>
    <p class="mb-4">
      {{ props.cc ? props.cc.join(", ") : "" }}
    </p>
    <div class="flex flex-nowrap items-center gap-2">
      <CalendarIcon class="w-6 h-6" />
      <span class="font-medium text-cyan-700 dark:text-cyan-500">Date:</span>
    </div>
    <p class="mb-4">
      {{ parseDate(props.date, options) }}
    </p>
  </div>
</template>
