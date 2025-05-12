<!-- pages/register.vue -->
<template>
  <AuthForm mode="register" @success="handleRegisterSuccess" />
</template>

<script setup>
  import { useRouter } from "nuxt/app";

  definePageMeta({
    layout: false,
  });

  const router = useRouter();

  const handleRegisterSuccess = (data) => {
    console.log("User registered:", data);

    // Maybe navigate to a welcome page or dashboard
    router.push("/welcome");
  };
</script>
