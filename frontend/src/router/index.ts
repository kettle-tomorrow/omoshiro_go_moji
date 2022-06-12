import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import HomeView from "../views/HomeView.vue";
import OmoshiroGoMojiView from "../views/OmoshiroGoMojiView.vue";
import OmoshiroGoMojiCreateFormView from "../views/OmoshiroGoMojiCreateFormView.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/home",
    name: "home",
    component: HomeView,
  },
  {
    path: "/",
    name: "omoshiro_go_moji",
    component: OmoshiroGoMojiView,
  },
  {
    path: "/omoshiro_go_moji/new",
    name: "omoshiro_go_moji_create_form",
    component: OmoshiroGoMojiCreateFormView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
