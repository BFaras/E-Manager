"use client";
import React from 'react'
import { BillboardColumn } from './columns'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuTrigger } from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'
import { Copy, Edit, MoreHorizontal, Trash } from 'lucide-react'
import toast from 'react-hot-toast'
import { useParams,useRouter } from 'next/navigation'

interface CellActionProps {
    data:BillboardColumn
}
export default function CellAction({data}:CellActionProps) {

    const router = useRouter()
    const params = useParams()

    const onCopy = (id:string) => {
        navigator.clipboard.writeText(id);
        toast.success("Billboard Id copied to the clipboard");
    }
  return (
    <DropdownMenu>
        <DropdownMenuTrigger >
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
            <DropdownMenuItem onClick={()=> router.push(`/${params.storeId}/billboards/${data.id}`)}>
                <Edit className='mr-2 h-4 w-4'></Edit>
                Update
            </DropdownMenuItem>
            <DropdownMenuItem>
                <Trash className='mr-2 h-4 w-4'></Trash>
                Delete
            </DropdownMenuItem>


        </DropdownMenuContent>
    </DropdownMenu>
  )
}

