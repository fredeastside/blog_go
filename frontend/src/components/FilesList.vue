<template>
  <section id="writing">
    <span class="h1">Files</span>
    <ul class="post-list">
      <li class="post-item" v-for="file in files" :key="file.id">
        <span>{{ file.name }}</span>
      </li>
    </ul>
  </section>
</template>

<script>
//import http from "@/http";

export default {
  name: "FilesList",
  data() {
    return {
      files: [
        {
          id: 1,
          name: "file.txt"
        },
        {
          id: 2,
          name: "file2.txt"
        },
        {
          id: 3,
          name: "file2.txt"
        }
      ]
    };
  }
};
</script>
