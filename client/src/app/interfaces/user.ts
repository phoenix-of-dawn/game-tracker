export interface UserCreationResponse {
    Authorization: string;
    Error: UserCreationError | null;
}

export enum UserCreationError {
    EmailNotUnique,
    ServerError,
}
