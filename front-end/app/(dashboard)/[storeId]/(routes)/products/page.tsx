import React from "react";
import { ProductColumn } from "./components/columns";
import { format } from "date-fns";
import { formatter } from "@/lib/utils";
import ProductClient from "./components/product-client";
import { auth } from "@clerk/nextjs/server";
import axiosInstance, { setUpInterceptor } from "@/app/utils/axios_instance";

export default async function ProductsPage({
  params,
}: {
  params: {
    storeId: string;
  };
}) {
  const { getToken } = auth();
  await setUpInterceptor(getToken);
  
  const response = await axiosInstance.get(
    `stores/${params.storeId}/products`
  );
  const products: any[] = response.data || [];

  const formatedProducts: ProductColumn[] = products.map((item) => ({
    id: item.id,
    name: item.name,
    isFeatured: item.isFeatured,
    isArchived: item.isArchived,
    price: formatter.format(item.price),
    size: item.size.name,
    category: item.category.name,
    color: item.color.value,
    createdAt: format(item.createdAt, "MMMM do, yyyy"),
  }));
  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <ProductClient data={formatedProducts}></ProductClient>
      </div>
    </div>
  );
}
