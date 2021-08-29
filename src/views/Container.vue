<template>
    <div>
        <Header name="容器"></Header>
        <div class="docker-top">
            <font-awesome-icon :icon="['fab', 'docker']"></font-awesome-icon>
        </div>
    </div>
</template>

<script>
import Header from "../components/Header";
export default {
    name: "Container",
    components: {Header}
}
</script>

<style scoped>
    .docker-top {
        padding: 10px;
        background-color: var(--light-bg-docker-color, #2a2b2b);
        border-radius: 4px;
        margin: 10px 0;
        font-size: 1.2rem;
    }
</style>