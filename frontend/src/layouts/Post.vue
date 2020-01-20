<template>
    <div class="content index py4">
        <div id="header-post">
            <a id="menu-icon" href="#"><i class="fas fa-bars fa-lg"></i></a>
            <a id="menu-icon-tablet" href="#"><i class="fas fa-bars fa-lg"></i></a>
            <a id="top-icon-tablet" href="#" onclick="$('html, body').animate({ scrollTop: 0 }, 'fast');" style="display:none;"><i class="fas fa-chevron-up fa-lg"></i></a>
            <span id="menu">
                <span id="nav">
                  <ul>
                      <li><router-link to="/">Home</router-link></li>
                      <li><router-link to="/about">About</router-link></li>
                  </ul>
                </span>
            </span>
        </div>
        <slot/>
        <Footer/>
    </div>
</template>

<script>
    import Footer from "@/components/Footer";
    export default {
        name: "Post",
        components: {Footer}
    }
</script>