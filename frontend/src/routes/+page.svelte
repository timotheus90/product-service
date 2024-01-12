<script lang="ts">
    import type {models_Product} from "../api_client";
    import {ApiClient,} from "../api_client";

    let products: models_Product[] = [];

    const appClient = new ApiClient({
        BASE: 'http://localhost:8080',
    });

    async function fetchProducts() {
        products = await appClient.default.getAllProducts()
    }

    fetchProducts();
</script>

<h1>Products</h1>
<pre>{JSON.stringify(products, null, 2)}</pre>