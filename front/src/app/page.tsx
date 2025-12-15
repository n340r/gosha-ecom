import { ProductsShowcase } from '@/components';
import { BaseLayout } from '@/layouts/BaseLayout';
import { fetchProducts as fetchProductsResponse } from '@/lib/server/products';
import { transformAllProductsData } from '@/lib/utils';
import { GetProductsResponse, ShopItem } from '@/types';

const fetchShopProducts = async (): Promise<ShopItem[]> => {
  try {
    const data: GetProductsResponse = await fetchProductsResponse({ cache: 'force-cache' });
    const { transformedProducts } = transformAllProductsData(data.products);
    return transformedProducts;
  } catch (error) {
    console.error('[Shop] fetchProducts failed', error);
    return [];
  }
};

const ShopPage = async () => {
  const products = await fetchShopProducts();

  return (
    <BaseLayout>
      <ProductsShowcase products={products} />
    </BaseLayout>
  );
};

export default ShopPage;
