<template>
    <article class="post" itemscope itemtype="http://schema.org/BlogPosting">
        <header>
            <h1 class="posttitle" itemprop="name headline">
                {{post.name}}
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
        <div class="content" itemprop="articleBody" v-html="post.content">
        </div>
    </article>
</template>

<script>
    import axios from "axios";

    export default {
        name: "PostItem",
        data() {
            return {
                post: null
            }
        },
        mounted() {
            axios.get("http://localhost:80/api/post/"+this.$route.params.slug)
                .then(response => {
                    this.post = response.data
                })
                .catch()
        }
    }
</script>