import { UserButton } from '@clerk/nextjs'
import React from 'react'
import MainNav from './main-nav'
import StoreSwitcher from './store-switcher'
import { redirect } from 'next/navigation'
import { auth } from '@clerk/nextjs/server'
import { ThemeToggle } from './theme-toggle'
import axiosInstance from '@/app/utils/axios_instance'
import toast from 'react-hot-toast'
import { Store } from '@/models/db'

async function fetchStore(userId: string): Promise<Store[]> {
  try {
      const response = await axiosInstance.get(`users/${userId}/stores`)
      return response.data as Store[];
  } catch (err) {
      toast.error("Error trying to get all stores for navigation ");
  }
  return []
}


export default async function Navbar() {
  const {userId} = auth()

  if (!userId) {
    redirect("/sign-in")
  }

  let stores: Store[] = await fetchStore(userId)

  return (
    <div className='border-b'>
      <div className = "flex h-16 items-center px-4">
        <StoreSwitcher items = {stores}></StoreSwitcher>
        <MainNav className='mx-6'/>
        <div className='ml-auto flex items-center space-x-4'>
          <ThemeToggle></ThemeToggle>
          <UserButton afterSignOutUrl="/" />
        </div>
      </div>
    </div>
  )
}

