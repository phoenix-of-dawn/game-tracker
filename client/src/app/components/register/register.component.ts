import { Component } from '@angular/core';
import {
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatError, MatFormField } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { RegistrationHandler } from '../../services/registeration-handler.service';
import { Router } from '@angular/router';
import { UserCreationError } from '../../interfaces/user';

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
    constructor(
        private registrationHandler: RegistrationHandler,
        private router: Router
    ) {}

    registerForm = new FormGroup({
        email: new FormControl('', [Validators.email, Validators.required]),
        password: new FormControl('', [
            Validators.required,
            Validators.minLength(6),
        ]),
    });

    get email() {
        return this.registerForm.get('email');
    }

    get password() {
        return this.registerForm.get('password');
    }

    onSubmit() {
        console.log('submitted!');
        if (this.registerForm.value.email && this.registerForm.value.password) {
            const result = this.registrationHandler.makeUser(
                this.registerForm.value.email,
                this.registerForm.value.password
            );
            this.registerForm.reset();

            if (result.Error == null) {
                // Move to user page when implemented
                this.router.navigate(['/search']);
            } else {
                this.router.navigate(['/register'], {
                    queryParams: { Error: result.Error },
                });
            }
        }
    }
}
