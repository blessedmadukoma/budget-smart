<template>
  <nav class="md:px-56 md:py-4">
    <div
      class="md:flex md:justify-between md:items-center md:py-2 md:px-6 rounded-lg border border-gray-200 text-md"
    >
      <NuxtLink to="/" class="flex items-center">
        <img src="/logo.svg" alt="Logo" class="h-7 w-7 mr-1" />
        BudgetSmart
      </NuxtLink>
      <div>
        <ul class="flex space-x-6 text-gray-500">
          <li
            :class="{
              'text-gray-900 cursor-pointer': $route.path === '/pricing',
              'hover:text-gray-900 cursor-pointer': $route.path !== '/pricing',
            }"
          >
            <NuxtLink to="/pricing">Pricing</NuxtLink>
          </li>
          <li
            :class="{
              'text-gray-900 cursor-pointer': $route.path === '/register',
              'hover:text-gray-900 cursor-pointer': $route.path !== '/register',
            }"
          >
            <NuxtLink to="/register">Register</NuxtLink>
          </li>
          <li
            :class="{
              'text-gray-900 cursor-pointer': $route.path === '/login',
              'hover:text-gray-900 cursor-pointer': $route.path !== '/login',
            }"
          >
            <NuxtLink to="/login">Log In</NuxtLink>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
  definePageMeta({
    name: "NavBar",
  });
</script>
