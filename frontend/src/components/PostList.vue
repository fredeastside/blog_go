<template>
  <section id="writing">
    <div v-for="year in postsList" :key="year">
      <h2>{{year}}</h2>
      <ul class="post-list">
        <li class="post-item" v-for="post in postsList[year]" :key="post.id">
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
    </div>
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
              const ordered = {}
              Object.keys(response.data).sort((a, b) => b-a).forEach(function(key) {
                ordered[key] = response.data[key]
              })
              this.postsList = ordered
            })
            .catch()
  }
}
</script>
