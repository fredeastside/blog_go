<template>
  <fragment>
    <article class="post" itemscope itemtype="http://schema.org/BlogPosting">
      <header>
        <h1 class="posttitle" itemprop="name headline">
          {{ post.name }}
        </h1>
        <div class="meta">
          <!--span class="author" itemprop="author" itemscope itemtype="http://schema.org/Person">
                      <span itemprop="name">Author</span>
                  </span-->
          <time :datetime="post.created" itemprop="datePublished">
            {{ post.created | moment("DD MMM YYYY") }}
          </time>
        </div>
      </header>
      <div class="content" itemprop="articleBody" v-html="post.content"></div>
    </article>
    <div id="footer-post-container">
      <div id="footer-post">

        <div id="nav-footer" style="display: none">
          <ul>
            <li><router-link to="/">Home</router-link></li>
            <li><router-link to="/about">About</router-link></li>
          </ul>
        </div>

        <div id="actions-footer">
          <a id="menu" class="icon" href="#" v-on:click="toggleMenu($event)">
            <font-awesome-icon :icon="['fas', 'bars']" size="lg" aria-hidden="true"
            /> Menu</a>
        </div>

      </div>
    </div>
  </fragment>
</template>

<script>
import http from "@/http";
import "highlight.js/styles/atom-one-dark.css";
import hljs from "highlight.js";
import $ from "jquery";

export default {
  name: "PostItem",
  data() {
    return {
      post: { name: "", created: null, content: "" }
    };
  },
  created() {
    http
      .get("/post/" + this.$route.params.slug)
      .then(response => {
        if (response.data === null) {
          this.$router.push({name: '404'});
          return;
        }
        this.post = response.data;
      })
      .catch((error)=>{
        console.log(error);
      });
  },
  updated() {
    this.highlightPost();
  },
  methods: {
    highlightPost() {
      document.querySelectorAll('pre code').forEach((block) => {
        hljs.highlightBlock(block);
      });
    },
    toggleMenu(event) {
      event.preventDefault();
      $('#nav-footer').toggle();
      return false;
    }
  }
};
</script>
