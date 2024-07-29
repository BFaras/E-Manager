"use client"
import React, { useState } from 'react'
import Heading from '@/components/heading'
import * as z from "zod";
import { Button } from '@/components/ui/button'
import { Trash } from 'lucide-react'
import { Separator } from '@/components/ui/separator'
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { Form, FormControl, FormField, FormItem, FormLabel } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { useParams, useRouter } from 'next/navigation';
import axios from 'axios';
import toast from 'react-hot-toast';
import { AlertModal } from '@/app/modals/alert-modal';
import { Color, Size } from '@prisma/client';
import ImageUpload from '@/components/ui/image-upload';


const formSchema = z.object({
    name: z.string().min(1),
    value: z.string().min(1).regex(/^#/,{
        message: "String must be a valid hex code"
    })
})

type ColorFormValues = z.infer<typeof formSchema>

interface ColorFormProps {
    initialData: Color | null
}

export default function ColorForm({initialData}:ColorFormProps) {

    const params = useParams()
    const router = useRouter()

    const [open,setOpen] = useState(false);
    const [loading,setLoading] = useState(false);

    const title = initialData ? "Edit color" : "Create color"
    const description = initialData ? "Edit color" : "Add a new color"
    const toastMessage = initialData ? "Color updated" : "Color created"
    const action = initialData ? "Save changes" : "Create color"
    
    const form = useForm<ColorFormValues>({
        resolver: zodResolver(formSchema),
        defaultValues: initialData || {
            name:"",
            value:""
        },
    });

    const onSubmit = async (data: ColorFormValues) => {
        try {
            setLoading(true)
            if (initialData) {
                await axios.patch(`/api/${params.storeId}/colors/${params.colorId}`,data)
            } else {
                await axios.post(`/api/${params.storeId}/colors`,data)
            }
            router.push(`/${params.storeId}/colors`)
            router.refresh()
            toast.success(toastMessage);
        } catch (error) {
            toast.error("Something went wrong")}
        finally {
            setLoading(false)
        }
    }

    const onDelette = async () => {
        try {
            setLoading(true)
            await axios.delete(`/api/${params.storeId}/colors/${params.colorId}`)
            router.push(`/${params.storeId}/colors`)
            router.refresh()
            toast.success("Color deleted")
        } catch (error) {
            toast.error("Make sure you removed all products using this colorfirst")
        } finally {
            setLoading(false);
            setOpen(false);
        }
}
  return (
    <>
    <AlertModal
    isOpen={open}
    onClose={()=>setOpen(false)}
    onConfirm={onDelette}
    loading={loading}

    ></AlertModal>
        <div className='flex items-center justify-between'>
                <Heading title={title}
                description= {description}/>
                { initialData && <Button variant = "destructive" size= "icon" onClick={()=>setOpen(true)}>
                    <Trash className='h-4 w-4'></Trash>
                </Button>}
        </div>
        <Separator></Separator>
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8 w-full ">
                <div className='grid grid-cols-3 gap-8'>
                    <FormField 
                    control={form.control}
                    name="name"
                    render={({field}) =>(
                        <FormItem >
                            <FormLabel >Name</FormLabel>
                            <FormControl >
                                <Input  disabled={loading} placeholder="Color name" {...field}></Input>
                            </FormControl>
                        </FormItem>
                    )}>
                    </FormField>
                    <FormField 
                    control={form.control}
                    name="value"
                    render={({field}) =>(
                        <FormItem >
                            <FormLabel >Name</FormLabel>
                            <FormControl >
                                <div className='flex items-center gap-x-4'>
                                <Input  disabled={loading} placeholder="Color value" {...field}></Input>
                                <div 
                                className='border p-4 rounded-full'
                                style = {{ backgroundColor :field.value}}/>
                                </div>
                            </FormControl>
                        </FormItem>
                    )}>
                    </FormField>
                </div>
                <Button disabled = {loading} className='ml-auto' type="submit" >
                    {action}
                </Button>
            </form>
        </Form>
   </>
  )
}

