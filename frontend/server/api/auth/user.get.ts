import { serverApi } from "~/server/api/utils/serverApi";

export default defineEventHandler(async (event) => {
  try {
    const token = getCookie(event, "auth_token");

    if (!token) {
      throw createError({
        statusCode: 401,
        message: "No authentication token found",
      });
    }

    const api = serverApi(event);
    const res = await api.raw("/users/me", "GET");

    return {
      data: res.data.data,
      authenticated: true,
    };
  } catch (error: any) {
    console.error("User fetch error details:", error);

    throw createError({
      statusCode: error.response?.status || 500,
      message: error.message || "Failed to fetch user",
    });
  }
});
