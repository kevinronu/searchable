import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import Search from "../views/Search.vue";
import Document from "../views/Document.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/search/:query/page/:currentPage",
    name: "Search",
    component: Search,
  },
  {
    path: "/document/:documentId",
    name: "Document",
    component: Document,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
