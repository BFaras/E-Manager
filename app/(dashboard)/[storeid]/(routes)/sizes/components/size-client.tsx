"use client";

import Heading from '@/components/heading'
import React from 'react'
import { Plus} from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Separator } from '@radix-ui/react-separator'
import { useRouter } from 'next/navigation'
import { useParams } from 'next/navigation'
import { SizeColumn, columns } from './columns';
import { DataTable } from '@/components/ui/data-table';
import ApiList from '@/components/ui/api-list';

interface SizeClientProps {
  data: SizeColumn[]
}
export default function SizeClient({data}:SizeClientProps) {
  const router = useRouter();
  const params = useParams();

  return (
    <>
    <div className='flex items-center justify-between '>
      <Heading 
      title={`Sizes (${data.length})`}
      description='Manage sizes for your store' />
      <Button onClick={()=> router.push(`/${params.storeId}/sizes/new`)}>
        <Plus className ="mr-2 h-4 w-4">
            Add new 
        </Plus>
      </Button>
    </div>
    <Separator></Separator>
    <DataTable columns={columns} data={data} searchKey="name"></DataTable>
    <Heading title="API" description="API calls for Sizes"></Heading>
    <Separator></Separator>
    <ApiList entityName='sizes' entityIdName="sizeId"></ApiList>

    </>
  )
}

