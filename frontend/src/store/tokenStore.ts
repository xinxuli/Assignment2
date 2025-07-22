import { makeAutoObservable } from "mobx";

export class TokenStore {
    constructor() {
        makeAutoObservable(this)
    }

    private _token = ''
    public get token() {
        return this._token
    }

    public set token(tk : string) {
        this._token = tk
    }
}

export class UserInfoStore{
    constructor() {
        makeAutoObservable(this)
    }

    private _nickname = ''
    public get nickname() {
        return this._nickname // avoid confusion with nickname getter and setter
    }

    public set nickname(nk :string){
        this._nickname = nk
    }
}