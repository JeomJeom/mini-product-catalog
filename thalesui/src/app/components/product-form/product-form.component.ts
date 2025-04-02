import {Component, inject, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {FormsModule, NgForm} from '@angular/forms';
import {ProductService} from '../../services/product.service';
import {ProductMutableAttrs} from '../../models/product.model';
import {InputText} from "primeng/inputtext";
import {SelectOption} from "../../models/common.model";
import {NgIf} from "@angular/common";
import {ThemeTypeService} from "../../services/theme-type.service";
import {v4 as uuidv4} from "uuid";
import {DropdownModule} from "primeng/dropdown";
import {Textarea} from "primeng/textarea";
import {Select} from "primeng/select";

@Component({
  selector: 'app-product-form',
  templateUrl: './product-form.component.html',
  imports: [
    FormsModule,
    InputText,
    DropdownModule,
    NgIf,
    Textarea,
    Select
  ],
  styleUrls: ['./product-form.component.css']
})

export class ProductFormComponent implements OnInit {
  private themeTypeService = inject(ThemeTypeService);
  themeTypeOptions: SelectOption[] = [];
  mode: 'create' | 'edit' = 'create';
  productId?: string;
  product: ProductMutableAttrs = {
    name: '',
    modelNo: '',
    year: new Date().getFullYear(),
    themeType: "",
    categoryType: '',
    imageURL: '',
    price: 0,
    description: ''
  };


  constructor(private productService: ProductService, private router: Router, private route: ActivatedRoute) {
    this.themeTypeService.getThemeTypes().subscribe(resp => {
      for (let name of resp) {
        this.themeTypeOptions.push({label: name, value: uuidv4()})
      }
    })
  }

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.mode = 'edit';
      this.productId = id;
      this.productService.getProductById(this.productId).subscribe(product => {
        this.product = {
          name: product.name,
          modelNo: product.modelNo,
          year: product.year,
          themeType: product.themeType,
          categoryType: product.categoryType,
          imageURL: product.imageURL,
          price: product.price,
          description: product.description
        };
      });
    }
  }

  goBack() {
    this.router.navigate(['/']);
  }

  submitForm(form: NgForm) {
    const id = this.route.snapshot.paramMap.get('id');
    if (form.invalid) return;
    if (this.mode === 'edit' && id) {
      this.productService.updateProduct(id, this.product).subscribe(() => {
        this.router.navigate(['/']);
      });
    } else {
      this.productService.createProduct(this.product).subscribe({
        next: () => this.router.navigate(['/']),
        error: (err) => console.error('Create failed', err)
      });
    }
  }

}
