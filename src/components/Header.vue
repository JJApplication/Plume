<template>
    <div class="header">
        <p class="header-title">{{name}}</p>
    </div>
</template>

<script>
export default {
    name: "Header",
    props: ['name']
}
</script>

<style scoped>
    .header {
        text-align: left;
        padding: 4px 0;
    }
    .header .header-title {
        font-size: 1.8rem;
        font-weight: bold;
        color: var(--light-text-title-color);
    }
</style>