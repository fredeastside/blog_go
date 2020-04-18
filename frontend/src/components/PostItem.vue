<template>
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
</template>

<script>
import http from "@/http";
import "highlight.js/styles/atom-one-dark.css";
import hljs from "highlight.js"

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
    }
  }
};
</script>
