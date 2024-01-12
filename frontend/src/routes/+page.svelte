<script lang="ts">
    let products = [];

    async function fetchProducts() {
        const response = await fetch('/products');
        products = await response.json();
    }

    fetchProducts();
</script>

<h1>Products</h1>
<pre>{JSON.stringify(products, null, 2)}</pre>