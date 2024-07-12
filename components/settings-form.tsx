"use client"

import React, { useState } from 'react'
import Heading from './heading'
import * as z from "zod";
import { Store } from '@/prisma/prisma-model'
import { Button } from './ui/button'
import { Trash } from 'lucide-react'
import { Separator } from './ui/separator'
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { Form, FormControl, FormField, FormItem, FormLabel } from './ui/form';
import { Input } from './ui/input';
import { useParams, useRouter } from 'next/navigation';
import axios from 'axios';
import toast from 'react-hot-toast';

interface SettingFormProps {
    initialData:Store
}

const formSchema = z.object({
    name: z.string().min(1),
})

type SettingFormValues = z.infer<typeof formSchema>

export default function SettingForm({initialData}:SettingFormProps) {

    const params = useParams()
    const router = useRouter()

    const [open,setOpen] = useState(false);
    const [loading,setLoading] = useState(false);

    const form = useForm<SettingFormValues>({
        resolver: zodResolver(formSchema),
        defaultValues: initialData,
    });
    
    const onSubmit = async (data: SettingFormValues) => {
        try {
            setLoading(true)
            await axios.patch(`/api/stores/${params.storeid}`,data)
            router.refresh()
            setLoading(false)
            toast.success("store updated")
        } catch (error) {
            toast.error("Something went wrong")        }
    }
  return (
    <>
        <div className='flex items-center justify-between'>
                <Heading title="Settings"
                description= "Manage store preference"/>
                <Button variant = "destructive" size= "icon"onClick={()=>setOpen(true)}>
                    <Trash className='h-4 w-4'></Trash>
                </Button>
        </div>
        <Separator></Separator>
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8 w-full border border-solid border-gray-600">
                <div className='grid grid-cols-3 gap-8'>
                    <FormField 
                    control={form.control}
                    name="name"
                    render={({field}) =>(
                        <FormItem >
                            <FormLabel >Name</FormLabel>
                            <FormControl >
                                <Input  disabled={loading} placeholder="Store name" {...field}></Input>
                            </FormControl>
                        </FormItem>
                    )}>
                    </FormField>
                </div>
                <Button disabled = {loading} className='ml-auto' type="submit" >
                    Save Changes
                </Button>
            </form>
        </Form>
   </>
  )
}

