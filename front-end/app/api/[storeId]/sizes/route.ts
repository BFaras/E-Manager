import prismaDB from "@/lib/prismadb";
import { auth } from "@clerk/nextjs/server";
import { NextResponse } from "next/server";

export async function POST(
  req: Request,
  { params }: { params: { storeId: string } }
) {
  try {
    const { userId } = auth();
    const body = await req.json();

    const { name, value } = body;

    if (!userId) {
      return new NextResponse("unAuthenticated", { status: 401 });
    }

    if (!name) {
      return new NextResponse("Name is required :", { status: 400 });
    }

    if (!value) {
      return new NextResponse("Value is required :", { status: 400 });
    }

    if (!params.storeId) {
      return new NextResponse("Store ID is required :", { status: 400 });
    }

    const storeByUserId = await prismaDB.store.findFirst({
      where: {
        id: params.storeId,
        userId,
      },
    });

    if (!storeByUserId) {
      return new NextResponse("Unauthorized", { status: 403 });
    }

    const size = await prismaDB.size.create({
      data: {
        name,
        value,
        storeId: params.storeId,
      },
    });

    return NextResponse.json(size);
  } catch (err) {
    console.log("[SIZE_POST]", err);
    return new NextResponse("Internal error", { status: 500 });
  }
}

export async function GET(
  req: Request,
  { params }: { params: { storeId: string } }
) {
  try {
    if (!params.storeId) {
      return new NextResponse("Store ID is required :", { status: 400 });
    }

    const size = await prismaDB.size.findMany({
      where: {
        storeId: params.storeId,
      },
    });

    return NextResponse.json(size);
  } catch (err) {
    console.log("[SIZE_GET]", err);
    return new NextResponse("Internal error", { status: 500 });
  }
}
