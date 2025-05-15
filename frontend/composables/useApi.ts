import { ApiWrapper } from "~/utils/api/api";

export const useApi = () => {
  const { apiBaseUrl } = useRuntimeConfig().public;

  const accessToken = useCookie("Authorization");
  // const refreshToken = useCookie("Refresh-Token");

  const apiClient = new ApiWrapper({
    baseURL: apiBaseUrl as string,
    headers: {
      ...(accessToken.value && {
        Authorization: `Bearer ${accessToken.value}`,
      }),
      // ...(refreshToken.value && { "Refresh-Token": refreshToken.value }),
    },
  });

  return apiClient;
};
