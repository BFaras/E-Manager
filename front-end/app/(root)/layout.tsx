import { auth } from "@clerk/nextjs/server";
import { redirect } from "next/navigation";
import axiosInstance, {  setUpInterceptor } from "../utils/axios_instance";
import { Store } from "@/models/db";

async function fetchStore(userId: string) {
  try {
    const response = await axiosInstance.get(`/users/${userId}/store`);
    return response.data;
  } catch (err: any) {
  }
}

export default async function SetupLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const { userId, getToken } = auth();

  if (!userId) {
    redirect("/sign-in");
  }

  await setUpInterceptor(getToken);
  let store: Store = await fetchStore(userId);
  if (store) {
    redirect(`/${store.id}`);
  }

  return <>{children}</>;
}
