import axios from "axios";

const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_GO_URL,
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
});

export async function setAuthorizationHeader(
  getToken: () => Promise<string | null>
) {
  const token = await getToken();
  if (!token) {
    throw new Error("There is no token ");
  }
  axiosInstance.defaults.headers.common["Authorization"] = `Bearer ${token}`;
}

export default axiosInstance;
