
import axiosInstance from "@/app/utils/axios_instance";
import Navbar from "@/components/navbar";
import prismaDB from "@/lib/prismadb";
import { Store } from "@/models/db";
import { auth } from "@clerk/nextjs/server";
import axios from "axios";
import {redirect} from "next/navigation";
import toast from "react-hot-toast";

async function fetchStore(userId:string,storeId: string) {
    try {
        console.log("TRy to get store")
        const response = await axiosInstance.get(`/users/${userId}/stores/${storeId}`)
        console.log(response.data);
        return response.data;
    } catch (err) {
        console.log("Failed to get store")
    }
}

export default async function DashboardLayout({
    children,
    params
}: {
    children: React.ReactNode;
    params: {storeId:string}
}){
    const {userId} = auth();

    if(!userId) {
        redirect('/sign-in')
    }

    let store: Store = await fetchStore(userId,params.storeId);

    if (!store) {
        redirect('/')
    }


    return (
        <>
        <Navbar></Navbar>
        {children}
        </>
    )
}