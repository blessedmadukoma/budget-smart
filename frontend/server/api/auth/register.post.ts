import { H3Event } from "h3";
import { serverApi } from "~/server/api/utils/serverApi";

export default defineEventHandler(async (event: H3Event) => {
  const api = serverApi(event);

  const { firstName, lastName, email, password } = await readBody(event);

  try {
    const res = await api.raw("/auth/register", "POST", {
      data: {
        firstName,
        lastName,
        email,
        password,
        authProvider: "local",
      },
    });

    const cookies = res.headers["set-cookie"] || [];

    for (const cookie of cookies) {
      appendHeader(event, "set-cookie", cookie);
    }

    return { data: res.data, cookies };
  } catch (error) {
    console.error("Register failed:", error);
    throw createError({
      statusCode: 500,
      message: "Register failed",
    });
  }
});
