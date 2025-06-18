import { Component } from '@angular/core';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { RouterLink, RouterOutlet } from '@angular/router';

@Component({
    selector: 'app-root',
    imports: [RouterOutlet, RouterLink, MatSnackBarModule],
    templateUrl: './app.component.html',
    styleUrl: './app.component.scss',
})
export class AppComponent {
    title = 'app';
}
