import { UserButton } from '@clerk/nextjs'
import React from 'react'
import MainNav from './main-nav'
import StoreSwitcher from './store-switcher'
import { redirect } from 'next/navigation'
import { auth } from '@clerk/nextjs/server'
import prismaDB from '@/lib/prismadb'

export default async function Navbar() {
  const {userId} = auth()

  if (!userId) {
    redirect("/sign-in")
  }

  const stores = await prismaDB.store.findMany({
    where: {
      userId
    }
  })

  return (
    <div className='border-b'>
      <div className = "flex h-16 items-center px-4">
        <StoreSwitcher items = {stores}></StoreSwitcher>
        <div>
          This will be routes
        </div>
        <MainNav className='mx-6'/>
        <div className='ml-auto flex items-center space-x-4'>
          <UserButton afterSignOutUrl="/" />
        </div>
      </div>
    </div>
  )
}

