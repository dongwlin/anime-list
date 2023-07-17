<script setup>
import useServer from "@/store/modules/server.js";
import { useOpenDir } from "@/hooks/open.js";
import { useStop } from "@/hooks/stop.js";

const serverStore = useServer();

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
        <el-button :type="'primary'" @click="useOpenDir(serverStore.rootDir)">open</el-button>
      </el-col>
    </el-row>
    <el-row :gutter="20" :align="'middle'">
      <el-col :span="6" :offset="1" class="item-name">
        <el-text>Server</el-text>
      </el-col>
      <el-col :span="12">
        <el-button
            type="danger"
            @click="useStop"
            :disabled="!serverStore.status"
        >stop
        </el-button>
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
