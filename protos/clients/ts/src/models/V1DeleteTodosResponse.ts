/* tslint:disable */
/* eslint-disable */
/**
 * go_app/v1/api.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: version not set
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
 * @interface V1DeleteTodosResponse
 */
export interface V1DeleteTodosResponse {
    /**
     * 
     * @type {Array<string>}
     * @memberof V1DeleteTodosResponse
     */
    errors?: Array<string>;
}

/**
 * Check if a given object implements the V1DeleteTodosResponse interface.
 */
export function instanceOfV1DeleteTodosResponse(value: object): boolean {
    return true;
}

export function V1DeleteTodosResponseFromJSON(json: any): V1DeleteTodosResponse {
    return V1DeleteTodosResponseFromJSONTyped(json, false);
}

export function V1DeleteTodosResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1DeleteTodosResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'errors': json['errors'] == null ? undefined : json['errors'],
    };
}

export function V1DeleteTodosResponseToJSON(value?: V1DeleteTodosResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'errors': value['errors'],
    };
}

