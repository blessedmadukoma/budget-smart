<template>
  <AuthForm mode="login" @success="handleLoginSuccess" />
</template>

<script setup>
  import { useRouter } from "nuxt/app";

  definePageMeta({
    layout: false,
  });

  const router = useRouter();

  const handleLoginSuccess = (data) => {
    console.log("User logged in:", data);

    // Navigate to dashboard or home page after login
    router.push("/");
  };
</script>
