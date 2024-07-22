"use client";

import Heading from '@/components/heading'
import React from 'react'
import { Plus} from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Separator } from '@radix-ui/react-separator'
import { useRouter } from 'next/navigation'
import { useParams } from 'next/navigation'
import { OrderColumn, columns } from './columns';
import { DataTable } from '@/components/ui/data-table';
import ApiList from '@/components/ui/api-list';

interface OrderClientProps {
  data: OrderColumn[]
}
export default function OrderClient({data}:OrderClientProps) {
  const router = useRouter();
  const params = useParams();

  return (
    <>
      <Heading 
      title={`Order (${data.length})`}
      description='Manage orders for your store' />
    <Separator></Separator>
    <DataTable columns={columns} data={data} searchKey="products"></DataTable>
    </>
  )
}

