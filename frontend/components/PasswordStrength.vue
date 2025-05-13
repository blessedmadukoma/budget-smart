<template>
  <div class="mt-1">
    <div class="w-full h-1 bg-gray-200 rounded-full overflow-hidden">
      <div
        class="h-full transition-all duration-300"
        :class="strengthColorClass"
        :style="{ width: `${strengthPercentage}%` }"
      ></div>
    </div>

    <p
      v-if="password.length > 0"
      class="text-xs mt-1"
      :class="strengthTextColorClass"
    >
      {{ strengthMessage }}
    </p>
  </div>
</template>

<script setup>
  import { computed } from "vue";

  const props = defineProps({
    password: {
      type: String,
      required: true,
    },
  });

  const passwordStrength = computed(() => {
    if (!props.password) return 0;

    let strength = 0;

    if (props.password.length >= 8) strength += 25;
    else if (props.password.length >= 6) strength += 10;
    else if (props.password.length > 0) strength += 5;

    if (/\d/.test(props.password)) strength += 25;

    if (/[!@#$%^&*(),.?":{}|<>]/.test(props.password)) strength += 25;

    if (/[a-z]/.test(props.password) && /[A-Z]/.test(props.password))
      strength += 25;

    return Math.min(100, strength);
  });

  const strengthPercentage = computed(() => {
    return passwordStrength.value;
  });

  const strengthColorClass = computed(() => {
    const strength = passwordStrength.value;

    if (strength < 20) return "bg-red-500";
    if (strength < 40) return "bg-orange-500";
    if (strength < 60) return "bg-yellow-500";
    if (strength < 80) return "bg-blue-500";
    return "bg-green-500";
  });

  const strengthTextColorClass = computed(() => {
    const strength = passwordStrength.value;

    if (strength < 20) return "text-red-500";
    if (strength < 40) return "text-orange-500";
    if (strength < 60) return "text-yellow-500";
    if (strength < 80) return "text-blue-500";
    return "text-green-500";
  });

  const strengthMessage = computed(() => {
    const strength = passwordStrength.value;

    if (props.password.length === 0) return "";
    if (props.password.length < 6) return "Password too short";
    if (strength < 20) return "Very weak password";
    if (strength < 40) return "Weak password";
    if (strength < 60) return "Medium strength password";
    if (strength < 80) return "Strong password";
    return "Very strong password";
  });
</script>
