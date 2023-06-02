<script setup>
import { config } from "@/layout/setting/config.js";
import { ref } from "vue";

const defaultActive = ref('');
const initDefaultActive = () => {
  defaultActive.value = location.hash.substring(1);
}

initDefaultActive()
</script>

<template>
  <el-menu
      :default-active="defaultActive"
      class="menu-container"
      router
  >
    <el-menu-item
        v-for="item in config"
        :key="item.id"
        :index="item.path"
    >
      <el-icon><Menu /></el-icon>
      <span>{{item.name}}</span>
    </el-menu-item>
  </el-menu>
</template>

<style scoped>
.menu-container {
  height: 100%;
}
</style>
