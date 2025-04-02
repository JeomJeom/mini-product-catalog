import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';

describe('AppComponent', () => {
  let component: AppComponent;
  let fixture: ComponentFixture<AppComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppComponent,
        RouterTestingModule
      ]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AppComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create the app component', () => {
    expect(component).toBeTruthy();
  });

  it(`should have as title 'thales-ui'`, () => {
    expect(component.title).toEqual('thales-ui');
  });

  it('should render the brand link with logo', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    const brandLink = compiled.querySelector('.brand');
    expect(brandLink).toBeTruthy();
    // Check that the brand link has the correct routerLink binding
    // Note: Angular adds an attribute like "ng-reflect-router-link" in development mode.
    expect(brandLink?.getAttribute('ng-reflect-router-link')).toBe('/');
    const logo = brandLink?.querySelector('.brand-logo');
    expect(logo).toBeTruthy();
    expect(logo?.getAttribute('src')).toContain('LEGO_logo.svg');
  });

  it('should render the router outlet', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    const outlet = compiled.querySelector('router-outlet');
    expect(outlet).toBeTruthy();
  });
});
