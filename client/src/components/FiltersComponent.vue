<script setup lang="ts">
import { useFiltersStore } from "../stores/FiltersStore";

const filters = useFiltersStore();

function handleFromChange(event: Event) {
  const fromDate = (event.target as HTMLInputElement).value;
  filters.setFrom(fromDate);
}

function handleToChange(event: Event) {
  const toDate = (event.target as HTMLInputElement).value;
  filters.setTo(toDate);
}

function handleSortChange(event: Event) {
  const sortValue = (event.target as HTMLSelectElement).value;
  filters.setSort(sortValue);
}
</script>

<template>
  <div class="container m-auto px-4 flex items-center justify-center">
    <form class="w-full md:w-fit md:flex md:items-center md:gap-4">
      <div
        class="flex flex-wrap items-center justify-between gap-4 mb-2 md:mb-0"
      >
        <div>
          <label class="block font-bold mb-1" for="from-date"> From: </label>
          <input
            class="appearance-none rounded w-full p-2 bg-stone-50 dark:bg-stone-800 dark:[color-scheme:dark] border-2 border-pink-600 hover:outline-none hover:border-blue-300 focus:border-blue-400"
            id="from-date"
            type="date"
            :value="filters.getFrom"
            @input="handleFromChange"
          />
        </div>
        <div>
          <label class="block font-bold mb-1" for="to-date"> To: </label>
          <input
            class="appearance-none rounded w-full p-2 bg-stone-50 dark:bg-stone-800 dark:[color-scheme:dark] border-2 border-pink-600 hover:outline-none hover:border-blue-300 focus:border-blue-400"
            id="to-date"
            type="date"
            :value="filters.getTo"
            @input="handleToChange"
          />
        </div>
      </div>
      <div class="flex items-center justify-between gap-8 mb-2 md:mb-0">
        <div>
          <label class="block font-bold mb-1" for="order-by-date">
            Order by:
          </label>
          <select
            class="block appearance-none w-full min-w-[10rem] px-4 py-2 rounded border-2 border-pink-600 bg-stone-50 dark:bg-stone-800 hover:outline-none hover:border-blue-300 focus:border-blue-400"
            id="order-by-date"
            :value="filters.getSort"
            @change="handleSortChange"
          >
            <option value="-date">Date descending</option>
            <option value="+date">Date ascending</option>
          </select>
        </div>
        <button
          class="place-self-end min-w-[9.5rem] text-white font-bold py-[0.65rem] p-4 rounded bg-teal-600 hover:bg-teal-500 focus:outline-none focus:ring focus:ring-blue-400 focus:rounded"
          type="reset"
          @click="filters.clearFilters"
        >
          Clear
        </button>
      </div>
    </form>
  </div>
</template>
