"use client";
import React, { useState } from 'react'
import { ColorColumn } from './columns'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuTrigger } from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'
import { Copy, Edit, MoreHorizontal, Trash } from 'lucide-react'
import toast from 'react-hot-toast'
import { useParams,useRouter } from 'next/navigation'
import axios from 'axios';
import { AlertModal } from '@/app/modals/alert-modal';

interface CellActionProps {
    data:ColorColumn
}
export default function CellAction({data}:CellActionProps) {

    const router = useRouter()
    const params = useParams()

    const [loading,setLoading] = useState(false)
    const [open,setOpen] = useState(false)

    const onCopy = (id:string) => {
        navigator.clipboard.writeText(id);
        toast.success("color Id copied to the clipboard");
    }
    
    const onDelette = async () => {
        try {
            setLoading(true)
            await axios.delete(`/api/${params.storeId}/colors/${data.id}`)
            router.refresh()
            toast.success("Color deleted")
        } catch (error) {
            toast.error("Make sure you removed all products using this color")
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
    <DropdownMenu>
        <DropdownMenuTrigger asChild>
            <Button variant="ghost" className='h-8 w-8 p-0'>
                <span className='sr-only'>Open menu</span>
                <MoreHorizontal className='h-4 w-4'></MoreHorizontal>
            </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align='end'>
            <DropdownMenuLabel>
                Action
            </DropdownMenuLabel>
            <DropdownMenuItem onClick={()=>onCopy(data.id)}>
                <Copy className='mr-2 h-4 w-4'></Copy>
                Copy Id
            </DropdownMenuItem>
            <DropdownMenuItem onClick={()=> router.push(`/${params.storeId}/colors/${data.id}`)}>
                <Edit className='mr-2 h-4 w-4'></Edit>
                Update
            </DropdownMenuItem>
            <DropdownMenuItem onClick={()=>{setOpen(true)}}>
                <Trash className='mr-2 h-4 w-4'></Trash>
                Delete
            </DropdownMenuItem>


        </DropdownMenuContent>
    </DropdownMenu>
    </>
  )
}

