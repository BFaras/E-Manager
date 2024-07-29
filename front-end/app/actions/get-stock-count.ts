

import prismaDB from "@/lib/prismadb";

export async function getStockCount(storeId:string) {
    const stockCount = await prismaDB.product.count({
        where: {
            storeId,
            isArchived:false
        }
    })


    return stockCount;

}