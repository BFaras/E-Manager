import React from 'react'
import prismaDB from '@/lib/prismadb'
import { ColorColumn } from './components/columns'
import { format } from "date-fns";
import ColorClient from './components/color-client';
import axiosInstance, { setUpInterceptor } from '@/app/utils/axios_instance';
import { auth } from '@clerk/nextjs/server';

export default async function ColorsPage({params}:{
  params: {
    storeId: string
  }
}) {
  const {getToken} = auth()
  await setUpInterceptor(getToken)

  const responseColors = await axiosInstance.get(`stores/${params.storeId}/colors`)
  const colors = responseColors.data || []

  const formatedColors :ColorColumn[] = colors.map((item:any)=> ({
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

