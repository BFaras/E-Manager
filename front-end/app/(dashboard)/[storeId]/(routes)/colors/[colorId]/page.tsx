import prismaDB from '@/lib/prismadb'
import React from 'react'
import ColorForm from './components/color-form';

export default async function  SizePage( {params} :
    {params:{colorId:string}}) {

    const color = await prismaDB.color.findUnique({
        where:{
            id: params.colorId
        }
    });

  return (
    <div className="flex-col">
      <div className = 'flex-1 space-y-4 p-8 pt-6'>
        < ColorForm initialData ={color}></ ColorForm>

      </div>
    </div>
  )
}

