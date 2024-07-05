"use client"

import React from 'react'
import Heading from './heading'
import * as z from "zod";
import { Store } from '@/prisma/prisma-model'
import { Button } from './ui/button'
import { Trash } from 'lucide-react'
import { Separator } from './ui/separator'
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';

interface SettingFormProps {
    initialData:Store
}

const formSchema = z.object({
    name: z.string().min(1),
})


type SettingFormValues = z.infer<typeof formSchema>

export default function SettingForm({initialData}:SettingFormProps) {

    const form = useForm<SettingFormValues>({
        resolver: zodResolver(formSchema),
        defaultValues: initialData,
    });

    
  return (
    <>
   <div className='flex items-center justify-between'>
        <Heading title="Settings"
        description= "Manage store preference"/>
        <Button variant = "destructive" size= "icon" onClick={()=>{}}>
            <Trash className='h-4 w-4'></Trash>
        </Button>
   </div>
   <Separator></Separator>
   </>
  )
}

