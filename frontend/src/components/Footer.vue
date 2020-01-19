<template>
    <footer id="footer">
        <div class="footer-left">
            Copyright &copy; {{ new Date() | moment("YYYY") }} fredrsf.ru
        </div>
        <!--<div class="footer-right">
          <nav>
            <ul>
              <% for (var i in theme.nav) { %>
              <li><a href="<%- url_for(theme.nav[i]) %>"><%= __('nav.'+i).replace("nav.", "") %></a></li>
              <% } %>
            </ul>
          </nav>
        </div>-->
    </footer>
</template>

<script>
    export default {
        name: "Footer"
    }
</script>