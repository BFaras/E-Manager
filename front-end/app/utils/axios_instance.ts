import axios from "axios";
import * as jwtDecodeNamespace from "jwt-decode";

const jwtDecode = ((jwtDecodeNamespace as any).default || jwtDecodeNamespace) as (token: string) => any;

const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_GO_URL || "http://back-end:8080",
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
});

let requestInterceptor: number | null = null;
let responseInterceptor: number | null = null;

export async function setUpInterceptor(getToken: any) {

  if (requestInterceptor !== null) {
    axiosInstance.interceptors.request.eject(requestInterceptor);
  }
  if (responseInterceptor !== null) {
    axiosInstance.interceptors.response.eject(responseInterceptor);
  }


  requestInterceptor = axiosInstance.interceptors.request.use(
    async (config) => {
      try {
        const token = await getToken({ refresh: true });
        if (token) {
          config.headers["Authorization"] = `Bearer ${token}`;


          try {
            const decoded: any = jwtDecode(token);
          } catch (decodeError) {
            console.warn(" Could not decode token:", decodeError);
          }
        } else {
          console.warn(" No token available!");
        }
      } catch (error) {
        console.error(" Error fetching token:", error);
      }
      return config;
    },
    (error) => {
      console.error(" Request error intercepted:", error);
      return Promise.reject(error);
    }
  );


  responseInterceptor = axiosInstance.interceptors.response.use(
    (response) => {
      return response;
    },
    async (error) => {
      const originalRequest = error.config;
      if (error.response?.status === 401 && !originalRequest._retry) {
        console.warn(" Token expired (401 Unauthorized). Attempting refresh...");
        originalRequest._retry = true;
        try {
          const newToken = await getToken({ refresh: true });
          if (newToken) {
            originalRequest.headers["Authorization"] = `Bearer ${newToken}`;
            return axiosInstance(originalRequest);
          } else {
            console.error(" Failed to obtain a new token. User needs to log in.");
          }
        } catch (refreshError) {
          console.error(" Token refresh error:", refreshError);
        }
        console.error(" Redirecting user to login due to token failure.");
        return Promise.reject("Unauthorized, please sign in again.");
      }
      console.error(" Response error intercepted:", error);
      return Promise.reject(error);
    }
  );
}

export default axiosInstance;
