
import Heading from '@/components/heading'
import { Separator } from '@radix-ui/react-separator'
import React from 'react'
import {Card, CardContent, CardHeader, CardTitle} from '@/components/ui/card'
import { formatter } from '@/lib/utils'
import { Overview } from '@/components/overview'
import axiosInstance from '@/app/utils/axios_instance'

interface DashboardPageProps {
  params: {storeId: string}
}

export interface GraphData {
  Name:string;
  Total:number;
}

export default async function DashboardPage({ params }: DashboardPageProps) {

  const responseRevenue  = await axiosInstance.get(`stores/${params.storeId}/revenue`)
  const totalRevenue:number = responseRevenue.data
  const reponseSales = await axiosInstance.get(`stores/${params.storeId}/sales`)
  const salesCount:number = reponseSales.data
  const responseGraphRevenue = await axiosInstance.get(`stores/${params.storeId}/graphRevenue`)
  const graphRevenue:GraphData[] = responseGraphRevenue.data

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
