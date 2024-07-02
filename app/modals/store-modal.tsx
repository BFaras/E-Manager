"use client";

import { Button } from '@/components/ui/button';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import Modal from '@/components/ui/modal';
import { useStoreModal } from '@/hooks/use-store-modal';
import { zodResolver } from '@hookform/resolvers/zod';
import React, { useState } from 'react'
import {  useForm } from 'react-hook-form';
import * as z from 'zod'
import axios from 'axios'
import {toast} from "react-hot-toast"


const formSchema = z.object({
    name: z.string().min(1),
})
export function StoreModal() {
    const storeModal = useStoreModal();

    const [loading, setLoading] = useState(false)

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: '',
        },
    });

    const onSubmit = async (values:z.infer<typeof formSchema>) => {
        try {
            setLoading(true);

            const response = await axios.post('/api/stores',values)

            toast.success("Store created successfully")
        } catch (error) {
            toast.error("something wwent wrontg")
        } finally {
            setLoading(false);
        }
    }
    
  return (
    <Modal
    title="Create a new store"
    description='Add a new store to manage products'
    isOpen={storeModal.isOpen}
    onClose={storeModal.onClose}
    >
    <div>
      <div className='space-y-4 py-2 pb-4'>
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)}>
                <FormField control={form.control}
                 name = "name"
                 render = {({field })=>(
                    <FormItem>
                        <FormLabel>Name</FormLabel>
                        <FormControl>
                            <Input disabled = {loading}
                             placeholder="E-Commerce"
                              {...field}></Input>
                        </FormControl>
                        <FormMessage></FormMessage>
                    </FormItem>
                 )}>
                </FormField>
                <div className='pt-6 space-x-2 flex items-center justify-end '>
                    <Button 
                    disabled ={loading}
                    variant="outline"
                    onClick={storeModal.onClose}>Cancel</Button>
                    <Button disabled ={loading} 
                    type="submit">Continue</Button>
                </div>
            </form>
        </Form>
      </div>
    </div>
    </Modal>
  )
}
