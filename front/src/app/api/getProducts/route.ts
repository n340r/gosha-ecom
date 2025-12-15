import { fetchProducts } from "@/lib/server/products";
import { GetProductsResponse } from "@/types";
import { NextResponse } from "next/server";

export const dynamic = "force-dynamic";
export const fetchCache = "force-no-store";

export async function GET(request: Request): Promise<NextResponse> {
  try {
    const data: GetProductsResponse = await fetchProducts({ cache: "no-store" });

    return NextResponse.json(data, {
      status: 200,
    });
  } catch (error) {
    console.error("❌ Error fetching products:", error);
    return NextResponse.json(
      { error: "Failed fetching data" },
      {
        status: 500,
      },
    );
  }
}
