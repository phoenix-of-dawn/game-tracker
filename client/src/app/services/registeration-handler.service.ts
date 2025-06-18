import { Injectable } from '@angular/core';
import { ErrorHandler } from './errors.service';
import { HttpHandler } from './http-handler.service';
import { catchError, Observable, Subject } from 'rxjs';
import { HttpErrorResponse, HttpStatusCode } from '@angular/common/http';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';

@Injectable({ providedIn: 'root' })
export class RegistrationHandler {
    constructor(
        private httpHandler: HttpHandler,
        private errorHandler: ErrorHandler,
        private router: Router,
        private snackBar: MatSnackBar
    ) {}

    public makeUser(email: string, password: string, username: string) {
        const response = this.httpHandler.postRegisterRequest(
            email,
            password,
            username
        );
        response
            .pipe(
                catchError((err: HttpErrorResponse) => {
                    if (err.status == HttpStatusCode.Unauthorized) {
                        this.errorHandler.showError(
                            'An account with this email already exists'
                        );
                    } else {
                        console.log(this.errorHandler);
                        this.errorHandler.showError('Server error');
                        console.log(this.errorHandler);

                        console.log('Server error, snackbar should appear');
                    }
                    throw err;
                })
            )
            .subscribe(() => {
                // TODO: Initiate login afterwards once login is implemented
                // For now, redirect to root
                this.router.navigate(['/']);
            });
    }
}
