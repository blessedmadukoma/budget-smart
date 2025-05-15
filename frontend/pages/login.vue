<template>
  <AuthForm mode="login" @success="handleLoginSuccess" />
</template>

<script setup>
  import { useRouter } from "nuxt/app";
  import no_auth from "~/middleware/no_auth";

  definePageMeta({
    layout: false,
    middleware: no_auth,
  });

  const router = useRouter();
  const notify = useNotification();

  const handleLoginSuccess = (data) => {
    notify.toast.success("user login successful!");

    router.push("/dashboard");
  };
</script>
