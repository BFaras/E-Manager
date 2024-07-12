import prismaDB from "@/lib/prismadb";
import { auth } from "@clerk/nextjs/server";
import { NextResponse } from "next/server";

export async function PATCH (
    req: Request,
    {params}: {params: {storeId:string}}
) {
    try {
        const {userId} = auth();
        const body = await req.json()
        const {name}= body

        if(!userId){
            return new NextResponse('Unauthenticated', {status:401})
        }

        if(!name){
            return new NextResponse('Name is required', {status:400})
        }

        if(!params.storeId){
            return new NextResponse('StoreId is required', {status:400})
        }

        const store = await prismaDB.store.updateMany({
            where: {
                id:params.storeId,
                userId
,            },
            data: {
                name
            }   
            }
        )
        return NextResponse.json(store)


    } catch (error) {
        console.log('[STORE_PATCH', error);
        return new NextResponse('Internal error', {status:500})
    }
    
}

export async function DELETE (
    req: Request,
    {params}: {params: {storeId:string}}
) {
    try {
        const {userId} = auth();

        if(!userId){
            return new NextResponse('Unauthenticated', {status:401})
        }

        if(!params.storeId){
            return new NextResponse('StoreId is required', {status:400})
        }

        console.log(params.storeId)

        const store = await prismaDB.store.deleteMany({
            where: {
                id:params.storeId,
                userId
,            } 
            }
        )

        return NextResponse.json(store)


    } catch (error) {
        console.log('[STORE_PATCH', error);
        return new NextResponse('Internal error', {status:500})
    }
}