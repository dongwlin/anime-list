<template>
  <transition name="fade" mode="out-in" appear>
    <el-row>
      <el-col :span="3"></el-col>
      <el-col :span="18">
        <el-card>
          <div>versionCode: {{about.versionCode}}</div>
          <div>versionName: {{about.versionName}}</div>
        </el-card>
      </el-col>
      <el-col :span="3"></el-col>
    </el-row>
  </transition>
</template>

<script setup>
import request from "@/utils/axiosInstance.js";
import {onMounted, reactive} from "vue";

const about = reactive({
  versionCode: 0,
  versionName: ''
});

const getAbout = () => {
  request.get(
      '/about',
      {
        params: {

        }
      }
  ).then(response => {
    about.versionCode = response.data.versionCode;
    about.versionName = response.data.versionName;
  }).catch(error => {
    console.log(error);
  });
}

onMounted(() => {
  getAbout();
});
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
</style>
