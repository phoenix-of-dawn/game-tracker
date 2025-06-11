import { Injectable } from '@angular/core';
import { HttpHandler } from './http-handler.service';
import { Observable } from 'rxjs';
import { UserCreationError, UserCreationResponse } from '../interfaces/user';

@Injectable({ providedIn: 'root' })
export class RegistrationHandler {
    constructor(private httpHandler: HttpHandler) {}

    public makeUser(email: string, password: string): UserCreationResponse {
        let result: UserCreationResponse = {
            Authorization: '',
            Error: UserCreationError.ServerError,
        };
        const response = this.httpHandler.postRegisterRequest(email, password);
        response.subscribe((res) => {
            result = res;
        });

        return result;
    }
}
