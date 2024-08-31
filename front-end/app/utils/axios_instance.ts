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

  axiosInstance.interceptors.response.use(
    (response) => {
      return response;
    },
    async (error) => {
      const originalRequest = error.config;

      if (error.response.status === 401 && !originalRequest._retry) {
        originalRequest._retry = true; 

        const token = await getToken();
        if (token) {
          axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
          originalRequest.headers["Authorization"] = `Bearer ${token}`;
          return axiosInstance(originalRequest); 
        }
      }
      return Promise.reject(error); 
    }
  );
}
export default axiosInstance;