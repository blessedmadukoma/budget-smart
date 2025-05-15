import { defineStore } from "pinia";
import type { User } from "~/types/api";

export const useStore = defineStore("main", {
  state: () => {
    return {
      user: null as User | null,
      isCheckingAuth: false,
      authStatus: "idle" as "idle" | "loading" | "authenticated" | "error",
    };
  },
  getters: {
    isAuthenticated: (state) => !!state.user,
  },
  actions: {
    setUser(newUser: User | null) {
      this.user = newUser;
      this.authStatus = newUser ? "authenticated" : "error";
      if (process.client) {
        if (newUser) {
          localStorage.setItem("auth_user", JSON.stringify(newUser));
        } else {
          localStorage.removeItem("auth_user");
        }
      }
    },

    async checkAuth() {
      if (this.authStatus === "loading" || this.isCheckingAuth) return;

      this.authStatus = "loading";
      this.isCheckingAuth = true;

      try {
        if (process.client) {
          const storedUser = JSON.parse(
            localStorage.getItem("auth_user") || "null"
          );
          if (storedUser) {
            this.setUser(storedUser);
            return;
          }
        }

        const { getUser } = useAuthService();
        const user = await getUser();
        this.setUser(user);
      } catch (error) {
        console.error("Auth check failed:", error);
        this.setUser(null);
      } finally {
        this.isCheckingAuth = false;
      }
    },

    async logout() {
      try {
        const { logout } = useAuthService();
        await logout();
      } catch (error) {
        console.error("Logout failed:", error);
      } finally {
        this.setUser(null);
      }
    },
  },
});
