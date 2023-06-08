<script setup>
import useServer from "@/store/modules/server.js";
import {onMounted} from "vue";

onMounted(() => {
  useServer().handleAbout();
});
</script>

<template>
  <el-card>
    <el-row :gutter="20" :align="'middle'">
      <el-col :span="6" :offset="1" class="item-name">
        <el-text>Version Code</el-text>
      </el-col>
      <el-col :span="12">
        <el-text>{{ useServer().versionCode }}</el-text>
      </el-col>
    </el-row>
    <el-row :gutter="20" :align="'middle'">
      <el-col :span="6" :offset="1" class="item-name">
        <el-text>Version Name</el-text>
      </el-col>
      <el-col :span="12">
        <el-text>{{ useServer().versionName }}</el-text>
      </el-col>
    </el-row>
    <el-row :gutter="20" :align="'middle'">
      <el-col :span="6" :offset="1" class="item-name">
        <el-text>Source Code</el-text>
      </el-col>
      <el-col :span="12">
        <el-link :type="'primary'" href="https://github.com/lin-Dongw/AnimeList" target="_blank">Github</el-link>
      </el-col>
    </el-row>
  </el-card>
</template>

<style scoped>
.el-row {
  margin-bottom: 20px;
}

.el-row:first-child {
  margin-top: 20px;
}

.item-name {
  text-align: end;
}
</style>
