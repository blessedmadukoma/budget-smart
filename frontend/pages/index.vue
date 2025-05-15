<template>
  <div>
    <h3 class="text-3xl text-red-500">BudgetSmart HQ</h3>
    <h1>Dashboard</h1>
    <p>Welcome to your dashboard!</p>

    <div class="space-x-10">
      <NuxtLink to="/register">Register</NuxtLink>
      <NuxtLink to="/login">Login</NuxtLink>
    </div>
  </div>
</template>

<script setup>
  import no_auth from "~/middleware/no_auth";

  definePageMeta({
    middleware: no_auth,
  });
</script>
