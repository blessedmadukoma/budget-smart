<template>
  <div
    class="flex flex-col items-center justify-center min-h-screen py-12 px-4 sm:px-6 lg:px-8 bg-gray-50 dark:bg-gray-900"
  >
    <!-- Logo -->
    <div class="mb-6">
      <img src="/logo.svg" alt="BudgetSmart Logo" class="w-16 h-16" />
    </div>

    <!-- Auth Card -->
    <div class="w-full max-w-md">
      <div
        class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 p-8"
      >
        <!-- Header -->
        <h1 class="text-2xl font-bold mb-1">
          {{ mode === "login" ? "Sign In" : "Sign Up" }}
        </h1>
        <p class="text-gray-600 dark:text-gray-400 mb-6">
          {{
            mode === "login"
              ? "Welcome back to BudgetSmart!"
              : "Create your BudgetSmart account"
          }}
        </p>

        <!-- Google Auth -->
        <CommonButton
          variant="secondary"
          fullWidth
          class="flex items-center justify-center bg-gray-100 hover:bg-gray-200 text-gray-800 dark:bg-gray-700 dark:hover:bg-gray-600 dark:text-gray-100 mb-4"
          @click="handleGoogleAuth"
        >
          <img
            src="/svgs/google-icon.svg"
            alt="Google Icon"
            class="w-5 h-5 mr-2"
          />

          {{ mode === "login" ? "Sign in with Google" : "Sign up with Google" }}
        </CommonButton>

        <!-- Divider -->
        <div class="relative my-6">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span
              class="px-2 bg-white dark:bg-gray-800 text-gray-500 dark:text-gray-300"
              >or</span
            >
          </div>
        </div>

        <!-- Form -->
        <form @submit.prevent="handleSubmit">
          <!-- Name fields (only shown for registration) -->
          <div v-if="mode === 'register'" class="mb-4 grid grid-cols-2 gap-4">
            <div>
              <label for="firstName" class="sr-only">First Name</label>
              <input
                id="firstName"
                v-model="form.firstName"
                type="text"
                required
                autocomplete="off"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-black"
                placeholder="First Name"
              />
            </div>
            <div>
              <label for="lastName" class="sr-only">Last Name</label>
              <input
                id="lastName"
                v-model="form.lastName"
                type="text"
                required
                autocomplete="off"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-black"
                placeholder="Last Name"
              />
            </div>
          </div>

          <!-- Email -->
          <div class="mb-4">
            <label for="email" class="sr-only">Email</label>
            <input
              id="email"
              v-model="form.email"
              type="email"
              required
              autocomplete="off"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-black"
              placeholder="Email"
            />
          </div>

          <!-- Password -->
          <div class="mb-4 relative">
            <label for="password" class="sr-only">Password</label>
            <input
              id="password"
              v-model="form.password"
              :type="showPassword ? 'text' : 'password'"
              required
              autocomplete="off"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-black"
              placeholder="Password"
            />
            <button
              type="button"
              @click="showPassword = !showPassword"
              class="absolute inset-y-0 right-0 b-3 pr-3 flex items-center"
            >
              <img
                :src="showPassword ? '/svgs/eye-off.svg' : '/svgs/eye-on.svg'"
                alt="Toggle Password Visibility"
                class="h-5 w-5 text-gray-400"
              />
            </button>

            <!-- Add Password Strength Indicator here -->
            <!-- <PasswordStrength class="pt-2" :password="form.password" /> -->
          </div>

          <!-- Confirm Password (only for registration) -->
          <div v-if="mode === 'register'" class="mb-4 relative">
            <label for="confirmPassword" class="sr-only"
              >Confirm Password</label
            >
            <input
              id="confirmPassword"
              v-model="form.confirmPassword"
              :type="showPassword ? 'text' : 'password'"
              required
              autocomplete="off"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-black"
              placeholder="Confirm Password"
            />
          </div>

          <!-- Cloudflare Message -->
          <!-- <div class="flex items-center mb-4">
            <div class="bg-green-100 rounded-full p-1">
              <svg
                class="w-5 h-5 text-green-600"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                <path
                  fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>
            <span class="ml-2 text-green-700">Success!</span>

            <div class="flex items-center ml-auto">
              <img
                src="https://www.cloudflare.com/img/logo-cloudflare-dark.svg"
                alt="Cloudflare"
                class="h-5"
              />
              <div class="text-xs ml-2">
                <a href="#" class="text-gray-600 hover:underline">Privacy</a> ·
                <a href="#" class="text-gray-600 hover:underline">Terms</a>
              </div>
            </div>
          </div> -->

          <!-- Actions -->
          <div class="flex justify-between items-center mt-6">
            <CommonButton variant="secondary" @click="goToHome">
              Back to Home
            </CommonButton>

            <CommonButton
              type="submit"
              variant="primary"
              size="medium"
              :disabled="loading"
              :class="{
                'cursor-not-allowed bg-gray-500 hover:bg-gray-500': loading,
              }"
            >
              <template v-if="loading">
                <loader class="w-5 h-5 mx-2 inline-block" />
              </template>
              <template v-else>
                {{ mode === "login" ? "Sign In" : "Sign Up" }}
              </template>
            </CommonButton>
          </div>
        </form>
      </div>

      <!-- Switch between login/register -->
      <div class="text-center mt-4">
        <p v-if="mode === 'login'">
          No account?
          <NuxtLink
            to="/register"
            class="font-medium text-blue-400 hover:text-blue-600"
          >
            Sign Up
          </NuxtLink>
        </p>
        <p v-else>
          Already have an account?
          <NuxtLink
            to="/login"
            class="font-medium text-blue-400 hover:text-blue-600"
          >
            Sign In
          </NuxtLink>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
  import { useRouter } from "nuxt/app";
  import { reactive, ref } from "vue";
  import { useAuthService } from "~/composables/useAuthService";
  import { useNotification } from "~/composables/useNotification";
  import { useStore } from "~/store/store";
  import loader from "./common/loader.vue";

  const props = defineProps({
    mode: {
      type: String,
      default: "login",
      validator: (value) => ["login", "register"].includes(value),
    },
  });

  const emit = defineEmits(["success"]);
  const router = useRouter();
  const notify = useNotification();

  const form = reactive({
    firstName: "",
    lastName: "",
    email: "",
    password: "",
    confirmPassword: "",
  });

  const showPassword = ref(false);
  const loading = ref(false);
  const success = ref(false);

  const authService = useAuthService();
  const store = useStore();

  const handleSubmit = async (event) => {
    if (form.password.length < 6) {
      notify.toast.error(
        "Password is too short. Please use at least 6 characters."
      );
      return;
    }

    if (props.mode === "register" && form.password !== form.confirmPassword) {
      notify.toast.error("Passwords do not match");
      return;
    }

    loading.value = true;

    try {
      if (props.mode === "login") {
        handleLogin(form);
        loading.value = false;
      } else {
        handleRegister(form);
        loading.value = false;
      }
    } catch (err) {
      notify.toast.error("Authentication error:", err);
    }
  };

  const handleRegister = async () => {
    try {
      await authService.register(form);

      // if successful, inform user should check email for OTP or something...

      const user = await authService.getUser();

      store.setUser(user);

      goToDashboard();
    } catch (error) {}
  };

  const handleLogin = async (form) => {
    try {
      const response = await authService.login({
        email: form.email,
        password: form.password,
      });

      await store.checkAuth();

      success.value = true;

      emit("success", {
        email: form.email,
        mode: props.mode,
      });

      router.push("/dashboard");
    } catch (error) {
      const errorMessage = error.message || "Authentication failed";

      notify.toast.error(`Authentication failed: ${errorMessage}`);
    }
  };

  const handleGoogleAuth = () => {
    console.log("handling oauth");

    // $auth.oauthLogin("google");
  };

  const goToHome = () => {
    router.push("/");
  };
  const goToDashboard = () => {
    router.push("/dashboard");
  };
</script>
