import { serverApi } from "~/server/api/utils/serverApi";

export default defineEventHandler(async (event) => {
  try {
    const { email, password } = await readBody(event);

    const api = serverApi(event);

    const res = await api.raw("/auth/login", "POST", {
      data: {
        email,
        password,
        authProvider: "local",
      },
    });

    const token = res.data.data.token;

    if (!token) {
      throw new Error("No token received from authentication service");
    }

    setCookie(event, "auth_token", token, {
      httpOnly: true,
      path: "/",
      maxAge: 60 * 60 * 24 * 7, // 7 days
      secure: process.env.NODE_ENV === "production",
      sameSite: "lax",
    });

    return {
      status_code: 200,
      message: res.data.message || "Login successful",
      data: {
        token: token,
      },
    };
  } catch (error: any) {
    console.error("Login error.....:", error);
    throw createError({
      statusCode: error.response?.status || 500,
      message: error.message || "Login failed",
    });
  }
});
