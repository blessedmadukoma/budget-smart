<template>
  <div>
    <div v-if="loading">Loading user data...</div>
    <div v-else-if="user">
      <!-- Main content max width -->
      <div class="flex justify-between items-center mb-4">
        <section>
          <h1 class="text-3xl font-bold capitalize">{{ currentPage }}</h1>
        </section>
        <div class="flex space-x-2">
          <LayoutTransactionsHeader />
        </div>
      </div>

      <p>
        Welcome to your Transactions, {{ user.firstName }} - {{ user.email }}
      </p>
    </div>
    <div v-else>
      <p>Not logged in. Redirecting...</p>
    </div>
  </div>
</template>

<script setup>
  import { useRoute } from "#app";
  import { computed, onMounted, ref, watch } from "vue";
  import { useStore } from "~/store/store";

  definePageMeta({
    middleware: "auth",
    layout: "logged-in",
    name: "transactions",
  });

  const route = useRoute();
  const currentPage = computed(() => route.name || "Transactions");

  const userStore = useStore();
  const user = ref(userStore.user);
  const loading = ref(true);

  onMounted(async () => {
    try {
      loading.value = true;
      if (!userStore.user) {
        await userStore.checkAuth();
      }
      user.value = userStore.user;

      user.value = user.value.data;
    } catch (err) {
      console.error("Failed to load user data:", err);
    } finally {
      loading.value = false;
    }
  });

  watch(
    () => userStore.user,
    (newUser) => {
      user.value = newUser;
    }
  );
</script>
