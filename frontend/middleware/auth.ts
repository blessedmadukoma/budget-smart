import { navigateTo } from "#app";
import { useStore } from "~~/store/store";

export default defineNuxtRouteMiddleware(async (to) => {
  const store = useStore();

  // Skip server-side redirects, let client handle it
  if (process.server) return;

  // Ensure auth state is checked
  if (store.authStatus === "idle" || store.authStatus === "error") {
    await store.checkAuth();
  }

  if (
    !store.isAuthenticated &&
    !(to.path === "/login") &&
    !(to.path === "/register")
  ) {
    return navigateTo("/login");
  }

  if (
    store.isAuthenticated &&
    (to.path === "/login" || to.path === "/register" || to.path === "/")
  ) {
    return navigateTo("/dashboard");
  }
});
