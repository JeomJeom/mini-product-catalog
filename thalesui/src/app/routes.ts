import {Routes} from '@angular/router';
import {ProductComponent} from './components/product/product.component';
import {DetailsComponent} from './components/details/details.component';
import {ProductFormComponent} from "./components/product-form/product-form.component";

const routeConfig: Routes = [
  {
    path: '',
    component: ProductComponent,
    title: 'Product page'
  },
  {
    path: 'details/:id',
    component: DetailsComponent,
    title: 'Product details'
  },
  {path: 'create', component: ProductFormComponent, data: {mode: 'create'}, title: 'Add new product'},
  {path: 'edit/:id', component: ProductFormComponent, data: {mode: 'edit'}, title: 'Update new product'},
];

export default routeConfig;
