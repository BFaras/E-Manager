"use client";

import Heading from '@/components/heading'
import React from 'react'
import { Plus} from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Separator } from '@radix-ui/react-separator'
import { useRouter } from 'next/navigation'
import { useParams } from 'next/navigation'
import { Billboard } from '@prisma/client';
import { BillboardColumn, columns } from './columns';
import { DataTable } from '@/components/ui/data-table';

interface BillboardClientProps {
  data: BillboardColumn[]
}
export default function BillboardClient({data}:BillboardClientProps) {
  const router = useRouter();
  const params = useParams();

  return (
    <>
    <div className='flex items-center justify-between '>
      <Heading 
      title={`Billboard (${data.length})`}
      description='Manage billboards for your store' />
      <Button onClick={()=> router.push(`/${params.storeId}/billboards/new`)}>
        <Plus className ="mr-2 h-4 w-4">
            Add new 
        </Plus>
      </Button>
    </div>
    <Separator></Separator>
    <DataTable columns={columns} data={data} searchKey="label"></DataTable>
      
    </>
  )
}

