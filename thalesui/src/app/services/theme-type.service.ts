import {Injectable} from "@angular/core";
import {environment} from "../../env/env";
import {Observable} from "rxjs";
import {HttpClient} from "@angular/common/http";

@Injectable({providedIn: 'root'})

export class ThemeTypeService {
  private apiUrl = environment.apiUrl + '/theme-types';

  constructor(private http: HttpClient) {
  }

  getThemeTypes(): Observable<string[]> {
    return this.http.get<string[]>(this.apiUrl)
  }
}
