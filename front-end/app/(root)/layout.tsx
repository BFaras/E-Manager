import prismaDB from "@/lib/prismadb";
import { auth } from "@clerk/nextjs/server";
import { redirect } from "next/navigation";
import axiosInstance from "../utils/axios_instance";


export default async function SetupLayout({
    children
} : {
    children: React.ReactNode;
}) {
    const {userId} = auth();

    if (!userId) {
        redirect('/sign-in')
    }

    console.log("test golang please work")

    try {
        const response = await axiosInstance.get(`/stores/user/${userId}`)
        console.log("test golang please work")
        console.log(response)

    } catch (err) {
        console.log("help with error")
        console.log(err)
    }

    const store = await prismaDB.store.findFirst({
        where: {
            userId
        }
    })

    if (store){
        redirect(`/${store.id}`)
    }

    return(
    <>
        {children}
    </>
    )

}