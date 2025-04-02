import {Injectable} from '@angular/core';
import {PaginatedProductResp, ProductFilter, ProductMutableAttrs, ProductView} from '../models/product.model';
import {HttpClient, HttpParams} from "@angular/common/http";
import {Observable} from "rxjs";
import {environment} from "../../env/env";

@Injectable({
  providedIn: 'root'
})
export class ProductService {
  private apiUrl = environment.apiUrl + '/products';

  constructor(private http: HttpClient) {
  }

  getProducts(filter?: ProductFilter): Observable<PaginatedProductResp> {
    let params = new HttpParams()
    if (filter) {
      if (filter.keyword) {
        params = params.set("keyword", filter.keyword)
      }
      if (filter.themeType) {
        params = params.set("themeType", filter.themeType)
      }

      if (filter.pageIndex) {
        params = params.set('pageIndex', filter.pageIndex);
      }

      if (filter.pageSize) {
        params = params.set('pageSize', filter.pageSize);
      }

      if (filter.columnName) {
        params = params.set('columnName', filter.columnName);
      }

      if (filter.orderBy) {
        params = params.set('orderBy', filter.orderBy); // e.g., 'asc' or 'desc'
      }
    }

    return this.http.get<PaginatedProductResp>(this.apiUrl, {params})
  }

  createProduct(product: Partial<ProductView>): Observable<ProductView> {
    return this.http.post<ProductView>(this.apiUrl, product);
  }

  updateProduct(id: string, product: Partial<ProductMutableAttrs>): Observable<ProductMutableAttrs> {
    return this.http.put<ProductMutableAttrs>(`${this.apiUrl}/${id}`, product);
  }

  getProductById(id: string): Observable<ProductView> {
    return this.http.get<ProductView>(`${this.apiUrl}/${id}`);
  }

  deleteProduct(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`);
  }

  submitApplication(firstName: string, lastName: string, email: string) {
    console.log(`Homes application received: firstName: ${firstName}, lastName: ${lastName}, email: ${email}.`);
  }
}
