import { ProductsShowcase } from "@/components";
import { BaseLayout } from "@/layouts/BaseLayout";
import { retailCrm } from "@/lib/server/config";
import { transformAllProductsData } from "@/lib/utils";
import { GetProductsResponse, ShopItem } from "@/types";

const fetchProducts = async (): Promise<ShopItem[]> => {
  try {
    const response = await fetch(`${retailCrm.endpoints.products}?apiKey=${retailCrm.apiKey}`, {
      cache: "force-cache",
    });

    if (!response.ok) {
      throw new Error(`[Shop] Failed to fetch products (HTTP ${response.status})`);
    }

    const data: GetProductsResponse = await response.json();
    const { transformedProducts } = transformAllProductsData(data.products);
    return transformedProducts;
  } catch (error) {
    console.error("[Shop] fetchProducts failed", error);
    return [];
  }
};

const ShopPage = async () => {
  const products = await fetchProducts();

  return (
    <BaseLayout>
      <ProductsShowcase products={products} />
    </BaseLayout>
  );
};

export default ShopPage;
