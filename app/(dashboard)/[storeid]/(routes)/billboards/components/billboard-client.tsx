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

  function getName(){
    console.log(process.env.NEXT_PUBLIC_PROJECT_ID)
  }

  async function upload(){

    console.log("help")
    console.log(firebaseDB)
    console.log(process.env.STORAGE_BUCKET)
    var storageRef = ref(firebaseDB,'public')
    
    const response = await fetch('https://res.cloudinary.com/demo/image/upload/v1689803100/ai/hiker.jpg') 

    console.log("response", response)
    const file = await response.blob()
    
    console.log("resfileponse", file)

    uploadBytes(storageRef, file).then((snapshot) => {
        console.log('Uploaded a blob or file!');
      });

}
  return (
    <>
    <div className='flex items-center justify-between '>
      <Heading 
      title="Billboard (0)"
      description='Manage billboards for your store' />
      <Button onClick={getName}>
        <Minus className ="mr-2 h-4 w-4">
            Add new 
        </Minus>
      </Button>

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

