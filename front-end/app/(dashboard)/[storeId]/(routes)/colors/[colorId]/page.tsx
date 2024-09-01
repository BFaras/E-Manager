import React from 'react'
import ColorForm from './components/color-form';
import axiosInstance, { setUpInterceptor } from '@/app/utils/axios_instance';
import { auth } from '@clerk/nextjs/server';

export default async function  SizePage( {params} :
    {params:{colorId:string,storeId:string}}) {
    const {getToken} = auth()
    await setUpInterceptor(getToken)
    
    const colorResponse = await axiosInstance.get(`stores/${params.storeId}/colors/${params.colorId}`);
    const color = colorResponse.data

  return (
    <div className="flex-col">
      <div className = 'flex-1 space-y-4 p-8 pt-6'>
        < ColorForm initialData ={color}></ ColorForm>

      </div>
    </div>
  )
}

