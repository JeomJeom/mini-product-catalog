import {Component} from '@angular/core';
import {RouterLink, RouterOutlet} from '@angular/router';

@Component({
  standalone: true,
  selector: 'app-root',
  imports: [
    RouterLink,
    RouterOutlet,
  ],
  templateUrl: `./app.component.html`,

  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'thales-ui';
}
