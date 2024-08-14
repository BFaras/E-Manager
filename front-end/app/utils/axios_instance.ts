import axios from 'axios';

const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_GO_URL,
  timeout: 10000, // Optional timeout setting
  headers: {
    'Content-Type': 'application/json',
  },
});

export default axiosInstance;