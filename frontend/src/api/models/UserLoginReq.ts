/* tslint:disable */
/* eslint-disable */
/**
 * Assignment2
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface UserLoginReq
 */
export interface UserLoginReq {
    /**
     * 
     * @type {string}
     * @memberof UserLoginReq
     */
    username: string;
    /**
     * 
     * @type {string}
     * @memberof UserLoginReq
     */
    password: string;
}

/**
 * Check if a given object implements the UserLoginReq interface.
 */
export function instanceOfUserLoginReq(value: object): value is UserLoginReq {
    if (!('username' in value) || value['username'] === undefined) return false;
    if (!('password' in value) || value['password'] === undefined) return false;
    return true;
}

export function UserLoginReqFromJSON(json: any): UserLoginReq {
    return UserLoginReqFromJSONTyped(json, false);
}

export function UserLoginReqFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserLoginReq {
    if (json == null) {
        return json;
    }
    return {
        
        'username': json['username'],
        'password': json['password'],
    };
}

export function UserLoginReqToJSON(json: any): UserLoginReq {
    return UserLoginReqToJSONTyped(json, false);
}

export function UserLoginReqToJSONTyped(value?: UserLoginReq | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'username': value['username'],
        'password': value['password'],
    };
}

