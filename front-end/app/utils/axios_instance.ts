import axios from "axios";

const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_GO_URL,
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
    (error) => {
      return Promise.reject(error);
    }
  );
}

export default axiosInstance;