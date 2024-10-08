import prismaDB from "@/lib/prismadb";
import { auth } from "@clerk/nextjs/server";
import { NextResponse } from "next/server";

export async function POST(req: Request) {
  try {
    const { userId } = auth();
    const body = await req.json();

    const { name } = body;

    if (!userId) {
      return new NextResponse("Internal error", { status: 401 });
    }

    if (!name) {
      return new NextResponse("Missing required field: name", { status: 400 });
    }

    const store = await prismaDB.store.create({
      data: {
        name,
        userId,
      },
    });

    return NextResponse.json(store);
  } catch (err) {
    console.log("[STORES_POST]", err);
    return new NextResponse("Internal error", { status: 500 });
  }
}

export async function GET() {
  try {
    const stores = await prismaDB.store.findMany();
    console.log(stores)
    return NextResponse.json(stores);
  } catch (err) {
    console.log("[STORES_GET]", err);
    return new NextResponse("Internal error", { status: 500 });
  }
}
