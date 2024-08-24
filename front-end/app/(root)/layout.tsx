import { auth } from "@clerk/nextjs/server";
import { redirect } from "next/navigation";
import axiosInstance, { setAuthorizationHeader } from "../utils/axios_instance";
import { Store } from "@/models/db";

async function fetchStore(userId: string) {
  try {
    const response = await axiosInstance.get(`/users/${userId}/store`);
    return response.data;
  } catch (err: any) {
    throw new Error("Could not fetch any store");
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

  await setAuthorizationHeader(getToken);
  let store: Store = await fetchStore(userId);
  if (store) {
    redirect(`/${store.id}`);
  }

  return <>{children}</>;
}
