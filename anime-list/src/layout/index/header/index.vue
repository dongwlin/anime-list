<script setup>
import { Setting } from "@element-plus/icons-vue";
import {useRouter} from "vue-router";

const router = useRouter();

const goSetting = () => {
  router.push('/setting');
}
</script>

<template>
  <div class="container">
    <span class="title">Anime List</span>
    <el-button :icon="Setting" @click="goSetting"></el-button>
  </div>
</template>

<style scoped>
.container {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.title {
  font-size: 36px;
  font-style: italic;
  font-weight: 700;
  color: #3f3f3f;
}
</style>
