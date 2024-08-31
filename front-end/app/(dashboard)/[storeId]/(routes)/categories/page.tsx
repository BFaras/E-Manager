import React from "react";
import CategoryClient from "./components/category-client";
import { CategoryColumn } from "./components/columns";
import { format } from "date-fns";
import axiosInstance, { setUpInterceptor } from "@/app/utils/axios_instance";
import { auth } from "@clerk/nextjs/server";
import { Category } from "@/models/db";

export default async function CategoriesPage({
  params,
}: {
  params: {
    storeId: string;
  };
}) {
  const {getToken} = auth()
  await setUpInterceptor(getToken)

  const response = await axiosInstance.get(`stores/${params.storeId}/categories`)

  const categories:Category[] = response.data || []

  const formatedCategories: CategoryColumn[] = categories.map((item:any) => ({
    id: item.id,
    name: item.name,
    billboardLabel: item.billboard.label,
    createdAt: format(item.createdAt, "MMMM do, yyyy"),
  }));

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <CategoryClient data={formatedCategories}></CategoryClient>
      </div>
    </div>
  );
}
