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

// currentPage = 2
// totalQuantity = 5
// quantityPerPage = 20
// numberOfPages = 0
// maxPagesGenerated = 10
const numberOfPages = computed(() =>
  Math.ceil(props.totalQuantity / props.quantityPerPage)
);
// numberOfPages = 1
const remainingPagesToDisplay = computed(() => {
  const pages = numberOfPages.value - props.currentPage + 1;
  return pages > 0 ? pages : 0;
});
// remainingPagesToDisplay = 1 - 2 + 1 = 0
const validation = computed(
  () => remainingPagesToDisplay.value > props.maxPagesGenerated
);
// validation = 0 > 10 = false
const pagesToDisplay = computed(() =>
  validation.value ? props.maxPagesGenerated : remainingPagesToDisplay.value
);
// pagesToDisplay = 0
</script>

<template>
  <nav class="my-4">
    <ul class="flex flex-wrap gap-3 justify-center">
      <router-link
        v-if="currentPage > 1 && remainingPagesToDisplay > 0"
        :to="`${location}/page/${currentPage - 1}`"
        class="cursor-pointer"
      >
        <CursorLeftIcon class="w-6 h-6" />
      </router-link>
      <router-link
        v-for="n in pagesToDisplay"
        :key="n"
        :to="`${location}/page/${currentPage + n - 1}`"
        class="cursor-pointer py-1 px-2 border-2"
        active-class="border-pink-600 bg-pink-100 dark:bg-pink-900"
      >
        <li>{{ currentPage + n - 1 }}</li>
      </router-link>
      <router-link
        v-if="validation"
        :to="`${location}/page/${currentPage + maxPagesGenerated}`"
        class="cursor-pointer"
      >
        <CursorRightIcon class="w-6 h-6" />
      </router-link>
    </ul>
  </nav>
</template>
