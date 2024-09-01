
import React from "react";
import ProductForm from "./components/product-form";
import { auth } from "@clerk/nextjs/server";
import axiosInstance, { setUpInterceptor } from "@/app/utils/axios_instance";

export default async function ProductPage({
  params,
}: {
  params: { productId: string; storeId: string };
}) {
  const {getToken} = auth()
  await setUpInterceptor(getToken)

  const responseProducts = await axiosInstance.get(`stores/${params.storeId}/products/${params.productId}/image`)
  const product = responseProducts.data;
  console.log("help")
  console.log(product)

  const responseCategories = await axiosInstance.get(`stores/${params.storeId}/categories`)
  const categories = responseCategories.data ;


  const responseSizes = await axiosInstance.get(`stores/${params.storeId}/sizes`)
  const sizes = responseSizes.data ;

  const responseColors = await axiosInstance.get(`stores/${params.storeId}/colors`)
  const colors = responseColors.data;

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <ProductForm
          initialData={product}
          categories={categories}
          colors={colors}
          sizes={sizes}
        ></ProductForm>
      </div>
    </div>
  );
}
