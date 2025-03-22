import { stripe } from "@/lib/stripe";
import { NextResponse } from "next/server";
import axiosInstance from "@/app/utils/axios_instance";
import Stripe from "stripe";

const corsHeaders = {
  "Access-Control-Allow-Origin": "http://localhost:3002",
  "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, PATCH, OPTIONS",
  "Access-Control-Allow-Headers": "Content-Type, Authorization",
};

export async function OPTIONS(
  req: Request,
  { params }: { params: { storeId: string } }
) {
  return new NextResponse(null, {
    status: 204,
    headers: corsHeaders,
  });
}

export async function POST(
  req: Request,
  { params }: { params: { storeId: string } }
) {
  const { productsId } = await req.json();
  const storeId = params.storeId;

  if (!productsId || productsId.length === 0) {
    return new NextResponse("Products IDs are required", {
      status: 400,
      headers: corsHeaders,
    });
  }

  try {
    const productPromises = productsId.map((id: string) =>
      axiosInstance.get(`/stores/${storeId}/products/${id}`).then((res) => res.data)
    );
    const products = await Promise.all(productPromises);

    const listItems: Stripe.Checkout.SessionCreateParams.LineItem[] = products.map((product: any) => ({
      quantity: 1,
      price_data: {
        currency: "USD",
        product_data: { name: product.name },
        unit_amount: product.price * 100,
      },
    }));

    const { data: order } = await axiosInstance.post(`/stores/${storeId}/orders`, {
      productsId,
    });

    const session = await stripe.checkout.sessions.create({
      line_items: listItems,
      mode: "payment",
      billing_address_collection: "required",
      phone_number_collection: { enabled: true },
      success_url: `${process.env.FRONTEND_STORE_URL}/cart?success=1`,
      cancel_url: `${process.env.FRONTEND_STORE_URL}/cart?canceled=1`,
      metadata: { orderId: order.id },
    });

    return new NextResponse(JSON.stringify({ url: session.url }), {
      status: 200,
      headers: corsHeaders,
    });
  } catch (error) {
    console.error("Checkout error:", error);
    return new NextResponse("Checkout failed", {
      status: 500,
      headers: corsHeaders,
    });
  }
}
