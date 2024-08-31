
import React from "react";
import BillboardForm from "./components/billboard-form";
import { auth } from "@clerk/nextjs/server";
import axiosInstance, { setUpInterceptor } from "@/app/utils/axios_instance";

export default async function BillboardPage({
  params,
}: {
  params: { billboardId: string };
}) {
  const { getToken } = auth();
  await setUpInterceptor(getToken);
  
  const response = await axiosInstance.get(`billboards/${params.billboardId}`)
  const billboard = response.data;

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <BillboardForm initialData={billboard}></BillboardForm>
      </div>
    </div>
  );
}
