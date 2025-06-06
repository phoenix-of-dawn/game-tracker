import { Component } from '@angular/core';
import { RegisterComponent } from '../../components/register/register.component';

@Component({
    selector: 'register-page',
    templateUrl: 'register-page.component.html',
    imports: [RegisterComponent],
})
export class RegisterPage {
    constructor() {}
}
