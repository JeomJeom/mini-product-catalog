/*
*  Protractor support is deprecated in Angular.
*  Protractor is used in this example for compatibility with Angular documentation tools.
*/
import {bootstrapApplication, provideProtractorTestingSupport} from '@angular/platform-browser';
import {AppComponent} from './app/components/app.component';
import {provideRouter} from '@angular/router';
import routeConfig from './app/routes';
import {provideHttpClient, withInterceptorsFromDi} from "@angular/common/http";
import {provideAnimationsAsync} from "@angular/platform-browser/animations/async";
import {providePrimeNG} from "primeng/config";
import Aura from '@primeng/themes/aura';

bootstrapApplication(AppComponent,
  {
    providers: [
      provideProtractorTestingSupport(),
      provideRouter(routeConfig),
      provideHttpClient(withInterceptorsFromDi()),
      provideAnimationsAsync(),
      providePrimeNG({
        theme: {
          preset: Aura
        }
      })
    ]
  }
).catch(err => console.error(err));
