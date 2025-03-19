import axios from "axios";

const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_GO_URL || "http://back-end:8080",
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
});

export async function setUpInterceptor(getToken: () => Promise<string | null>) {
  axiosInstance.interceptors.request.use(
    async (config) => {
      const token = await getToken();
      if (token) {
        config.headers["Authorization"] = `Bearer ${token}`;
      } else {
        throw new Error("No token available");
      }
      return config;
    },
    (error) => Promise.reject(error)
  );

  axiosInstance.interceptors.response.use(
    (response) => response,
    async (error) => {
      const originalRequest = error.config;

      // Handle token expiration
      if (error.response?.status === 401 && !originalRequest._retry) {
        console.warn("Token expired. Attempting refresh...");

        originalRequest._retry = true;
        const newToken = await getToken(); // Refresh token

        if (newToken) {
          originalRequest.headers["Authorization"] = `Bearer ${newToken}`;
          return axiosInstance(originalRequest);
        } else {
          console.error("Failed to refresh token. Redirecting to login.");
          return Promise.reject("Unauthorized, please sign in again.");
        }
      }

      return Promise.reject(error);
    }
  );
}

export default axiosInstance;
