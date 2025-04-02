export type OrderBy = 'desc' | 'asc' | ""

export interface PageInfo {
  pageIndex?: number;
  pageSize?: number;
  columnName?: string;
  orderBy?: OrderBy;
}

export interface Pagination {
  pageIndex: number;
  pageSize: number;
  totalItems: number;
  totalPages: number;
}
