import SettingForm from '@/components/settings-form';
import prismaDB from '@/lib/prismadb';
import { auth } from '@clerk/nextjs/server';
import { redirect } from 'next/navigation';
import React from 'react'


interface SettingPageProps {
    params:{
        storeId: string
    }
}
export default async function SettingsPage({ params }: SettingPageProps) {

    const {userId} = auth();

    if(!userId) {
        redirect("/sign-in")
    }

    const store = await prismaDB.store.findFirst({
        where: {
            id: params.storeId,
            userId
        }
    })

    if(!store) {
        redirect("/")
    }
    
  return (
    <div className="flex-col">
        <div className="flex-1 space-y-4 p-8 pt-6">
            <SettingForm initialData={store}></SettingForm>
        </div>
      
    </div>
  )
}

