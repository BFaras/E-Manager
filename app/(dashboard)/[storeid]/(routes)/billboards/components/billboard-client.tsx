"use client";

import Heading from '@/components/heading'
import React from 'react'
import {Minus, Plus} from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Separator } from '@radix-ui/react-separator'
import { useRouter } from 'next/navigation'
import { useParams } from 'next/navigation'
import { firebaseDB } from '@/app/utils/firebaseConfig';
import { ref, uploadBytes } from 'firebase/storage';

require('dotenv').config()

export default function BillboardClient() {
  const router = useRouter();
  const params = useParams();

  return (
    <>
    <div className='flex items-center justify-between '>
      <Heading 
      title="Billboard (0)"
      description='Manage billboards for your store' />
      <Button onClick={()=> router.push(`/${params.storeId}/billboards/new`)}>
        <Plus className ="mr-2 h-4 w-4">
            Add new 
        </Plus>
      </Button>
    </div>
    <Separator></Separator>
      
    </>
  )
}

