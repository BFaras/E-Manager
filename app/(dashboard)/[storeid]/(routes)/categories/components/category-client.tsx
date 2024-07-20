"use client";

import Heading from '@/components/heading'
import React from 'react'
import { Plus} from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Separator } from '@radix-ui/react-separator'
import { useRouter } from 'next/navigation'
import { useParams } from 'next/navigation'
import { CategoryColumn, columns } from './columns';
import { DataTable } from '@/components/ui/data-table';
import ApiList from '@/components/ui/api-list';

interface CategoryClientProps {
  data: CategoryColumn[]
}
export default function CategoryClient({data}:CategoryClientProps) {
  const router = useRouter();
  const params = useParams();

  return (
    <>
    <div className='flex items-center justify-between '>
      <Heading 
      title={`Categories (${data.length})`}
      description='Manage categories for your store' />
      <Button onClick={()=> router.push(`/${params.storeId}/categories/new`)}>
        <Plus className ="mr-2 h-4 w-4">
            Add new 
        </Plus>
      </Button>
    </div>
    <Separator></Separator>
    <DataTable columns={columns} data={data} searchKey="name"></DataTable>
    <Heading title="API" description="API calls for Categories"></Heading>
    <Separator></Separator>
    <ApiList entityName='categories' entityIdName="categoryId"></ApiList>

    </>
  )
}

