import {keyboardState} from "@testing-library/user-event/dist/keyboard/types";

export type Task = {
    id :number,
    title : string,
    created_at : Date,
    updated_at : Date
}

export type CsrfToken = {
    csrf_token : string
}

export type Credential = {
    email : string
    password : string
}