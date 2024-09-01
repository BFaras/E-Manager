import prismaDB from "@/lib/prismadb";
import React from "react";
import SizeForm from "./components/size-form";
import { auth } from "@clerk/nextjs/server";
import axiosInstance, { setUpInterceptor } from "@/app/utils/axios_instance";

export default async function SizePage({
  params,
}: {
  params: { sizeId: string, storeId: string };
}) {

  const {getToken} = auth()
  await setUpInterceptor(getToken)

  const responseSize = await axiosInstance.get(`stores/${params.storeId}/sizes/${params.sizeId}`);
  const size = responseSize.data;

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <SizeForm initialData={size}></SizeForm>
      </div>
    </div>
  );
}
