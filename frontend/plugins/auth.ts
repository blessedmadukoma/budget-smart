import { defineNuxtPlugin } from "#app";
import { useStore } from "~/store/store";

export default defineNuxtPlugin(async (nuxtApp) => {
  const store = useStore();
  await store.checkAuth();
});
