import React from "react";
import { format } from "date-fns";
import { formatter } from "@/lib/utils";
import OrderClient from "./components/order-client";
import { OrderColumn } from "./components/columns";
import { auth } from "@clerk/nextjs/server";
import axiosInstance, { setUpInterceptor } from "@/app/utils/axios_instance";

export default async function OrdersPage({
  params,
}: {
  params: {
    storeId: string;
  };
}) {
  const {getToken} = auth()
  await setUpInterceptor(getToken)

  const response = await axiosInstance.get(`stores/${params.storeId}/orders`)
  const orders:any[] = response.data || []

  /*the filter here is because OrderItem_producId_fk deleted cascade, but order might still have orderItems 
  TODO: when deleting product I need to delete orderId */
  const formatedOrders: OrderColumn[] = orders.filter((item) => item.orderItems !== null).map((item) => ({
    id: item.id,
    phone: item.phone,
    address: item.address,
    products: item.orderItems
      .map((orderItem:any) => orderItem.product.name)
      .join(", "),
    totalPrice: formatter.format(
      item.orderItems.reduce((total:string, item:any) => {
        return total + Number(item.product.price);
      }, 0)
    ),
    isPaid: item.isPaid,
    createdAt: format(item.createdAt, "MMMM do, yyyy"),
  }));
  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <OrderClient data={formatedOrders}></OrderClient>
      </div>
    </div>
  );
}
