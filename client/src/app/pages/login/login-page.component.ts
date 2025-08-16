import { Component } from '@angular/core';
import { LoginComponent } from '../../components/login/login.component';

@Component({
    selector: 'login-page',
    templateUrl: 'login-page.component.html',
    styleUrl: 'login-page.component.scss',
    imports: [LoginComponent],
})
export class LoginPage {}
