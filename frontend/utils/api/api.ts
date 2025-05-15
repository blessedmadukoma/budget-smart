import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import axios, { AxiosError } from "axios";

export interface ApiResponse<T> {
  data: T;
  message?: string;
  status: number;
}

export class ApiWrapper {
  protected instance: AxiosInstance;

  constructor(configOverrides: AxiosRequestConfig = {}) {
    const runtimeConfig = useRuntimeConfig();

    this.instance = axios.create({
      baseURL: runtimeConfig.public.apiBaseUrl as string,
      timeout: 10000,
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      withCredentials: true,
      ...configOverrides,
    });

    this.setupInterceptors();
  }

  protected setupInterceptors(): void {
    this.instance.interceptors.response.use(
      (response: AxiosResponse) => response,
      (error: AxiosError) => {
        if (error.response?.status === 401 && process.client) {
          window.location.href = "/login";
        }
        return Promise.reject(error);
      }
    );
  }

  public getInstance(): AxiosInstance {
    return this.instance;
  }

  public async raw(
    url: string,
    method:
      | "GET"
      | "HEAD"
      | "PATCH"
      | "POST"
      | "PUT"
      | "DELETE"
      | "CONNECT"
      | "OPTIONS"
      | "TRACE",
    options: AxiosRequestConfig = {}
  ): Promise<AxiosResponse> {
    try {
      return await this.instance({
        ...options,
        url,
        method,
      });
    } catch (err) {
      return Promise.reject(err);
    }
  }
}
