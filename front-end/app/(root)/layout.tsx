
import { auth } from "@clerk/nextjs/server";
import { redirect } from "next/navigation";
import axiosInstance from "../utils/axios_instance";
import { Store } from "@/models/db";

async function fetchStore(userId: string) {
    try {
        const response = await axiosInstance.get(`/users/${userId}/store`)
        return response.data;

    } catch (err: any) {
        //TO DO: redirect to a 500 page 
    }

}

export default async function SetupLayout({
    children
} : {
    children: React.ReactNode;
}) {
    const {userId} = auth();

    if (!userId) {
        redirect('/sign-in')
    }
    let store: Store = await fetchStore(userId);

    if (store){
        redirect(`/${store.id}`)
    }

    return(
    <>
        {children}
    </>
    )

}