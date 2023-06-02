<script setup>
import Header from '@/layout/setting/header/index.vue';
import Menu from '@/layout/setting/menu/index.vue';
import useServer from "@/store/modules/server.js";
import {onMounted} from "vue";

onMounted(() => {
  useServer().handleHi();
  useServer().handleAbout();
});
</script>

<template>
  <el-container class="app-wrapper">
    <el-aside width="200px">
      <Menu></Menu>
    </el-aside>
    <el-container>
      <el-header class="header-container">
        <Header></Header>
      </el-header>
      <el-main class="main-container">
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<style scoped>
.app-wrapper {
  height: 100vh;
}

.header-container {
  display: flex;
  align-items: center;
}

.main-container {
  background: #F2F3F5;
}
</style>
