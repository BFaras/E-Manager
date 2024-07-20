import React from 'react'
import BillboardClient from './components/size-client'
import prismaDB from '@/lib/prismadb'
import { BillboardColumn } from './components/columns'
import { format} from "date-fns";

export default async function BillboardsPage({params}:{
  params: {
    storeId: string
  }
}) {

  const billboards = await prismaDB.billboard.findMany({
    where: {
      storeId: params.storeId
    },
    orderBy:{
      createdAt: 'desc'
    }
  })

  const formatedBillboards :BillboardColumn[] = billboards.map((item)=> ({
    id:item.id,
    label: item.label,
    createdAt:format(item.createdAt,"MMMM do, yyyy")
  }))
  return (
    <div className='flex-col'>
      <div className ="flex-1 space-y-4 p-8 pt-6">
        <BillboardClient data={formatedBillboards}></BillboardClient>


      </div>
    </div>
  )
}

