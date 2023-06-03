<script setup>
import useServer from "@/store/modules/server.js";
import {onMounted} from "vue";
import {open} from "@/api/open.js";
import {ElNotification} from "element-plus";

const openDir = async (folderName) => {
  await open(folderName)
      .then(response => {
        if (response.code === 200)
        {
          ElNotification({
            title: 'Success',
            message: 'Open folder success.',
            type: 'success',
            position: 'bottom-right'
          });
        }
        else if (response.code === 404)
        {
          ElNotification({
            title: 'Error',
            message: 'Open folder fail.',
            type: 'error',
            position: 'bottom-right'
          });
        }
      })
      .catch(error => {
        console.log(error);
      });
}

onMounted(() => {
  useServer().handleAbout();
});
</script>

<template>
  <el-card>
    <el-row :gutter="20" :align="'middle'">
      <el-col :span="6" :offset="1" class="item-name">
        <el-text>Root Dir</el-text>
      </el-col>
      <el-col :span="12">
        <el-button :type="'primary'" @click="openDir(useServer().rootDir)">open</el-button>
      </el-col>
    </el-row>
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
