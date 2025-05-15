<!-- pages/register.vue -->
<template>
  <AuthForm mode="register" @success="handleRegisterSuccess" />
</template>

<script setup>
  import { useRouter } from "nuxt/app";
  import no_auth from "~/middleware/no_auth";

  definePageMeta({
    layout: false,
    middleware: no_auth,
  });

  const router = useRouter();

  const handleRegisterSuccess = (data) => {
    console.log("User registered:", data);

    router.push("/dashboard");
  };
</script>
