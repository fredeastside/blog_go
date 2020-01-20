<template>
    <footer id="footer">
        <div class="footer-left">
            Copyright &copy; {{ new Date() | moment("YYYY") }} fredrsf.ru
        </div>
        <div class="footer-right">
          <nav>
            <ul>
                <li><router-link to="/">Home</router-link></li>
                <li><router-link to="/about">About</router-link></li>
            </ul>
          </nav>
        </div>
    </footer>
</template>

<script>
    export default {
        name: "Footer"
    }
</script>