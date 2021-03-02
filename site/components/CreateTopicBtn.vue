<template>
  <div>
    <el-dropdown placement="bottom" trigger="hover">
      <span class="el-dropdown-link">
        <el-button @click="handlePostCommand" type="primary" icon="el-icon-plus">发表</el-button>
      </span>
    </el-dropdown>
  </div>
</template>
<script>
export default {
  data() {
    return {}
  },
  computed: {},
  methods: {
    handlePostCommand() {
      this.$linkTo('/topic/create')
    },
  },
}
</script>
<style lang="scss" scoped></style>
