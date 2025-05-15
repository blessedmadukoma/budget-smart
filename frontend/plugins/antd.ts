import { defineNuxtPlugin } from "#app";
import { message, notification } from "ant-design-vue";

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.provide("notification", notification);
  nuxtApp.provide("message", message);
});
