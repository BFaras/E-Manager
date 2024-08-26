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

    const { label, imageUrl,isActive } = body;

    if (!userId) {
      return new NextResponse("unAuthenticated", { status: 401 });
    }

    if (!label) {
      return new NextResponse("Label is required :", { status: 400 });
    }

    if (!label) {
      return new NextResponse("Image Url is required :", { status: 400 });
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
    /*make sure this one has isActive to create it */

    const billboard = await prismaDB.billboard.create({
      data: {
        label,
        imageUrl,
        storeId: params.storeId
      },
    });

    return NextResponse.json(billboard);
  } catch (err) {
    console.log("[BILLBOARD_POST]", err);
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

    const billboard = await prismaDB.billboard.findMany({
      where: {
        storeId: params.storeId,
      },
    });

    return NextResponse.json(billboard);
  } catch (err) {
    console.log("[BILLBOARD_POST]", err);
    return new NextResponse("Internal error", { status: 500 });
  }
}
