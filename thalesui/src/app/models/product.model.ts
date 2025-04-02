import {PageInfo, Pagination} from "./pagination.model";

export interface ProductView {
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  modelNo: string;
  year: number;
  themeType: string;
  categoryType: string;
  imageURL: string;
  price: number;
  description: string;
}

export interface PaginatedProductResp {
  data: ProductView[];
  pagination: Pagination;
}

export interface ProductMutableAttrs {
  name: string;
  modelNo: string;
  year: number;
  themeType?: string;
  categoryType: string;
  imageURL: string;
  price: number;
  description: string;
}

export interface ProductFilter extends PageInfo {
  themeType?: string
  keyword?: string
}
