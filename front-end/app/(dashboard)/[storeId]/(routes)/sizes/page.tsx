import React from "react";
import SizeClient from "./components/size-client";
import prismaDB from "@/lib/prismadb";
import { SizeColumn } from "./components/columns";
import { format } from "date-fns";
import { auth } from "@clerk/nextjs/server";
import axiosInstance, { setUpInterceptor } from "@/app/utils/axios_instance";

export default async function SizesPage({
  params,
}: {
  params: {
    storeId: string;
  };
}) {
  const {getToken} = auth()
  await setUpInterceptor(getToken)

  const responseSize = await axiosInstance.get(`stores/${params.storeId}/sizes`)
  const sizes = responseSize.data || []

  const formatedSizes: SizeColumn[] = sizes.map((item:any) => ({
    id: item.id,
    name: item.name,
    value: item.value,
    createdAt: format(item.createdAt, "MMMM do, yyyy"),
  }));
  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <SizeClient data={formatedSizes}></SizeClient>
      </div>
    </div>
  );
}
