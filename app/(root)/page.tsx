import Image from "next/image";
import {Button } from '@/components/ui/button'
import { UserButton } from "@clerk/nextjs";

export default function SetUpPage() {
  return (
   <div>
    <div className="p-4">
    <UserButton afterSignOutUrl="/sign-in"></UserButton>
    </div>
   </div>
  );
}
