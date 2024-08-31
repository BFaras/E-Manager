import React, { useEffect } from "react";
import CategoryForm from "./components/category-form";
import { auth } from "@clerk/nextjs/server";
import axiosInstance, { setUpInterceptor } from "@/app/utils/axios_instance";

export default async function CategoryPage({
  params,
}: {
  params: { categoryId: string; storeId: string };
}) {

  const {getToken} = auth()

  await setUpInterceptor(getToken)

  const responseCategory = await axiosInstance.get(`stores/${params.storeId}/categrories/${params.categoryId}`)
  const category = responseCategory.data
  const reponseBillboards = await axiosInstance.get(`stores/${params.storeId}/billboards`)
  const billboards = reponseBillboards.data

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <CategoryForm
          billboards={billboards}
          initialData={category}
        ></CategoryForm>
      </div>
    </div>
  );
}
