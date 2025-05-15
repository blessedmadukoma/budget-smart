import type { LoginRequest, RegisterRequest, User } from "~/types/api";

export function useAuthService() {
  const register = async (credentials: RegisterRequest) => {
    try {
      await $fetch("/api/auth/register", {
        method: "POST",
        body: credentials,
      });
    } catch (error) {
      console.error("Register failed:", error);
      return Promise.reject(error);
    }
  };

  const login = async (credentials: LoginRequest) => {
    try {
      const response = await $fetch("/api/auth/login", {
        method: "POST",
        body: {
          email: credentials.email,
          password: credentials.password,
        },
      });

      return response;
    } catch (error: any) {
      let errorMessage = "Login failed";

      console.log(`Login fauled: ${error}`);

      if (error.response && error.response.data) {
        errorMessage =
          error.response.data.message ||
          error.response.data.error ||
          errorMessage;
      } else if (error.data) {
        errorMessage = error.data.message || error.data.error || errorMessage;
      } else if (error.message) {
        errorMessage = error.message;
      }

      console.error(`Login failed: ${errorMessage}`);

      const enhancedError: any = new Error(errorMessage);
      enhancedError.status = error.response?.status;
      enhancedError.data = error.response?._data;

      return Promise.reject(enhancedError);
    }
  };

  const getUser = async (): Promise<User> => {
    try {
      const user = await $fetch<User>("/api/auth/user", {
        method: "GET",
        // headers: useRequestHeaders(["cookies"]), // sends the httpOnly cookies to the nitro server
        headers: useRequestHeaders(["cookie"]),
      });

      return user;
    } catch (error) {
      console.error("Get user failed:", error);
      return Promise.reject(error);
    }
  };

  const logout = async () => {
    try {
      const user = await $fetch("/api/auth/logout", {
        method: "POST",
      });

      return user;
    } catch (error) {
      console.error("Logout failed:", error);
      return Promise.reject(error);
    }
  };

  return {
    login,
    logout,
    register,
    getUser,
  };
}
