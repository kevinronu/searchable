import { createRouter, createWebHistory } from "vue-router";

import HomeView from "../views/HomeView.vue";
import SearchView from "../views/SearchView.vue";
import DocumentView from "../views/DocumentView.vue";

const routes = [
  {
    path: "/",
    name: "HomeView",
    component: HomeView,
  },
  {
    path: "/search/:query/page/:currentPage",
    name: "SearchView",
    component: SearchView,
  },
  {
    path: "/document/:documentId",
    name: "DocumentView",
    component: DocumentView,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
