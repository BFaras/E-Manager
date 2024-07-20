import React from 'react'
import prismaDB from '@/lib/prismadb'
import { ColorColumn } from './components/columns'
import { format } from "date-fns";
import ColorClient from './components/color-client';

export default async function ColorsPage({params}:{
  params: {
    storeId: string
  }
}) {

  const colors = await prismaDB.color.findMany({
    where: {
      storeId: params.storeId
    },
    orderBy:{
      createdAt: 'desc'
    }
  })

  const formatedColors :ColorColumn[] = colors.map((item)=> ({
    id:item.id,
    name: item.name,
    value: item.value,
    createdAt:format(item.createdAt,"MMMM do, yyyy")
  }))
  return (
    <div className='flex-col'>
      <div className ="flex-1 space-y-4 p-8 pt-6">
        <ColorClient data={formatedColors}></ColorClient>
      </div>
    </div>
  )
}

