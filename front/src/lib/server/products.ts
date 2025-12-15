import "server-only";

import { GetProductsResponse } from "@/types";

import { retailCrm } from "./config";

type FetchProductsOptions = RequestInit & {
  ids?: Array<string | number>;
};

function isDev() {
  return process.env.NODE_ENV !== "production";
}

function devCatalogProductsUrl(ids?: string[]) {
  const baseUrl = process.env.CATALOG_SERVICE_BASE_URL ?? "http://localhost:8080";
  const url = new URL("/products", baseUrl);

  if (ids?.length) {
    url.searchParams.set("ids", ids.join(","));
  }

  return url.toString();
}

function retailCrmProductsUrl(ids?: string[]) {
  const url = new URL(retailCrm.endpoints.products);
  url.searchParams.set("apiKey", retailCrm.apiKey);

  ids?.forEach((id) => {
    url.searchParams.append("filter[ids][]", id);
  });

  return url.toString();
}

export async function fetchProducts(options: FetchProductsOptions = {}): Promise<GetProductsResponse> {
  const { ids, ...init } = options;
  const idStrings = ids?.map(String).filter(Boolean);

  const url = isDev() ? devCatalogProductsUrl(idStrings) : retailCrmProductsUrl(idStrings);
  const response = await fetch(url, init);

  if (!response.ok) {
    throw new Error(`[products] Failed to fetch (HTTP ${response.status})`);
  }

  return (await response.json()) as GetProductsResponse;
}
