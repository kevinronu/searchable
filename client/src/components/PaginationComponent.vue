<script setup lang="ts">
import { computed } from "vue";
import CursorLeftIcon from "./icons/CursorLeftIcon.vue";
import CursorRightIcon from "./icons/CursorRightIcon.vue";

const props = defineProps({
  location: {
    type: String,
    required: true,
  },
  currentPage: {
    type: Number,
    required: true,
  },
  totalQuantity: {
    type: Number,
    required: false,
    default: 0,
  },
  quantityPerPage: {
    type: Number,
    required: true,
  },
  maxPagesGenerated: {
    type: Number,
    required: true,
  },
});

const initialPage = computed(
  () => (Math.ceil(props.currentPage / props.maxPagesGenerated) - 1) * 10 + 1
);

const numberOfPages = computed(() =>
  Math.ceil(props.totalQuantity / props.quantityPerPage)
);
const remainingPagesToDisplay = computed(() => {
  const pages = numberOfPages.value - initialPage.value + 1;
  return pages > 0 ? pages : 0;
});

const validation = computed(
  () => remainingPagesToDisplay.value > props.maxPagesGenerated
);

const pagesToDisplay = computed(() =>
  validation.value ? props.maxPagesGenerated : remainingPagesToDisplay.value
);
</script>

<template>
  <nav class="text-sm">
    <ul class="flex flex-wrap gap-2 justify-center items-center">
      <router-link
        v-if="initialPage > 1 && remainingPagesToDisplay > 0"
        :to="`${location}/page/${initialPage - maxPagesGenerated}`"
        class="cursor-pointer"
      >
        <CursorLeftIcon class="w-6 h-6" />
      </router-link>
      <router-link
        v-for="n in pagesToDisplay"
        :key="n"
        :to="`${location}/page/${initialPage + n - 1}`"
        class="cursor-pointer py-1 px-1 rounded"
        active-class="border-2 border-pink-600 dark:border-pink-600 bg-pink-200 dark:bg-pink-900"
      >
        <li>{{ initialPage + n - 1 }}</li>
      </router-link>
      <router-link
        v-if="validation"
        :to="`${location}/page/${initialPage + maxPagesGenerated}`"
        class="cursor-pointer"
      >
        <CursorRightIcon class="w-6 h-6" />
      </router-link>
    </ul>
  </nav>
</template>
