import {Component, inject} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ProductListComponent} from '../product-list/product-list.component';
import {ProductService} from '../../services/product.service';
import {FormsModule} from "@angular/forms";
import {PaginatorModule, PaginatorState} from "primeng/paginator";
import {ButtonModule} from "primeng/button";
import {DividerModule} from "primeng/divider";
import {ProductFilter, ProductView} from "../../models/product.model";
import {Select} from "primeng/select";
import {SelectChangeEvent} from "primeng/select/select.interface";
import {DropdownModule} from "primeng/dropdown";
import {ThemeTypeService} from "../../services/theme-type.service";
import {v4 as uuidv4} from 'uuid';
import {OrderBy} from "../../models/pagination.model";
import {RouterLink} from "@angular/router";
import {SelectOption} from "../../models/common.model";

@Component({
  standalone: true,
  selector: 'app-home',
  imports: [
    CommonModule,
    ProductListComponent,
    FormsModule,
    PaginatorModule, ButtonModule, DividerModule, FormsModule, Select, DropdownModule, RouterLink,
  ],
  templateUrl: `./product.component.html`,
  styleUrls: ['./product.component.css']
})

export class ProductComponent {
  private productService = inject(ProductService);
  private themeTypeService = inject(ThemeTypeService);

  paginatedProducts: ProductView[] = [];
  productFilter: ProductFilter = {
    pageSize: 10,
    pageIndex: 0,
    columnName: "",
    orderBy: "desc",
    keyword: "",
    themeType: "",
  }

  first: number = 0;
  rows: number = 10;
  totalItems: number = 0;
  // Selected options
  selectedSort: OrderBy = "";
  themeTypeValue: string = ""

  themeTypeOptions: SelectOption[] = [
    {label: 'All', value: ''},
  ];

  sortOptions = [
    {label: 'Default', value: ''},
    {label: 'Price: Low to High', value: 'asc'},
    {label: 'Price: High to Low', value: 'desc'},
    // add more sort options here
  ];


  options = [
    {label: 5, value: 5},
    {label: 10, value: 10},
    {label: 20, value: 20},
    {label: 120, value: 120}
  ];

  constructor() {
    this.loadProducts();

    this.themeTypeService.getThemeTypes().subscribe(resp => {
      for (let name of resp) {
        this.themeTypeOptions.push({label: name, value: uuidv4()})
      }
    })
  }

  loadProducts(filter?: ProductFilter) {
    this.productFilter = {...this.productFilter, ...filter}
    this.productService.getProducts(this.productFilter)
      .subscribe((resp) => {
        // TODO: use https://picsum.photos/200 as a sample of product image,
        //  as the dataset is from https://mavenanalytics.io/data-playground?order=date_added%2Cdesc&page=5&pageSize=5
        //  the image url provided is being blocked by the remote server (images.brickset.com).
        this.paginatedProducts = resp.data.map(prod => ({...prod, imageURL: "https://picsum.photos/200"
      }))
        this.first = resp.pagination.pageIndex * resp.pagination.pageSize;
        this.rows = resp.pagination.pageSize ?? 20;
        this.totalItems = resp.pagination.totalItems
      });
  }

  onPageSizeChange(event: SelectChangeEvent) {
    this.rows = event.value;
    this.first = 0;
    this.loadProducts({pageIndex: 0, pageSize: event.value})
  }

  onPageChange(event: PaginatorState) {
    this.first = event.first ?? 0;
    this.rows = event.rows ?? 10;
    this.loadProducts({pageIndex: event.page, pageSize: event.rows})
  }

  onSearchKeyword(value: string) {
    this.loadProducts({keyword: value})
  }

  onThemeTypeChange(): void {
    let themeType = this.themeTypeOptions.find(item => item.value === this.themeTypeValue && this.themeTypeValue !== this.themeTypeOptions[0].value)?.label ?? ""

    this.loadProducts({themeType});
  }

  onSortBy(): void {
    this.loadProducts({columnName: "price", orderBy: this.selectedSort})
  }

}
