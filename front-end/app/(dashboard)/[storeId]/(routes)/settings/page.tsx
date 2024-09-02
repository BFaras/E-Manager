import axiosInstance, { setUpInterceptor } from "@/app/utils/axios_instance";
import SettingForm from "./components/settings-form";
import prismaDB from "@/lib/prismadb";
import { auth } from "@clerk/nextjs/server";
import { redirect } from "next/navigation";
import React from "react";

interface SettingPageProps {
  params: {
    storeId: string;
  };
}
export default async function SettingsPage({ params }: SettingPageProps) {
  const { userId, getToken } = auth();

  if (!userId) {
    redirect("/sign-in");
  }

  await setUpInterceptor(getToken);

  const storeResponse = await axiosInstance.get(`users/${userId}/stores/${params.storeId}`);
  const store = storeResponse.data
  
  if (!store) {
    redirect("/");
  }

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <SettingForm initialData={store}></SettingForm>
      </div>
    </div>
  );
}
