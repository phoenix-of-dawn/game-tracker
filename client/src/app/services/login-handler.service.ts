import { Router } from '@angular/router';
import { ErrorHandler } from './errors.service';
import { HttpHandler } from './http-handler.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Injectable } from '@angular/core';
import { catchError } from 'rxjs';
import { HttpErrorResponse, HttpStatusCode } from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class LoginHandler {
    constructor(
        private httpHandler: HttpHandler,
        private errorHandler: ErrorHandler,
        private router: Router,
        private snackBar: MatSnackBar
    ) {}

    public loginUser(email: string, password: string) {
        const response = this.httpHandler
            .postLoginRequest(email, password)
            .pipe(
                catchError((err: HttpErrorResponse) => {
                    this.errorHandler.showError(err.status.toString());
                    throw err;
                })
            )
            .subscribe(() => this.router.navigate(['/']));
    }
}
