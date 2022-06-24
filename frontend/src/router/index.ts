import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import HomeView from "../views/HomeView.vue";
import OmoshiroGoMojiView from "../views/OmoshiroGoMojiView.vue";
import OmoshiroGoMojiCreateFormView from "../views/OmoshiroGoMojiCreateFormView.vue";
import OmoshiroGoMojiUpdateFormView from "../views/OmoshiroGoMojiUpdateFormView.vue";

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
  {
    path: "/omoshiro_go_moji/edit/:id",
    name: "omoshiro_go_moji_update_form",
    component: OmoshiroGoMojiUpdateFormView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
