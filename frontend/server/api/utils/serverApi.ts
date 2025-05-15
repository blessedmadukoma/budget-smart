import { H3Event } from "h3";
import { ApiWrapper } from "~/utils/api/api";

export const serverApi = (event: H3Event) => {
  const {
    public: { apiBaseUrl },
  } = useRuntimeConfig();
  // const accessToken = getCookie(event, "Authorization");
  // const refreshToken = getCookie(event, "Refresh-Token");

  const accessToken = getCookie(event, "auth_token");

  const apiClient = new ApiWrapper({
    baseURL: apiBaseUrl as string,
    headers: {
      ...(accessToken && { Authorization: `Bearer ${accessToken}` }),
      // ...(refreshToken && { "Refresh-Token": refreshToken }),
    },
  });

  return apiClient;
};
