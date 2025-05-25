<template>
  <div class="flex space-x-4 bg-gray-200 min-h-screen">
    <div class="w-64 bg-white max-w-xs flex flex-col justify-between">
      <!-- Sidebar max width -->
      <nav>
        <div class="flex items-center justify-between my-3 border-b">
          <div class="flex items-center justify-start p-4">
            <img src="/logo.svg" alt="Logo" class="h-8 w-8 mr-1" />
            <h1 class="text-gray-800 text-xl">BudgetSmart</h1>
          </div>

          <div class="mr-3 cursor-pointer">
            <img src="/svgs/sidebar.svg" alt="" />
          </div>
        </div>
        <ul class="px-4 py-2 space-y-3">
          <li>
            <NuxtLink
              to="/"
              class="flex block py-3 px-2 rounded hover:bg-gray-100 hover:text-gray-600"
              :class="{
                'font-semibold bg-blue-50 text-blue-800':
                  route?.name === 'dashboard',
                'text-gray-600': route?.name !== 'dashboard',
              }"
            >
              <img
                src="/svgs/dashboard.svg"
                alt=""
                class="mr-2 hover:text-blue-800"
              />
              <p class="block">Dashboard</p>
            </NuxtLink>
          </li>

          <li>
            <NuxtLink
              to="/transactions"
              class="flex block py-3 px-2 rounded hover:bg-gray-100 hover:text-gray-600"
              :class="{
                'font-semibold bg-blue-50 text-blue-800':
                  route?.name === 'transactions',
                'text-gray-600': route?.name !== 'transactions',
              }"
            >
              <img
                src="/svgs/transaction.svg"
                alt=""
                class="mr-2 hover:text-blue-800"
              />
              <p class="block">Transactions</p>
            </NuxtLink>
          </li>

          <li>
            <NuxtLink
              to="#"
              class="flex block py-3 px-2 rounded hover:bg-gray-100 hover:text-gray-600"
              :class="{
                'font-semibold bg-blue-50 text-blue-800':
                  route?.name === 'budgets',
                'text-gray-600': route?.name !== 'budgets',
              }"
            >
              <img
                src="/svgs/pie-chart.svg"
                alt=""
                class="mr-2 hover:text-blue-800"
              />
              <p class="block">Budgets</p>
            </NuxtLink>
          </li>

          <li>
            <NuxtLink
              to="#"
              class="flex block py-3 px-2 rounded hover:bg-gray-100 hover:text-gray-600"
              :class="{
                'font-semibold bg-blue-50 text-blue-800':
                  route?.name === 'analytics',
                'text-gray-600': route?.name !== 'analytics',
              }"
            >
              <img
                src="/svgs/analytics.svg"
                alt=""
                class="mr-2 hover:text-blue-800"
              />
              <p class="block">Analytics</p>
            </NuxtLink>
          </li>

          <li>
            <NuxtLink
              to="#"
              class="flex block py-3 px-2 rounded hover:bg-gray-100 hover:text-gray-600"
              :class="{
                'font-semibold bg-blue-50 text-blue-800':
                  route?.name === 'predictions',
                'text-gray-600': route?.name !== 'predictions',
              }"
            >
              <img
                src="/svgs/prediction.svg"
                alt=""
                class="mr-2 hover:text-blue-800"
              />
              <p class="block">Predictions</p>
            </NuxtLink>
          </li>
        </ul>
      </nav>

      <div class="mt-auto p-4">
        <button class="bg-blue-100 py-2 px-6 rounded-lg w-full">AI Chat</button>
      </div>
    </div>
    <div class="flex-1 py-4 container mx-auto w-full px-4">
      <slot />
    </div>
  </div>
</template>

<script setup>
  import { useRoute } from "#app";

  const route = useRoute();

  console.log(route.name);
</script>
