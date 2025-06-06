import { Component } from '@angular/core';
import {
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormField } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';

@Component({
    selector: 'register-component',
    templateUrl: 'register.component.html',
    styleUrl: 'register.component.scss',
    imports: [
        MatFormField,
        MatInputModule,
        ReactiveFormsModule,
        MatButtonModule,
    ],
})
export class RegisterComponent {
    constructor() {}

    registerForm = new FormGroup({
        email: new FormControl('', [Validators.email, Validators.required]),
        password: new FormControl('', Validators.required),
    });

    onSubmit() {
        console.log('submitted!');
    }
}
