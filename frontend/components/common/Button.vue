<template>
  <button
    :type="type"
    :class="[
      'px-4 py-2 rounded-md text-sm font-medium transition-colors',
      variant === 'primary' &&
        'bg-black text-white hover:bg-gray-800 dark:bg-white dark:text-black dark:hover:bg-gray-200',
      variant === 'secondary' &&
        'border border-gray-300 text-gray-700 hover:bg-gray-50 dark:border-gray-600 dark:text-gray-200 dark:hover:bg-gray-800',
      size === 'small' && 'px-3 py-1 text-xs',
      size === 'large' && 'px-6 py-3 text-base',
      fullWidth && 'w-full',
      className,
    ]"
    @click="$emit('click')"
  >
    <slot></slot>
  </button>
</template>

<script setup>
  defineProps({
    type: {
      type: String,
      default: "button",
      validator: (value) => ["button", "submit", "reset"].includes(value),
    },
    variant: {
      type: String,
      default: "primary",
      validator: (value) => ["primary", "secondary"].includes(value),
    },
    size: {
      type: String,
      default: "medium",
      validator: (value) => ["small", "medium", "large"].includes(value),
    },
    fullWidth: {
      type: Boolean,
      default: false,
    },
    className: {
      type: String,
      default: "",
    },
  });

  defineEmits(["click"]);
</script>
