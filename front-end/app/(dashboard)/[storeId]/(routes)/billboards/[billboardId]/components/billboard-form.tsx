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
import { ApiAlert } from '@/components/ui/api-alert';
import { useOrigin } from '@/hooks/use-origin';
import { Billboard } from '@prisma/client';
import ImageUpload from '@/components/ui/image-upload';


const formSchema = z.object({
    label: z.string().min(1),
    imageUrl: z.string().min(1)
})

type BillboardFormValues = z.infer<typeof formSchema>

interface BillboardFormProps {
    initialData: Billboard | null
}

export default function BillboardForm({initialData}:BillboardFormProps) {

    const params = useParams()
    const router = useRouter()

    const [open,setOpen] = useState(false);
    const [loading,setLoading] = useState(false);

    const title = initialData ? "Edit billboard" : "Create billboard"
    const description = initialData ? "Edit billboard" : "Add a new billboard"
    const toastMessage = initialData ? "Billboard updated" : "Billboard created"
    const action = initialData ? "Save changes" : "Create billboard"
    
    const form = useForm<BillboardFormValues>({
        resolver: zodResolver(formSchema),
        defaultValues: initialData || {
            label:"",
            imageUrl:""
        },
    });

    const onSubmit = async (data: BillboardFormValues) => {
        try {
            setLoading(true)
            if (initialData) {
                await axios.patch(`/api/${params.storeId}/billboards/${params.billboardId}`,data)
            } else {
                await axios.post(`/api/${params.storeId}/billboards`,data)
            }
            router.push(`/${params.storeId}/billboards`)
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
            await axios.delete(`/api/${params.storeId}/billboards/${params.billboardId}`)
            router.push(`/${params.storeId}/billboards`)
            router.refresh()
            toast.success("Billboard deleted")
        } catch (error) {
            toast.error("Make sure you removed all categories using billboard")
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
                { initialData && <Button variant = "destructive" size= "icon"onClick={()=>setOpen(true)}>
                    <Trash className='h-4 w-4'></Trash>
                </Button>}
        </div>
        <Separator></Separator>
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8 w-full ">
                <FormField 
                        control={form.control}
                        name="imageUrl"
                        render={({field}) =>(
                            <FormItem >
                                <FormLabel >Background image</FormLabel>
                                <FormControl >
                                    <ImageUpload  
                                    value = {field.value ? [field.value] : [] }
                                    disabled = {loading}
                                    onChange={(url) => field.onChange(url)}
                                    onRemove={() => field.onChange("")}>

                                    </ImageUpload>
                                </FormControl>
                            </FormItem>
                        )}>
                    </FormField>
                <div className='grid grid-cols-3 gap-8'>
                    <FormField 
                    control={form.control}
                    name="label"
                    render={({field}) =>(
                        <FormItem >
                            <FormLabel >Name</FormLabel>
                            <FormControl >
                                <Input  disabled={loading} placeholder="Billboard label" {...field}></Input>
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

