<template>
  <section id="writing">
    <h2>2019</h2>
    <ul class="post-list">
      <li class="post-item" v-for="post in postsList" :key="post.id">
        <div class="meta">
          <time :datetime="post.created" itemprop="dateModified">
            {{ post.created | moment("DD MMM YYYY") }}
          </time>
        </div>
        <span>
          <router-link :to="{ name: 'post', params: { slug: post.slug }}">{{post.name}}</router-link>
        </span>
      </li>
    </ul>
  </section>
</template>

<script>
import axios from "axios";

export default {
  name: 'PostList',
  data() {
    return {
      postsList: []
    }
  },
  mounted() {
    axios.get("http://localhost:80/api/post")
            .then(response => {
              this.postsList = response.data
            })
            .catch()
  }
}
</script>
