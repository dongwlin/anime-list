<script setup>
import { useRoute, useRouter } from 'vue-router';
import {ref, watch} from 'vue';

const route = useRoute();
const router = useRouter();
const title = ref('');

const initTitle = () => {
  title.value = route.name;
}

const goIndex = () => {
  router.push('/index');
}

watch(route, () => {
  initTitle()
}, {deep: true, immediate: true})
</script>

<template>
  <el-page-header @back="goIndex">
    <template #content>
      <span class="title">{{ title }}</span>
    </template>
  </el-page-header>
</template>

<style scoped>
.title {
  font-weight: 600;
  font-size: large;
}
</style>
