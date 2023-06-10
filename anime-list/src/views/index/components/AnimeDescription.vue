<script setup>
defineProps({
  content: {
    type: String,
    default: '',
  }
});
</script>

<template>
  <div class="description-container">
    <el-text truncated size="small">{{content}}</el-text>
  </div>
</template>

<style scoped>
.description-container {
  width: 100%;
  height: 20px;
  background-color: rgba(177, 179, 184, 0.85);
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 0 0 0.5rem 0.5rem;
}
</style>
