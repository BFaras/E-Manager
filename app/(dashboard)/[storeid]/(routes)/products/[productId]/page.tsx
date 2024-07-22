import prismaDB from '@/lib/prismadb'
import React from 'react'
import ProductForm from './components/product-form';

export default async function ProductPage( {params} :
    {params:{productId:string,storeId:string}}) {

  const product = await prismaDB.product.findUnique({
      where:{
          id: params.productId
      },
      include: {
        images: true,
      }
  });

  const categories = await prismaDB.category.findMany({
    where:{
      storeId: params.storeId
    }
});

  const sizes = await prismaDB.size.findMany({
    where:{
      storeId: params.storeId
    }
  });

  const colors = await prismaDB.color.findMany({
    where:{
      storeId: params.storeId
    }
  });

  return (
    <div className="flex-col">
      <div className = 'flex-1 space-y-4 p-8 pt-6'>
        <ProductForm initialData ={product}
        categories={categories}
        colors={colors}
        sizes={sizes}
        ></ProductForm>

      </div>
    </div>
  )
}

