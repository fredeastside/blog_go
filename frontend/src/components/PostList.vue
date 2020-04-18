<template>
  <section id="writing">
    <span class="h1">Writing</span>
    <div v-for="year in Object.keys(postsList).sort((a, b) => parseInt(b)-parseInt(a))" :key="year">
      <h2>{{ year }}</h2>
      <ul class="post-list">
        <li class="post-item" v-for="post in postsList[year]" :key="post.id">
          <div class="meta">
            <time :datetime="post.created" itemprop="dateModified">
              {{ post.created | moment("DD MMM YYYY") }}
            </time>
          </div>
          <span>
            <router-link :to="{ name: 'post', params: { slug: post.slug } }">{{
              post.name
            }}</router-link>
          </span>
        </li>
      </ul>
    </div>
  </section>
</template>

<script>
import http from "@/http";

export default {
  name: "PostList",
  data() {
    return {
      postsList: []
    };
  },
  created() {
    http
      .get("/post")
      .then(response => {
        this.postsList = response.data;
      })
      .catch((error) => {
        console.log(error);
      });
  }
};
</script>
