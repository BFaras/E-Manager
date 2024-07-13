import prismaDB from '@/lib/prismadb'
import React from 'react'
import BillboardForm from './components/billboard-form';

export default async function BillboardPage( {params} :
    {params:{billboardId:string}}) {

    const billboard = await prismaDB.billboard.findUnique({
        where:{
            id: params.billboardId
        }
    });

  return (
    <div className="flex-col">
      <div className = 'flex-1 space-y-4 p-8 pt-6'>
        <BillboardForm initialData ={billboard}></BillboardForm>

      </div>
    </div>
  )
}

