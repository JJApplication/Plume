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
    }
    .header .header-title {
        font-size: 1.8rem;
        font-weight: bold;
        color: #404040;
    }
</style>