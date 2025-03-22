import { NextRequest, NextResponse } from "next/server";
import axiosInstance, { setUpInterceptor } from "@/app/utils/axios_instance";
import { auth } from "@clerk/nextjs/server";

export async function middleware(req: NextRequest) {
  try {
    const { userId, getToken } = auth();
    if (!userId) return NextResponse.json({ error: "Unauthorized" }, { status: 401 });

    await setUpInterceptor(getToken); 

    const path = req.nextUrl.searchParams.get("path");
    if (!path) return NextResponse.json({ error: "Missing path" }, { status: 400 });

    const method = req.method;
    const body = method !== "GET" && method !== "DELETE" ? await req.json() : undefined; 

    const response = await axiosInstance({
      method,
      url: path,
      data: body,
    });


    if (response.data && typeof response.data === "object") {
      return NextResponse.json(response.data);
    }

    return new NextResponse(null, { status: response.status }); 
  } catch (error: any) {
    console.error("Proxy API Error:", error.message || error);

    return NextResponse.json(
      { error: error.message || "Something went wrong" },
      { status: error.response?.status || 500 }
    );
  }
}

export const GET = middleware;
export const POST = middleware;
export const PUT = middleware;
export const PATCH = middleware;
export const DELETE = middleware;