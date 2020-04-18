<template>
  <header id="header">
    <router-link to="/">
      <div id="logo" :style="{ backgroundImage: 'url('+ require('@/assets/logo.png') +')' }"></div>
      <div id="title">
        <h1>fredrsf.ru</h1>
      </div>
    </router-link>
    <div id="nav">
      <ul>
        <li class="icon">
          <a href="#"><font-awesome-icon :icon="['fas', 'bars']" size="2x" /></a>
        </li>
        <li><router-link to="/">Home</router-link></li>
        <li><router-link to="/about">About</router-link></li>
      </ul>
    </div>
  </header>
</template>

<script>
import $ from "jquery";
export default {
  name: "Header",
  mounted() {
    /**
     * Shows the responsive navigation menu on mobile.
     */
    $("#header > #nav > ul > .icon").click(function() {
      $("#header > #nav > ul").toggleClass("responsive");
    });
  }
};
</script>

<style scoped></style>
