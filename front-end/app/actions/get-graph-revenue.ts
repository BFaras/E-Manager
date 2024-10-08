

import prismaDB from "@/lib/prismadb";

interface GraphDataDTO {
    name:string;
    total:number;

}
export async function getGraphRevenue(storeId:string) {
    const paidOrders = await prismaDB.order.findMany({
        where: {
            storeId,
            isPaid:true
        },
        include: {
            orderItems:{
                include:{
                    product:true
                }
            }
        }
    })

    const monthlyRevenue: { [key:number] : number } = {}



    for (const order of paidOrders) {
        const month = order.createdAt.getMonth();
        let revenueForOrder = 0;
        for (const item of order.orderItems) {
            revenueForOrder += item.product.price.toNumber();
        }


        monthlyRevenue[month] = (monthlyRevenue[month] || 0) + revenueForOrder;
    };

    const graphData:GraphDataDTO[] = [
        {name:"Jan", total:0},
        {name:"Feb", total:0},
        {name:"Mar", total:0},
        {name:"Apr", total:0},
        {name:"May", total:0},
        {name:"June", total:0},
        {name:"July", total:0},
        {name:"Aug", total:0},
        {name:"Sept", total:0},
        {name:"Oct", total:0},
        {name:"Nov", total:0},
        {name:"Dec", total:0},
    ] 

    for (const month in monthlyRevenue) {
        graphData[parseInt(month)].total = monthlyRevenue[parseInt(month)]
    }

    return graphData;

}