import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import SearchView from "../views/SearchView.vue";
import EmailView from "../views/EmailView.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: HomeView,
  },
  {
    path: "/search/:query",
    name: "Search",
    component: SearchView,
  },
  {
    path: "/document/:documentId",
    name: "Document",
    component: EmailView,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
