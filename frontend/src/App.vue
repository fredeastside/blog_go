<template>
  <component :is="layout">
    <router-view/>
  </component>
</template>

<style lang="stylus">
  @import '../css/style.styl'
</style>

<script>
  const default_layout = 'default'
  // @ is an alias to /src
  export default {
    computed: {
      layout() {
        //console.log(this.$route.meta.layout);
        return (this.$route.meta.layout || default_layout)+'-layout'
      }
    }
  }
</script>
