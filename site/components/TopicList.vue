<template>
  <ul class="topic-list">
    <li v-for="article in articles" :key="article.article_id" class="topic-item">
      <div class="topic-avatar" :href="'/user/' + article.user.id" :title="article.user.nickname">
        <avatar :user="article.user" />
      </div>
      <div class="topic-main-content">
        <div class="topic-top">
          <div class="topic-userinfo">
            <avatar class="topic-inline-avatar" :user="article.user" size="20" />
            <a :href="'/user/' + article.user.id">{{ article.user.nickname }}</a>
          </div>
          <div class="topic-time">发布于{{ article.create_time | prettyDate }}</div>
        </div>
        <div class="topic-content" :class="{ 'topic-tweet': false }">
          <template>
            <h1 class="topic-title">
              <a :href="'/topics/' + article.article_id">{{ article.title }}</a>
            </h1>
            <a :href="'/topics/' + article.article_id" class="topic-summary">{{ article.title }}</a>
          </template>
          <!-- <template v-if="topic.type === 1">
            <a
              v-if="topic.content"
              :href="'/topic/' + topic.topicId"
              class="topic-summary"
            >{{ topic.content }}</a>
            <ul v-if="topic.imageList && topic.imageList.length" class="topic-image-list">
              <li v-for="(image, index) in topic.imageList" :key="index">
                <a :href="'/topic/' + topic.topicId" class="image-item">
                  <img v-lazy="image.preview" />
                </a>
              </li>
            </ul>
          </template>-->
        </div>
        <div class="topic-handlers">
          <div class="btn" :class="{ liked: false }" @click="like(article)">
            <i class="iconfont icon-like" />
            {{ article.liked ? '已赞' : '赞' }}
            <span
              v-if="article.like_count > 0"
            >{{ article.like_count }}</span>
          </div>
          <div class="btn" @click="toTopicDetail(article.article_id)">
            <i class="iconfont icon-comments" />评论
            <span v-if="article.comment_count > 0">{{ article.comment_count }}</span>
          </div>
          <div class="btn" @click="toTopicDetail(article.article_id)">
            <i class="iconfont icon-read" />浏览
            <span v-if="article.view_count > 0">{{ article.view_count }}</span>
          </div>
        </div>
      </div>
    </li>
  </ul>
</template>

<script>
import Avatar from '@/components/Avatar'
export default {
  components: {
    Avatar,
  },
  props: {
    articles: {
      type: Array,
      default() {
        return []
      },
      required: false,
    },
    showAvatar: {
      type: Boolean,
      default: true,
    },
    showAd: {
      type: Boolean,
      default: false,
    },
  },
  methods: {
    async like(article) {
      try {
        let data = {
          article_id: article.article_id,
          user_id: article.user.id,
        }
        await this.$axios.post('/api/topics/like', data)
        article.liked = true
        article.like_count++
        this.$message.success('点赞成功')
      } catch (e) {
        if (e.errorCode === 1) {
          this.$msgSignIn()
        } else {
          this.$message.error(e.message || e)
        }
      }
    },
    toTopicDetail(topicId) {
      this.$linkTo(`/topic/${topicId}`)
    },
  },
}
</script>

<style lang="scss" scoped></style>
