<template>
  <div>
    <div v-if="loading">Loading user data...</div>
    <div v-else-if="user">
      <h1>
        Welcome to your Dashboard, {{ user.firstName }} - {{ user.email }}
      </h1>
    </div>
    <div v-else>
      <p>Not logged in. Redirecting...</p>
    </div>
  </div>
</template>

<script setup>
  import { onMounted, ref, watch } from "vue";
  import { useStore } from "~/store/store";

  definePageMeta({
    middleware: "auth",
  });

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
