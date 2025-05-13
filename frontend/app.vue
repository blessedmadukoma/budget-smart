<template>
  <div>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>

<style>
  html {
    /* Add any transition for smooth color changes */
    transition: background-color 0.3s ease, color 0.3s ease;
  }

  /* Base styling for light mode */
  html {
    background-color: #ffffff;
    color: #000000;
  }

  /* Styling for dark mode */
  html.dark {
    background-color: #111827; /* Tailwind's gray-900 */
    color: #f3f4f6; /* Tailwind's gray-100 */
  }
</style>
