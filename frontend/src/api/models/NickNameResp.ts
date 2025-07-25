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
 * @interface NickNameResp
 */
export interface NickNameResp {
    /**
     * 
     * @type {string}
     * @memberof NickNameResp
     */
    nickName: string;
}

/**
 * Check if a given object implements the NickNameResp interface.
 */
export function instanceOfNickNameResp(value: object): value is NickNameResp {
    if (!('nickName' in value) || value['nickName'] === undefined) return false;
    return true;
}

export function NickNameRespFromJSON(json: any): NickNameResp {
    return NickNameRespFromJSONTyped(json, false);
}

export function NickNameRespFromJSONTyped(json: any, ignoreDiscriminator: boolean): NickNameResp {
    if (json == null) {
        return json;
    }
    return {
        
        'nickName': json['nick_name'],
    };
}

export function NickNameRespToJSON(json: any): NickNameResp {
    return NickNameRespToJSONTyped(json, false);
}

export function NickNameRespToJSONTyped(value?: NickNameResp | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'nick_name': value['nickName'],
    };
}

