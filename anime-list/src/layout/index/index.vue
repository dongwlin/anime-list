<script setup>
import Header from '@/layout/index/header/index.vue';
</script>

<template>
  <el-container class="app-wrapper">
    <el-header class="header-container">
      <Header></Header>
    </el-header>
    <el-main class="main-container">
      <router-view></router-view>
    </el-main>
  </el-container>
</template>

<style scoped>
.app-wrapper {
  height: 100vh;
  background: #f2f3f5;
}

.header-container {
  background: #ffffff;
  display: flex;
  align-items: center;
}
</style>
