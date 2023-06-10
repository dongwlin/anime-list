<script setup>
import { stop } from "@/api/stop.js";
import {ElMessage, ElNotification} from "element-plus";
import useStore from '@/store';
import {open} from "@/api/open.js";

const serverStore = useStore().useServer();

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
        else if (response.code === 500)
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

const handleStop = async () => {
  await stop()
      .then(() => {
        serverStore.status = false;
        ElMessage({
          message: 'The server is stopped.',
          type: 'success'
        });
      })
      .catch(error => {
        console.log(error);
        ElMessage({
          message: 'The server is failed to stop.',
          type: 'error'
        });
      })
}

serverStore.handleHi();
</script>

<template>
  <el-card>
    <el-row :gutter="20" :align="'middle'">
      <el-col :span="6" :offset="1" class="item-name">
        <el-text>Status</el-text>
      </el-col>
      <el-col :span="12">
        <el-tag :type="serverStore.status ? 'success' : 'danger'">
          {{ serverStore.status ? 'Running' : 'Stopped' }}
        </el-tag>
      </el-col>
    </el-row>
    <el-divider></el-divider>
    <el-row :gutter="20" :align="'middle'">
      <el-col :span="6" :offset="1" class="item-name">
        <el-text>Root Dir</el-text>
      </el-col>
      <el-col :span="12">
        <el-button :type="'primary'" @click="openDir(serverStore.rootDir)">open</el-button>
      </el-col>
    </el-row>
    <el-row :gutter="20" :align="'middle'">
      <el-col :span="6" :offset="1" class="item-name">
        <el-text>Server</el-text>
      </el-col>
      <el-col :span="12">
        <el-button
            type="danger"
            @click="handleStop"
            :disabled="!serverStore.status"
        >stop</el-button>
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
