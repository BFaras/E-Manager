import React from "react";
import BillboardClient from "./components/billboard-client";
import { BillboardColumn } from "./components/columns";
import { format } from "date-fns";
import axiosInstance from "@/app/utils/axios_instance";
import { Billboard } from "@/models/db";

export default async function BillboardsPage({
  params,
}: {
  params: {
    storeId: string;
  };
}) {
  
  const response = await axiosInstance.get(
    `stores/${params.storeId}/billboards`
  );
  const billboards: Billboard[] = response.data || [];

  const formatedBillboards: BillboardColumn[] = billboards.map((item) => ({
    id: item.id,
    label: item.label,
    isActive: item.isActive,
    createdAt: format(item.createdAt, "MMMM do, yyyy"),
  }));
  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <BillboardClient data={formatedBillboards}></BillboardClient>
      </div>
    </div>
  );
}
