
import Heading from '@/components/heading'
import prismaDB from '@/lib/prismadb'
import { Separator } from '@radix-ui/react-separator'
import React from 'react'
import {Card, CardContent, CardHeader, CardTitle} from '@/components/ui/card'
import { CreditCard, DollarSign } from 'lucide-react'
import { formatter } from '@/lib/utils'
import { getTotalRevenue } from '@/app/actions/get-total-revenue'
import { getSalesCount } from '@/app/actions/get-sales-count'
import { getStockCount } from '@/app/actions/get-stock-count'
import { Overview } from '@/components/overview'
import { getGraphRevenue } from '@/app/actions/get-graph-revenue'

interface DashboardPageProps {
    params: {storeId: string}
}

export default async function DashboardPage({ params }: DashboardPageProps) {

  const totalRevenue = await getTotalRevenue(params.storeId)

  const salesCount = await getSalesCount(params.storeId)

  const graphRevenue = await getGraphRevenue(params.storeId)


  return (
    <div className='flex-col'>
      <div className='flex-1 space-y-4 p-8 pt-6'>
        <Heading title = "Dashboard" description="Overview of your store"/>
          <Separator></Separator>
          <div className='grid gap-4 grid-cols-2'>
            <Card>
              <CardHeader className='flex flex-row items-center justify-between space-y-0 pb-2 '>
                <CardTitle  className='text-lg font-bold'>Total Revenue</CardTitle>
              </CardHeader>
              <CardContent>
                  <div className='text-2xl font-bold'>
                    {formatter.format(totalRevenue)}
                  </div>
              </CardContent>
            </Card>
            <Card>
              <CardHeader className='flex flex-row items-center justify-between space-y-0 pb-2 '>
                <CardTitle  className='text-lg font-bold'>Total Sales</CardTitle>
              </CardHeader>
              <CardContent>
                  <div className='text-2xl font-bold'>
                    {salesCount}
                  </div>
              </CardContent>
            </Card>
          </div>
          <Card>
              <CardHeader className='flex flex-row items-center justify-center space-y-0 pb-4 '>
              <CardTitle  className='text-3xl font-bold'>Overview</CardTitle>
              </CardHeader>
              <CardContent>
                  <div className='pl-2'>
                    <Overview data = {graphRevenue}></Overview>
                  </div>
              </CardContent>
            </Card>
      </div>
    </div>
  )
}
