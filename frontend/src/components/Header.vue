<template>
    <header id="header">
        <router-link to="/">
            <div id="title">
                <h1>fredrsf.ru</h1>
            </div>
        </router-link>
        <div id="nav">
            <ul>
                <li class="icon">
                    <a href="#"><i class="fas fa-bars fa-2x"></i></a>
                </li>
                <li><router-link to="/">Home</router-link></li>
                <li><router-link to="/about">About</router-link></li>
            </ul>
        </div>
    </header>
</template>

<script>
    export default {
        name: "Header"
    }
</script>

<style scoped>

</style>