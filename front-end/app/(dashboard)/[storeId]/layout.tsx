import axiosInstance from "@/app/utils/axios_instance";
import Navbar from "@/components/navbar";
import { Store } from "@/models/db";
import { auth } from "@clerk/nextjs/server";
import { redirect } from "next/navigation";

async function fetchStore(userId: string, storeId: string) {
  try {
    const response = await axiosInstance.get(
      `users/${userId}/stores/${storeId}`
    );
    return response.data;
  } catch (err) {}
}

export default async function DashboardLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { storeId: string };
}) {
  const { userId } = auth();

  if (!userId) {
    redirect("/sign-in");
  }

  let store: Store = await fetchStore(userId, params.storeId);

  if (!store) {
    redirect("/");
  }

  return (
    <>
      <Navbar></Navbar>
      {children}
    </>
  );
}
