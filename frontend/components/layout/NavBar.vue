<template>
  <nav class="md:px-56 md:py-3">
    <div
      class="md:flex md:justify-between md:items-center md:py-2 md:px-6 rounded-lg border border-gray-400 dark:border-gray-700"
    >
      <NuxtLink to="/" class="flex items-center">
        <img src="/logo.svg" alt="Logo" class="h-7 w-7 mr-1" />
        BudgetSmart
      </NuxtLink>
      <div>
        <ul class="flex space-x-6 text-gray-500 dark:text-gray-400">
          <li
            :class="{
              'hover:text-gray-900 hover:dark:text-gray-100 cursor-pointer underline':
                $route.path === '/',
              'hover:text-gray-900 hover:dark:text-gray-100 cursor-pointer':
                $route.path !== '/',
            }"
          >
            <NuxtLink to="/">Pricing</NuxtLink>
          </li>
          <li
            :class="{
              'hover:text-gray-900 hover:dark:text-gray-100 cursor-pointer underline':
                $route.path === '/register',
              'hover:text-gray-900 hover:dark:text-gray-100 cursor-pointer':
                $route.path !== '/register',
            }"
          >
            <NuxtLink to="/register">Register</NuxtLink>
          </li>
          <li
            :class="{
              'hover:text-gray-900 hover:dark:text-gray-100 cursor-pointer underline':
                $route.path === '/login',
              'hover:text-gray-900 hover:dark:text-gray-100 cursor-pointer':
                $route.path !== '/login',
            }"
          >
            <client-only>
              <button v-if="userStore.isAuthenticated" @click="handleLogout">
                Log out
              </button>
              <NuxtLink v-else to="/login">Log In</NuxtLink>
            </client-only>
          </li>
          <li>
            <button @click="toggleColorMode" class="ml-2">
              <span v-if="isDark">🌞</span>
              <span v-else>🌙</span>
            </button>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script setup>
  import { useRouter } from "nuxt/app";
  import { computed } from "vue";
  import { useStore } from "~/store/store";

  const colorMode = useColorMode();
  const userStore = useStore();
  const router = useRouter();

  const isDark = computed(() => colorMode.value === "dark");

  function toggleColorMode() {
    colorMode.preference = colorMode.value === "dark" ? "light" : "dark";
  }

  const handleLogout = async () => {
    try {
      await userStore.logout();

      router.push("/login");
    } catch (err) {
      console.error("Logout error:", err);
    }
  };
</script>
