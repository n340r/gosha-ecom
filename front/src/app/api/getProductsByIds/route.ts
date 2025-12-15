import { NextResponse } from "next/server";
import { fetchProducts } from "@/lib/server/products";

export const dynamic = "force-dynamic";
export const fetchCache = "force-no-store";

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url);
  const ids = searchParams.get("ids");

  if (!ids) {
    return NextResponse.json({ error: "Product IDs are required" }, { status: 400, headers: createCorsHeaders() });
  }

  const idsArray = ids.split(",");

  try {
    const data = await fetchProducts({ ids: idsArray, cache: "no-store" });

    return NextResponse.json(data, { headers: createCorsHeaders() });
  } catch (error) {
    console.error("Error fetching products by IDs:", error);
    return NextResponse.json({ error: "Failed fetching data" }, { status: 500, headers: createCorsHeaders() });
  }
}

function createCorsHeaders() {
  return {
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
    "Access-Control-Allow-Headers": "Content-Type",
  };
}
