<template>
    <div id="app" class="content index py4">
        <Header />
        <slot/>
        <Footer />
    </div>
</template>

<script>
    import Footer from "@/components/Footer";
    import Header from "@/components/Header";
    export default {
        components: {Header, Footer}
    }
</script>