import { Injectable } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';

@Injectable({
    providedIn: 'root',
})
export class ErrorHandler {
    constructor(private snackBar: MatSnackBar) {}

    errorDuration = 3000;

    showError(error: string) {
        this.snackBar.open(error, '', {
            verticalPosition: 'top',
            horizontalPosition: 'center',
            duration: this.errorDuration,
        });
    }
}
