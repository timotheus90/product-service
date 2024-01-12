/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { models_Product } from '../models/models_Product';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class DefaultService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * Get all products
     * @returns models_Product OK
     * @throws ApiError
     */
    public getAllProducts(): CancelablePromise<Array<models_Product>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/products',
        });
    }
    /**
     * Create a new product with the input payload
     * @param product Create product
     * @returns models_Product Created
     * @throws ApiError
     */
    public createProduct(
        product: models_Product,
    ): CancelablePromise<models_Product> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/products',
            body: product,
        });
    }
}
