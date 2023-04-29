<template>
  <transition name="fade" mode="out-in" appear>
    <el-row>
      <el-col :span="3"></el-col>
      <el-col :span="18">
        <el-card>
          <div class="setting-item">
            <span>Close Server</span>
            <el-button
                :type="'danger'"
                @click="stopServer"
            >Stop</el-button>
          </div>
        </el-card>
      </el-col>
      <el-col :span="3"></el-col>
    </el-row>
  </transition>
</template>

<script setup>
import request from "@/utils/axiosInstance.js";
import {ElNotification} from "element-plus";

const stopServer = () => {
  request.get(
      '/stop',
      {
        params: {

        }
      }
  ).then(() => {
    ElNotification({
      title: 'Success',
      message: 'Server is Close.',
      type: 'success',
      position: 'bottom-right'
    });
  }).catch(error => {
    console.log(error);
  })
}
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: all 200ms ease-out;
}

.fade-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.fade-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}

.el-row:first-child {
  margin-top: 20px;
}

.el-row {
  margin-bottom: 20px;
}

.el-row:last-child {
  margin-bottom: 0;
}

.setting-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
