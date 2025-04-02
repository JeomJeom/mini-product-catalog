import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {ProductService} from '../../services/product.service';
import {ProductView} from "../../models/product.model";
import {NgIf} from "@angular/common";
import {formatNumberWithCommas} from "../../../utils/utils";

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  imports: [
    NgIf
  ],
  styleUrls: ['./details.component.css']
})
export class DetailsComponent implements OnInit {
  readonly formatNumberWithCommas = formatNumberWithCommas;

  product: Partial<ProductView> = {
    createdAt: "",
    updatedAt: "",
    name: "",
    modelNo: "",
    year: 0,
    themeType: "",
    categoryType: "",
    imageURL: "",
    price: 0,
    description: ""
  };

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private productService: ProductService
  ) {
  }

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.productService.getProductById(id).subscribe({
        next: (res) => {
          this.product = res
          this.product.imageURL = "https://picsum.photos/200"
        },
        error: (err) => console.error('Error loading product', err)
      });
    }
  }


  goBack() {
    this.router.navigate(['/']);
  }

  editProduct() {
    this.router.navigate(['/edit', this.product?.id]);
  }

  deleteProduct() {
    if (this.product?.id && confirm('Are you sure you want to delete this product?')) {
      this.productService.deleteProduct(this.product.id).subscribe(() => {
        this.router.navigate(['/']);
      });
    }
  }

}
