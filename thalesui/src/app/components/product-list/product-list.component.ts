import {Component, Input} from '@angular/core';
import {CommonModule, NgOptimizedImage} from '@angular/common';
import {ProductView} from '../../models/product.model';
import {RouterModule} from '@angular/router';
import {FormsModule} from "@angular/forms";
import {formatNumberWithCommas} from "../../../utils/utils";

@Component({
  selector: 'app-product-list',
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    NgOptimizedImage
  ],
  templateUrl: `./product-list.component.html`,
  standalone: true,
  styleUrls: ['./product-list.component.css']
})

export class ProductListComponent {
  @Input() product!: ProductView;

  protected readonly formatNumberWithCommas = formatNumberWithCommas;
}
