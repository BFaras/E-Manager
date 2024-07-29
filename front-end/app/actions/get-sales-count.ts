

import prismaDB from "@/lib/prismadb";

export async function getSalesCount(storeId:string) {
    const salesCount = await prismaDB.order.count({
        where: {
            storeId,
            isPaid:true
        }
    })


    return salesCount;

}